package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

// データベース接続に必要なハンドルを生成
// 構造体Dbを用意、SQLと接続はしていない
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	fmt.Println(Db)
	fmt.Println("--init--")
	if err != nil {
		panic(err)
	}
}

// 投稿を作成
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		return
	}
	return
}

// idで指定された投稿を取得
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		fmt.Println(err)
		return
	}

	// コメント機能の追加
	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			fmt.Println(err)
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

// 投稿の修正
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// 投稿の削除
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

// limit個まで投稿を取得する
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// コメントを投稿
func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post nit found")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id",
		comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func main() {
	post := Post{
		Content: "Hello World in gwp DB!",
		Author:  "Sau Sheong",
	}

	fmt.Println(post)
	err := post.Create()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println("Read...")
	fmt.Println(readPost)

	readPost.Content = "Bobjour!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(16)
	fmt.Println("view all posts...")
	fmt.Println(posts)

	readPost.Delete()

	fmt.Println("END")
	posts, _ = Posts(16)
	fmt.Println(posts)

	// コメント機能の追加
	post2 := Post{
		Content: "Hello World in gwp DB!",
		Author:  "Sau Sheong コメント待ち",
	}
	post2.Create()

	comment := Comment{
		Content: "いいね❤️",
		Author:  "Joe",
		Post:    &post2,
	}
	comment.Create()

	fmt.Println("Comment機能")
	readPost, _ = GetPost(post2.Id)
	fmt.Println(readPost)
}
