package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DB  DBConfig
	AWS AWSConfig
}

type DBConfig struct {
	Addr   string `json:"addr"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DB     string `json:"db"`
	DBType string `json:"dbType"`
}

type AWSConfig struct {
	Region      string `json:"region"`
	AccessKeyId string `json:"accessKeyId"`
	Secret      string `json:"secret"`
	BucketName  string `json:"bucketName"`
}

var DB DBConfig
var AWS AWSConfig

func LoadConfig(path string) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	var conf Config
	decoder := json.NewDecoder(file)
	decoder.Decode(&conf)
	DB = conf.DB
	AWS = conf.AWS
}
