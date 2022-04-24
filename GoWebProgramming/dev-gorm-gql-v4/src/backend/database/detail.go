package database

import (
	"github.com/jinzhu/gorm"
)

type Detail struct {
	ID      int    `gorm:"column:id;primary_key"`
	Content string `gorm:"column:content"`
	Done    bool   `gorm:"column:done"`
	TodoID  int    `gorm:column"todo_id"`
	UserID  int    `gorm:"column:user_id"`
}

func (u *Detail) TableName() string {
	return "details_v4"
}

type DetailDao interface {
	InsertOne(u *Detail) (*Detail, error)
	FindAll() ([]*Detail, error)
	FindByUserID(userID int) ([]*Detail, error)
	FindByTodoID(todoID int) ([]*Detail, error)
	FindOne(id int) (*Detail, error)
	DeleteOne(id int) error
	DoneOne(id int) (*Detail, error)
}

type detailDao struct {
	db *gorm.DB
}

func NewDetailDao(db *gorm.DB) DetailDao {
	return &detailDao{db: db}
}

func (d *DetailDao) InsertOne(u *Detail) (*Detail, error) {
	res := d.db.Create(&u)
	if err := res.Error; err != nil {
		return u, err
	}
	return u, nil
}

func (d *DetailDao) FindAll() ([]*Detail, error) {
	var todos []*Detail
	res := d.db.Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *DetailDao) FindOne(id int) (*Detail, error) {
	var todos []*Detail
	res := d.db.Where("id = ?", id).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(todos) < 1 {
		return nil, nil
	}
	return todos[0], nil
}

func (d *DetailDao) FindByUserID(userID int) ([]*Detail, error) {
	var todos []*Detail
	res := d.db.Where("user_id = ?", userID).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *DetailDao) DeleteOne(id int) error {
	var todos []*Detail
	res := d.db.Where("id = ?", id).Delete(&todos)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (d *DetailDao) DoneOne(id int) (*Detail, error) {
	todo := Todo{}
	// fmt.Println("call for done.")
	res := d.db.Where("id = ?", id).Find(&todo).Update("done", true)
	// fmt.Println("res", res)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
