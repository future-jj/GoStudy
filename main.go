// main.go
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func main() {
	config := Config{}

	// 读取 YAML 配置文件
	data, err := ioutil.ReadFile("config.yaml") // 挂载的路径
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 解析 YAML
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Host: %s, Port: %d\n", config.Host, config.Port)
}
