package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"openAI/model"
	"openAI/pool"
	"time"
)

type FollowAgentRequest struct {
	AgentId int    `json:"agent_id"`
	WxId    string `json:"wx_id"`
}
type FollowAgentResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func FollowAgent(w http.ResponseWriter, r *http.Request) {

	var agentId FollowAgentRequest
	err := json.NewDecoder(r.Body).Decode(&agentId)

	if err != nil {
		http.Error(w, "FollowAgent解析参数失败", http.StatusBadRequest)
		return
	}
	// 更新follow 关注
	db := pool.DB
	stmt1, err := db.Prepare(`UPDATE agent SET is_followed=? WHERE id=?`)
	checkErr(err)
	res, err := stmt1.Exec(true, agentId.AgentId)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)

	// 插入到对话列表
	rows, err := db.Query("SELECT * FROM agent WHERE id = ?", agentId.AgentId)
	for rows.Next() {
		var agent model.Agent
		stirs, er := rows.Columns()
		fmt.Println(stirs, er)
		err = rows.Scan(&agent.ID, &agent.Avatar, &agent.Name, &agent.Description,
			&agent.Firepower, &agent.Author, &agent.WxId, &agent.CreatedAt, &agent.UpdatedAt, &agent.IsFollowed)
		checkErr(err)
		fmt.Println(agent)
		stmt3, err := db.Prepare(`INSERT conversation_list (conversation_name, avatar,wx_id,conversation_id,updated_at,description) values (?,?,?,?,?,?)`)
		checkErr(err)
		res, err := stmt3.Exec(agent.Name, agent.Avatar, agent.WxId, agent.ID, time.Now(), agent.Description)
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		log.Println("关注成功，关注Id={},wxId={}", id, agentId.WxId)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response FollowAgentResponse
	response.Status = 0
	response.Message = "success"
	json.NewEncoder(w).Encode(response)
}
