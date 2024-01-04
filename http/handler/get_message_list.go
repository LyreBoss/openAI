package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"openAI/model"
	"openAI/pool"
)

// WxIdRequest 请求参数
type WxIdReq struct {
	WxId string `json:"wx_id,omitempty"`
}

type MessageListResponse struct {
	Users []model.User `json:"users"`
}

func GetMessageList(w http.ResponseWriter, r *http.Request) {
	// 请求参数
	var wxId WxId
	err := json.NewDecoder(r.Body).Decode(&wxId)
	if err != nil {
		http.Error(w, "解析参数失败", http.StatusBadRequest)
		return
	}
	// 数据库连接池
	db := pool.DB
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)
	// 请求返回
	var userMsgList MessageListResponse
	for rows.Next() {
		var user model.User
		stirs, er := rows.Columns()
		fmt.Println(stirs, er)
		err = rows.Scan(&user.ID, &user.WxID, &user.Name, &user.Avatar, &user.Description, &user.CreatedAt, &user.UpdatedAt, &user.IsTop)
		checkErr(err)
		fmt.Println(user)
		userMsgList.Users = append(userMsgList.Users, user)
	}

	// 设置响应头和写入响应体
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userMsgList)

}
