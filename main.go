package main

import (
	"bytes"
	json "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"openAI/http/handler"
)

var url = "https://api.openai-hk.com/v1/chat/completions"
var apiKey = "hk-v2th7s10000061096bd1f550d8a1a5a021fa787fe3047643"

func main() {
	// 定义 API 处理函数
	http.HandleFunc("/api/chat/completions/v0", openAPIChat)
	http.HandleFunc("/api/chat/getMessageList/v0", handler.GetMessageList)
	http.HandleFunc("/api/chat/GetConversations/v0", handler.GetConversations)
	http.HandleFunc("/api/chat/GetAgents/v0", handler.GetAgents)
	http.HandleFunc("/api/chat/FollowAgent/v0", handler.FollowAgent)
	http.HandleFunc("/api/chat/DeleteConversation/v0", handler.DeleteConversation)

	// 启动 HTTP 服务器并监听指定端口
	port := 8080
	fmt.Printf("Server listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

type AskContent struct {
	AskContent string `json:"ask_content,omitempty"`
}

type OpenAPIResponse struct {
	Status int    `json:"status"`
	Answer string `json:"answer"`
}
type OriginalResponse struct {
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func openAPIChat(w http.ResponseWriter, r *http.Request) {
	// 解码请求体
	var ask AskContent
	err := json.NewDecoder(r.Body).Decode(&ask)
	if err != nil {
		log.Println("解析问题失败:", err)
		http.Error(w, "解析问题失败", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Println("请求gpt的问题:", ask.AskContent)

	// 构建请求数据
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
		log.Println("Error encoding JSON payload:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error making API request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println("请求gpt返回答案:", string(body))

	// 解析响应数据
	var response OpenAPIResponse
	var oResp OriginalResponse
	err = json.Unmarshal(body, &oResp)
	if err != nil {
		log.Println("error msg", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.Status = 0
	response.Answer = oResp.Choices[0].Message.Content

	// 设置响应头和写入响应体
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
