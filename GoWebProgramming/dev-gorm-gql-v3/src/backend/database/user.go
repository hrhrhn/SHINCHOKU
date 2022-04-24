package database

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID   int    `gorm:"column:id;primary_key"`
	Name string `gorm:"column:user_name"`
}

func (u *User) TableName() string {
	return "users"
}

type UserDao interface {
	InsertOne(u *User) (*User, error)
	FindAll() ([]*User, error)
	FindOne(id int) (*User, error)
	FindByTodoID(todoID int) (*User, error)
	DeleteOne(id int) error
}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) InsertOne(u *User) (*User, error) {
	res := d.db.Create(&u)
	if err := res.Error; err != nil {
		return u, err
	}
	return u, nil
}

func (d *userDao) FindAll() ([]*User, error) {
	var users []*User
	res := d.db.Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (d *userDao) FindOne(id int) (*User, error) {
	var users []*User
	res := d.db.Where("id = ?", id).Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return users[0], nil
}

func (d *userDao) FindByTodoID(todoID int) (*User, error) {
	var users []*User
	res := d.db.Table("users").
		Select("users.*").
		Joins("LEFT JOIN todos ON todos.user_id = users.id").
		Where("todos.id = ?", todoID).
		First(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return users[0], nil
}

func (d *userDao) DeleteOne(id int) error {
	var users []*User
	res := d.db.Where("id = ?", id).Delete(&users)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
