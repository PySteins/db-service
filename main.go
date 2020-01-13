package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Comic struct {
	Name     string `gorm:"size:64;unique_index"`
	Status   uint
	Chapters []Chapter
	gorm.Model
}

type Chapter struct {
	ComicId uint `gorm:"index"`
	Number  uint
	Picture string `gorm:"type:text"`
	gorm.Model
}

func main() {
	dsn := "root:acdhb123@tcp(127.0.0.1:3306)/comic?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("open mysql database error")
	}
	defer db.Close()

	db.AutoMigrate(&Comic{}, &Chapter{})

	db.Create(&Comic{
		Name:     "我的漫画",
		Status:   1,
		Chapters: nil,
	})

	var comic Comic
	db.First(&comic, 1)

	fmt.Println(comic)
}
