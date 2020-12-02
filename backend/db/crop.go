package db

import (
	"../model"
	_ "github.com/go-sql-driver/mysql"
)

func FindCropInfo(crop string) ([]model.CropInfo, error) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT name, area, SUM(amount) from crop_amount_area GROUP BY area having name=?", crop)
	if err != nil {
		return nil, err
	}

	err = searchCount(crop)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var info model.CropInfo
	infoList := make([]model.CropInfo, 0)
	for rows.Next() {
		err := rows.Scan(&info.Name, &info.Area, &info.Count)
		if err != nil {
			return nil, err
		}
		infoList = append(infoList, info)
	}

	return infoList, nil
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
