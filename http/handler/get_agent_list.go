package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"openAI/model"
	"openAI/pool"
)

type AgentListResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Agents  []model.Agent `json:"agents"`
}

func GetAgents(w http.ResponseWriter, r *http.Request) {
	// 请求参数
	var wxId WxId
	err := json.NewDecoder(r.Body).Decode(&wxId)
	if err != nil {
		http.Error(w, "GetAgentList解析参数失败", http.StatusBadRequest)
		return
	}
	// 数据库连接池
	db := pool.DB
	rows, err := db.Query("SELECT * FROM agent WHERE wx_id = ?", wxId.WxId)
	checkErr(err)
	// 请求返回
	var agentsList AgentListResponse
	for rows.Next() {
		var agent model.Agent
		stirs, er := rows.Columns()
		fmt.Println(stirs, er)
		err = rows.Scan(&agent.ID, &agent.Avatar, &agent.Name, &agent.Description,
			&agent.Firepower, &agent.Author, &agent.WxId, &agent.CreatedAt, &agent.UpdatedAt, &agent.IsFollowed)
		checkErr(err)
		fmt.Println(agent)
		agentsList.Agents = append(agentsList.Agents, agent)
	}

	// 设置响应头和写入响应体
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(agentsList)

}
