package database

import (
	"github.com/jinzhu/gorm"
)

type Parent struct {
	ID   int    `gorm:"column:id;primary_key"`
	Name string `gorm:"column:parent_name"`
}

func (u *Parent) TableName() string {
	return "parents_v4"
}

type ParentDao interface {
	InsertOne(u *Parent) (*Parent, error)
	FindAll() ([]*Parent, error)
	FindOne(id int) (*Parent, error)
	FindByUserID(userID int) (*Parent, error)
}

type parentDao struct {
	db *gorm.DB
}

func NewParentDao(db *gorm.DB) ParentDao {
	return &parentDao{db: db}
}

func (d *parentDao) InsertOne(u *Parent) (*Parent, error) {
	res := d.db.Create(&u)
	if err := res.Error; err != nil {
		return u, err
	}
	return u, nil
}

func (d *parentDao) FindAll() ([]*Parent, error) {
	var parents []*Parent
	res := d.db.Find(&parents)
	if err := res.Error; err != nil {
		return nil, err
	}
	return parents, nil
}

func (d *parentDao) FindOne(id int) (*Parent, error) {
	var parents []*Parent
	res := d.db.Where("id = ?", id).Find(&parents)
	if err := res.Error; err != nil {
		return nil, err
	}
	return parents[0], nil
}

func (d *parentDao) FindByUserID(userID int) (*Parent, error) {
	var parents []*Parent
	res := d.db.Table("parents_v4").
		Select("parents_v4.*").
		Joins("LEFT JOIN users_v4 ON users_v4.parent_id = parents_v4.id").
		Where("users_v4.id = ?", userID).
		First(&parents)
	if err := res.Error; err != nil {
		return nil, err
	}
	return parents[0], nil
}
