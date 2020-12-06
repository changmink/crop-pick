package db

import (
	"../model"
	_ "github.com/go-sql-driver/mysql"
)

func FindCropInfo(crop string) (model.CropInfo, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * from year_price WHERE crop=?", crop)
	if err != nil {
		return model.CropInfo{}, err
	}
	defer rows.Close()

	var yearPriceList [][13]int
	var yearPrice [13]int
	var name string
	for rows.Next() {
		err := rows.Scan(&yearPrice[0], &yearPrice[1], &yearPrice[2], &yearPrice[3],
			&yearPrice[4], &yearPrice[5], &yearPrice[6], &yearPrice[7], &yearPrice[8],
			&yearPrice[9], &yearPrice[10], &yearPrice[11], &yearPrice[12], &name)
		if err != nil {
			return model.CropInfo{}, err
		}
		yearPriceList = append(yearPriceList, yearPrice)
	}

	var consumerPriceList []model.ConsumerPrice
	var consumerPrice model.ConsumerPrice
	rows1, err := db.Query("SELECT * FROM consumer_price WHERE crop=?", crop)
	if err != nil {
		return model.CropInfo{}, err
	}
	defer rows1.Close()
	for rows1.Next() {
		err := rows1.Scan(&consumerPrice.Crop, &consumerPrice.Kind, &consumerPrice.Price, &consumerPrice.Rate)
		if err != nil {
			return model.CropInfo{}, err
		}
		consumerPriceList = append(consumerPriceList, consumerPrice)
	}

	err = searchCount(crop)
	if err != nil {
		return model.CropInfo{}, err
	}

	return model.CropInfo{Name: crop, YearPrice: yearPriceList, ConsumerPrice: consumerPriceList}, nil
}

func searchCount(crop string) error {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT count FROM crop_count WHERE name=?", crop)
	if err != nil {
		return err
	}
	defer rows.Close()

	var count int
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return err
		}

		_, err = db.Query("UPDATE crop_count SET count=? WHERE name=?", count+1, crop)
		if err != nil {
			return err
		}
	} else {
		_, err = db.Query("INSERT INTO crop_count VALUES(?, 1)", crop)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetCropCount() ([]model.CropCount, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM crop_count")
	if err != nil {
		return nil, err
	}

	countList := make([]model.CropCount, 0)
	var count model.CropCount
	for rows.Next() {
		err := rows.Scan(&count.Name, &count.Count)
		if err != nil {
			return nil, err
		}
		countList = append(countList, count)
	}

	return countList, nil
}
