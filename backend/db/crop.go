package db

import (
	"fmt"

	"../model"
	"../util"
	_ "github.com/go-sql-driver/mysql"
)

func FindInfo(crop string) []model.CropInfo {
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT name, area, SUM(amount) from crop_amount_area GROUP BY area having name=?", crop)
	fmt.Println(crop)
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
