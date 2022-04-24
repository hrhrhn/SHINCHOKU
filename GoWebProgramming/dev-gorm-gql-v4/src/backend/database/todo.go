package database

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID      int    `gorm:"column:id;primary_key"`
	Content string `gorm:"column:content"`
	Done    bool   `gorm:"column:done"`
	UserID  int    `gorm:"column:user_id"`
}

func (u *Todo) TableName() string {
	return "todos_v4"
}

type TodoDao interface {
	InsertOne(u *Todo) (*Todo, error)
	FindAll() ([]*Todo, error)
	FindByUserID(userID int) ([]*Todo, error)
	FindOne(id int) (*Todo, error)
	DeleteOne(id int) error
	DoneOne(id int) (*Todo, error)
}

type todoDao struct {
	db *gorm.DB
}

func NewTodoDao(db *gorm.DB) TodoDao {
	return &todoDao{db: db}
}

func (d *todoDao) InsertOne(u *Todo) (*Todo, error) {
	res := d.db.Create(&u)
	if err := res.Error; err != nil {
		return u, err
	}
	return u, nil
}

func (d *todoDao) FindAll() ([]*Todo, error) {
	var todos []*Todo
	res := d.db.Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *todoDao) FindOne(id int) (*Todo, error) {
	var todos []*Todo
	res := d.db.Where("id = ?", id).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(todos) < 1 {
		return nil, nil
	}
	return todos[0], nil
}

func (d *todoDao) FindByUserID(userID int) ([]*Todo, error) {
	var todos []*Todo
	res := d.db.Where("user_id = ?", userID).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *todoDao) DeleteOne(id int) error {
	var todos []*Todo
	res := d.db.Where("id = ?", id).Delete(&todos)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (d *todoDao) DoneOne(id int) (*Todo, error) {
	todo := Todo{}
	// fmt.Println("call for done.")
	res := d.db.Where("id = ?", id).Find(&todo).Update("done", true)
	// fmt.Println("res", res)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
