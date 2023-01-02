package models

import (
	"time"
)

type Product struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(300)" json:"nama"`
	Deskripsi string    `gorm:"type:text" json:"deskripsi"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Company struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Image     string    `gorm:"type:varchar(255)" json:"image"`
	Address   string    `gorm:"type:varchar(255)" json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"type:varchar(255)" json:"username"`
	Fullname   string    `gorm:"type:varchar(255)" json:"fullname"`
	Password   string    `gorm:"type:varchar(255)" json:"password"`
	NoKtp      string    `gorm:"type:varchar(255)" json:"noktp"`
	NoKaryawan string    `gorm:"type:varchar(255)" json:"nokaryawan"`
	Address    string    `gorm:"type:text" json:"address"`
	Position   string    `gorm:"type:varchar(255)" json:"position"`
	Image      string    `gorm:"type:varchar(255)" json:"image"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type LogsPresent struct {
	Id        uint `gorm:"primaryKey" json:"id"`
	UserId    uint
	CompanyId uint
	CheckIn   time.Time `json:"checkin"`
	CheckOut  time.Time `json:"checkout"`
	Status    string    `gorm:"type:varchar(255)" json:"status"`
	Location  string    `gorm:"type:varchar(255)" json:"location"`
	Uid       string    `gorm:"type:varchar(255)" json:"uid"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// User      User `gorm:"foreignKey:UserId;references:Id"`
	// Company   Company   `gorm:"foreignKey:CompanyId;references:Id"`
}

type LogsLogin struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"userid"`
	IsLogin   string    `gorm:"type:varchar(255)" json:"islogin"`
	LoginAt   time.Time `json:"loginat"`
	Uid       string    `gorm:"type:varchar(255)" json:"uid"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// User      User      `gorm:"foreignKey:UserId;references:Id"`
}
