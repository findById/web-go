package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

const (
	VERSION string = "0.0.1"
)

var (
	Env       Config
	GoVersion = runtime.Version()
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func parseConfig(path string, v interface{}) {
	file, err := os.Open(path)
	if err != nil {
		log.Println("配置文件读取失败:", err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	err = dec.Decode(v)
	if err != nil {
		log.Println("配置文件解析失败:", err)
	}
}
