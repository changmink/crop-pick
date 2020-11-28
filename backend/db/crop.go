package db

import (
	"../model"
	"../util"
	_ "github.com/go-sql-driver/mysql"
)

func FindCropInfo(crop string) []model.CropInfo {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT name, area, SUM(amount) from crop_amount_area GROUP BY area having name=?", crop)
	searchCount(crop)
	util.CheckError(err)

	defer rows.Close()

	var info model.CropInfo
	infoList := make([]model.CropInfo, 0)
	for rows.Next() {
		err := rows.Scan(&info.Name, &info.Area, &info.Count)
		util.CheckError(err)
		infoList = append(infoList, info)
	}

	return infoList
}

func searchCount(crop string) {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT count FROM crop_count WHERE name=?", crop)
	util.CheckError(err)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		util.CheckError(err)

		_, err = db.Query("UPDATE crop_count SET count=? WHERE name=?", count+1, crop)
		util.CheckError(err)
		return
	}

	_, err = db.Query("INSERT INTO crop_count VALUES(?, 1)", crop)
	util.CheckError(err)
}

func GetCropCount() []model.CropCount {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM crop_count")
	util.CheckError(err)

	countList := make([]model.CropCount, 0)
	var count model.CropCount
	for rows.Next() {
		err := rows.Scan(&count.Name, &count.Count)
		util.CheckError(err)
		countList = append(countList, count)
	}

	return countList
}
