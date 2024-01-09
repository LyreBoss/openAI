package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"openAI/model"
	"openAI/pool"
)

// WxId 请求参数
type WxId struct {
	WxId string `json:"wx_id,omitempty"`
}

type ConversationsResponse struct {
	Status      int                  `json:"status"`
	Message     string               `json:"message"`
	Conversions []model.Conversation `json:"conversions"`
}

func GetConversations(w http.ResponseWriter, r *http.Request) {
	// 请求参数
	var wxId WxId
	err := json.NewDecoder(r.Body).Decode(&wxId)
	if err != nil || wxId.WxId == "" {
		http.Error(w, "GetConversations解析参数失败", http.StatusBadRequest)
		return
	}
	// 数据库连接池
	db := pool.DB
	rows, err := db.Query("SELECT * FROM conversation_list where wx_id = ?", wxId.WxId)
	checkErr(err)
	// 请求返回
	var conversations ConversationsResponse
	for rows.Next() {
		var conversation model.Conversation
		stirs, er := rows.Columns()
		fmt.Println(stirs, er)
		err = rows.Scan(&conversation.ID, &conversation.ConversationName,
			&conversation.Avatar, &conversation.WxID, &conversation.ConversationID, &conversation.CreatedAt, &conversation.UpdatedAt,
			&conversation.IsPinned, &conversation.Description)
		checkErr(err)
		fmt.Println(conversation)
		conversations.Conversions = append(conversations.Conversions, conversation)
	}

	// mock

	//var conversations ConversationsResponse
	//
	//conversations.Status = 0
	//conversations.Message = "success"
	//var conversation model.Conversation
	//conversation.ConversationID = "1"
	//conversation.IsPinned = true
	//conversation.ConversationName = "test"
	//conversation.Description = "this is openai"
	//conversation.WxID = "1231"
	//var conversationss []model.Conversation
	//conversationss = make([]model.Conversation, 1)
	//conversationss[0] = conversation
	//conversations.Conversions = conversationss
	//// 设置响应头和写入响应体
	// 设置响应头和写入响应体
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(conversations)

}
func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
