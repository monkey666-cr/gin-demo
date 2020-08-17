package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id"json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"_"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"-"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"say_hello"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt  string `json:"update_at"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type Token struct {
	Token string `json:"token"`
}
