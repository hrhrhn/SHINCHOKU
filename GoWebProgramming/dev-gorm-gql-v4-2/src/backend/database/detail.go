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

func (d *detailDao) InsertOne(u *Detail) (*Detail, error) {
	res := d.db.Create(&u)
	if err := res.Error; err != nil {
		return u, err
	}
	return u, nil
}

func (d *detailDao) FindAll() ([]*Detail, error) {
	var details []*Detail
	res := d.db.Find(&details)
	if err := res.Error; err != nil {
		return nil, err
	}
	return details, nil
}

func (d *detailDao) FindByUserID(userID int) ([]*Detail, error) {
	var details []*Detail
	res := d.db.Where("user_id = ?", userID).Find(&details)
	if err := res.Error; err != nil {
		return nil, err
	}
	return details, nil
}

func (d *detailDao) FindByTodoID(todoID int) ([]*Detail, error) {
	var details []*Detail
	res := d.db.Where("todo_id = ?", todoID).Find(&details)
	if err := res.Error; err != nil {
		return nil, err
	}
	return details, nil
}

func (d *detailDao) FindOne(id int) (*Detail, error) {
	var details []*Detail
	res := d.db.Where("id = ?", id).Find(&details)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(details) < 1 {
		return nil, nil
	}
	return details[0], nil
}

func (d *detailDao) DeleteOne(id int) error {
	var details []*Detail
	res := d.db.Where("id = ?", id).Delete(&details)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (d *detailDao) DoneOne(id int) (*Detail, error) {
	todo := Detail{}
	// fmt.Println("call for done.")
	res := d.db.Where("id = ?", id).Find(&todo).Update("done", true)
	// fmt.Println("res", res)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
