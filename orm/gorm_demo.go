package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
)

func init() {

}

func main() {
	db := getDb()
	defer db.Close()

	dao := GetPODemoDaoInstance(db)

	name := "John Smith"

	p1 := PODemo{Name: name, Age: 18}
	fmt.Println("create:", p1)
	dao.create(p1)
	pdb := dao.retrieve(name)
	fmt.Println("in db:", pdb)
	pdb.Name = fmt.Sprintf("%s_NEW", pdb.Name)
	dao.update(pdb)
	fmt.Println("after update:", dao.list())
	pdb = dao.list()[0]
	dao.updateField(pdb.ID, "name", fmt.Sprintf("%s_NEW", pdb.Name))
	fmt.Println("after update name:", dao.list())
	dao.delete(pdb)
	fmt.Println("after delete:", dao.list())
	dao.transaction()
	fmt.Println("after transaction", dao.list())
}

func getDb() *gorm.DB {
	db, _ := gorm.Open("mysql", "kevin:1234@tcp(localhost:3306)/go-basics?charset=utf8&parseTime=True&loc=Local")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db
}

func initTbl(db *gorm.DB) {
	if !db.HasTable(&PODemo{}) {
		db.CreateTable(&PODemo{})
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			CreateTable(&PODemo{})
	}
}

type PODemoDao struct {
	db *gorm.DB
}

func GetPODemoDaoInstance(db *gorm.DB) *PODemoDao {
	var dao *PODemoDao
	var once sync.Once
	once.Do(func() {
		initTbl(db)
		dao = &PODemoDao{db}
	})
	return dao
}

func (dao PODemoDao) create(o PODemo) {
	dao.db.Create(&o)
}

func (dao PODemoDao) retrieve(name string) PODemo {
	var p PODemo
	dao.db.Model(&PODemo{}).Where(&PODemo{Name: name}).First(&p)
	return p
}

func (dao PODemoDao) list() (pos []PODemo)  {
	dao.db.Find(&pos)
	return
}

func (dao PODemoDao) update(po PODemo) {
	dao.db.Save(po)
}

func (dao PODemoDao) updateField(id int, attr string, newVal interface{}) {
	dao.db.Model(&PODemo{}).
		Where("id=?", id).
		Update(attr, newVal)
}

func (dao PODemoDao) delete(po PODemo) {
	dao.db.Delete(&po)
}

func (dao PODemoDao) deleteByField(attr string, val interface{}) {
	dao.db.Where(fmt.Sprintf("%s=?", attr), val).Delete(&PODemo{})
}

func (dao PODemoDao) transaction() (err error) {
	tx := dao.db.Begin()
	if err = tx.Save(&PODemo{
		Name:"Kevin Tian",
		Age:28,
	}).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return nil
}

type PODemo struct {
	ID   int    `gorm:"primary_key;auto_increment"`
	Name string `gorm:"type:varchar(20);not null;unique_index:idx_name"`
	Age  int    `gorm:"type:int"`
}
