package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var url = "https://api.openai-hk.com/v1/chat/completions"
var apiKey = "hk-v2th7s10000061096bd1f550d8a1a5a021fa787fe3047643"

func main() {
	// 定义 API 处理函数
	http.HandleFunc("/api/chat/completions/v0", openAPIChat)

	// 启动 HTTP 服务器并监听指定端口
	port := 8080
	fmt.Printf("Server listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

type AskContent struct {
	AskContent string `json:"ask_content,omitempty"`
}

func openAPIChat(w http.ResponseWriter, r *http.Request) {
	// 对话问题
	askQuestion := r.FormValue("ask_content")

	data, err := io.ReadAll(r.Body)
	var ask AskContent
	if err = json.Unmarshal(data, &ask); err != nil {
		log.Println("解析问题失败！")
	}
	log.Println(ask)

	log.Println("请求gpt的问题：", askQuestion)
	// 其他参数先写死
	payload := map[string]interface{}{
		"max_tokens":       1200,
		"model":            "gpt-3.5-turbo",
		"temperature":      0.8,
		"top_p":            1,
		"presence_penalty": 1,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are ChatGPT, a large language model trained by OpenAI. Answer as concisely as possible.",
			},
			{
				"role":    "user",
				"content": ask.AskContent,
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON payload:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}

	defer resp.Body.Close()

	// 请根据实际需求解析和处理响应数据
	fmt.Println("Response HTTP Status:", resp.StatusCode)

	// 设置响应的 Content-Type
	w.Header().Set("Content-Type", "application/json")

	// 编写要返回的 JSON 数据
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("请求gpt返回答案：", string(body))

	// 将 JSON 数据写入响应体
	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
