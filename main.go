// main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
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

// 健康检查
func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "OK"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/goodbye", handleGoodbye)
	http.HandleFunc("/health", handleHealth) // 增加健康检查路由

	// 启动 HTTP 服务器
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
