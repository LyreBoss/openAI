package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"openAI/pool"
)

type DeleteConversationReq struct {
	WxId           string `json:"wx_id,omitempty"`
	ConversationID string `json:"conversation_id"`
}

func DeleteConversation(w http.ResponseWriter, r *http.Request) {
	var deleteConversation DeleteConversationReq

	err := json.NewDecoder(r.Body).Decode(&deleteConversation)
	if err != nil {
		http.Error(w, "GetAgentList解析参数失败", http.StatusBadRequest)
		return
	}

	// 开始删除对话
	db := pool.DB
	stmt, err := db.Prepare(`DELETE from conversation_list WHERE conversation_id=?`)
	checkErr(err)
	res, err := stmt.Exec(deleteConversation.ConversationID)
	checkErr(err)
	num, err := res.RowsAffected()
	fmt.Println(num)
	if num > 0 {
		fmt.Println("删除成功对话，对话Id=", deleteConversation.ConversationID)
	}
	var response FollowAgentResponse
	response.Status = 0
	response.Message = "success"
	json.NewEncoder(w).Encode(response)

}
