package models

type Product struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	Nama      string `gorm:"type:varchar(300)" json:"nama"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}
