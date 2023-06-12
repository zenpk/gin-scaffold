package dal

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Id      uint64 `gorm:"primaryKey" json:"-"`
	Name    string `gorm:"size:20; not null" json:"name,omitempty"`
	Gender  string `gorm:"size:10; not null" json:"gender,omitempty"`
	Role    int32  `gorm:"not null; default:1" json:"-"`
	Deleted gorm.DeletedAt
}

func (u User) FindById(ctx context.Context, userId uint64) (user *User, err error) {
	return user, DB.WithContext(ctx).Take(&user, userId).Error
}

func (u User) FindAll(ctx context.Context) (users []*User, err error) {
	return users, DB.WithContext(ctx).Find(&users).Error
}

func (u User) FindAllByIds(ctx context.Context, ids []uint64) (users []*User, err error) {
	return users, DB.WithContext(ctx).Where("id IN ?", ids).Find(&users).Error
}

func (u User) Create(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(&user).Error
}

func (u User) Update(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Save(&user).Error
}
