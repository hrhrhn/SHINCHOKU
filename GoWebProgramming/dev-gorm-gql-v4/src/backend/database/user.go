package database

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `gorm:"column:id;primary_key"`
	Name     string `gorm:"column:user_name"`
	ParentID int    `gorm:"column:parent_id"`
}

func (u *User) TableName() string {
	return "users_v4"
}

type UserDao interface {
	InsertOne(u *User) (*User, error)
	FindAll() ([]*User, error)
	FindOne(id int) (*User, error)
	FindByTodoID(todoID int) (*User, error)
	DeleteOne(id int) error
	FindByParentID(parentID int) ([]*User, error)
}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) InsertOne(u *User) (*User, error) {
	user := User{}
	res := d.db.Where("id = ?", u.ID).Find(&user).Update("user_name", u.Name)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &user, nil
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
	res := d.db.Table("users_v4").
		Select("users_v4.*").
		Joins("LEFT JOIN todos_v4 ON todos_v4.user_id = users_v4.id").
		Where("todos_v4.id = ?", todoID).
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

func (d *userDao) FindByParentID(parentID int) ([]*User, error) {
	var users []*User
	res := d.db.Where("parent_id = ?", parentID).Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return users, nil
}
