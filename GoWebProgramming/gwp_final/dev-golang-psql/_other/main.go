package main

import (
	"database/sql"
	"fmt"
	"log"

	// postgres ドライバ
	_ "github.com/lib/pq"
)

// TestUser : テーブルデータ
type TestUser struct {
	UserID   int
	Password string
}

// メイン関数
func main() {
	fmt.Println("It works!")

	// Db: データベースに接続するためのハンドラ
	var Db *sql.DB
	// Dbの初期化
	Db, err := sql.Open("postgres", "host=devenv_postgresql user=gwp password=gwp dbname=gwp sslmode=disable")
	if err != nil {
		fmt.Println("It works! aaaa")
		log.Fatal(err)
	}

	// SQL文の構築
	sql := "SELECT user_id, user_password FROM test_user WHERE user_id=$1;"

	// preparedstatement の生成
	pstatement, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println("It works! bbbb")
		log.Fatal(err)
	}

	// 検索パラメータ（ユーザID）
	queryID := 1
	// 検索結果格納用の TestUser
	var testUser TestUser

	// queryID を埋め込み SQL の実行、検索結果1件の取得
	err = pstatement.QueryRow(queryID).Scan(&testUser.UserID, &testUser.Password)
	if err != nil {
		fmt.Println("It works! cccc")
		log.Fatal(err)
	}

	// 検索結果の表示
	fmt.Println(testUser.UserID, testUser.Password)
}
