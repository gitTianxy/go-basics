package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {

}

func main() {
	db := getDb()
	defer db.Close()

	initTbl(db)

	//db.AutoMigrate(&Person{})

	p1 := Person{Name: "John Smith", Age: 18}
	p2 := Person{Name: "Jane White", Age: 20}

	db.Create(&p1)
	var p3 Person // identify a Person type for us to store the results in
	db.First(&p3) // Find the first record in the Database and store it in p3

	fmt.Println(p1.Name)
	fmt.Println(p2.Name)
	fmt.Println(p3.Name)

}

func getDb() *gorm.DB {
	db, _ := gorm.Open("mysql", "kevin:1234/golang-basics?charset=utf8")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db
}

func initTbl(db *gorm.DB) {
	if !db.HasTable(&Person{}) {
		db.CreateTable(&Person{})
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			CreateTable(&Person{})
	}
}

type PersonDao struct {
	db *gorm.DB
}

func (dao PersonDao) create(o *interface{}) {
	dao.db.Create(o)
}

func (dao PersonDao) retrieve(name string) {
	dao.db.Model(&Person{}).Where(&Person{Name:name})
}

func (dao PersonDao) update() {

}

func (dao PersonDao) delete() {

}

func (dao PersonDao) transaction() {

}

type Person struct {
	ID   int    `gorm:"primary_key;auto_increment"`
	Name string `gorm:"type:varchar(20);not null;unique_index:idx_name"`
	Age  int    `gorm:"type:int"`
}
