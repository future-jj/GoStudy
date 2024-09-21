// main.go
package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 定义一个简单的响应结构体
type Response struct {
	Message string `json:"message"`
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Welcome to the Go Web API!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 处理 "/goodbye" 路径的请求
func handleGoodbye(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Goodbye!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func welcome(w http.ResponseWriter, r *http.Request) {

}

type Config struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func main() {

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/goodbye", handleGoodbye)

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

	time.Sleep(100000)
}
