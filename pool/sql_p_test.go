package pool

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {

	db, err := sql.Open("mysql", "root:root123..A@tcp(localhost:3306)/openai_schema?charset=utf8&parseTime=True&loc=Local") // 使用本地时间，即东八区，北京时间

	fmt.Println(err)

	fmt.Println(db.Ping())
}
