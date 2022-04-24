package database

import (
	"app/graph/model"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// type Todo struct {
// 	ID      int    `gorm:"column:id;primary_key"`
// 	Content string `gorm:"column:content"`
// 	Done    bool   `gorm:"column:done"`
// 	UserID  int    `gorm:"column:user_id"`
// }

type Todo struct {
	ID        int        `gorm:"column:id;primary_key"`
	Title     string     `gorm:"column:title"`
	Notes     string     `gorm:"column:notes"`
	Completed bool       `gorm:"column:completed"`
	Due       *time.Time `gorm:"column:due"`
}

func (u *Todo) TableName() string {
	return "todos"
}

type TodoDao interface {
	InsertOne(u *Todo) (*Todo, error)
	FixOne(id int, title *string, notes *string, completed *bool, due *time.Time) (*Todo, error)
	Pagenate(completed *bool, orderBy model.TodoOrderFields, page *model.PaginationInput) (*model.TodoConnection, error)
}

type todoDao struct {
	db *gorm.DB
}

func NewTodoDao(db *gorm.DB) TodoDao {
	return &todoDao{db: db}
}

func (d *todoDao) InsertOne(u *Todo) (*Todo, error) {
	fmt.Println(u)
	res := d.db.Create(&u)
	if err := res.Error; err != nil {
		return u, err
	}
	return u, nil
}

func (d *todoDao) FixOne(id int, title *string, notes *string, completed *bool, due *time.Time) (*Todo, error) {
	todo := Todo{}
	var res *gorm.DB
	if title != nil {
		res = d.db.Where("id = ?", id).Find(&todo).Update("title", *title)
	}
	if notes != nil {
		res = d.db.Where("id = ?", id).Find(&todo).Update("notes", *notes)
	}
	if completed != nil {
		res = d.db.Where("id = ?", id).Find(&todo).Update("completed", *completed)
	}
	if due != nil {
		res = d.db.Where("id = ?", id).Find(&todo).Update("due", *due)
	}
	// fmt.Println("res", res)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (d *todoDao) Pagenate(completed *bool, orderBy model.TodoOrderFields, page *model.PaginationInput) (*model.TodoConnection, error) {
	var res *gorm.DB
	var err error

	if completed != nil {
		res = d.db.Where("completed = ?", *completed)
	} else {
		res = d.db
	}
	switch orderBy {
	case model.TodoOrderFieldsLatest:
		res, err = pageDB(res, "id", "desc", page)
		if err != nil {
			return nil, err
		}
		var todos []*model.Todo
		if err := res.Find(&todos).Error; err != nil {
			return &model.TodoConnection{PageInfo: &model.PageInfo{}}, err
		}
		return convertToConnection(todos, orderBy, *page), nil
	case model.TodoOrderFieldsDue:
		res, err = pageDB(res, "due", "asc", page)
		if err != nil {
			return &model.TodoConnection{PageInfo: &model.PageInfo{}}, err
		}

		var todos []*model.Todo
		if err := res.Find(&todos).Error; err != nil {
			return &model.TodoConnection{PageInfo: &model.PageInfo{}}, err
		}

		return convertToConnection(todos, orderBy, *page), nil
	default:
		return &model.TodoConnection{PageInfo: &model.PageInfo{}}, errors.New("invalid order by")
	}
}
