package graph

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-graqhql-postgres/graph/model"
	_ "github.com/lib/pq"
)

type Resolver struct{
        // 追加
	DB *sql.DB
}

// 追加
func (r *Resolver) createUser(input model.NewUser) (*model.User, error){
  // dbにuserを保存する処理
	cmd := "INSERT INTO users (name) VALUES ($1)"
	_, err := r.DB.Exec(cmd, name)
	if err != nil {
		return nil, err
	}
       // 省略しますが、dbに保存されたデータをここで取得して、idとnameという変数に保存
	var user model.User = model.User{
		ID: id,
    Name: name,
	}

	return &user, nil
}