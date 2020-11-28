package config

import (
	"encoding/json"
	"log"
	"os"
)

type DBConfig struct {
	Addr   string `json:"addr"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DB     string `json:"db"`
	DBType string `json:"dbType"`
}

var DB DBConfig

func LoadConfig(path string) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	decoder := json.NewDecoder(file)
	decoder.Decode(&DB)
}
