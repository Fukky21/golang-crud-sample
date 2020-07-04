package infrastructure

import (
	"fmt"
	"log"
	"example.com/m/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const(
	DBUser = "root"
	DBPass = "pass"
	DBProtocol = "tcp(golang-crud-sample_db)"
	DBName = "golang_crud_sample_db"
)

var DB *gorm.DB

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	connect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open("mysql", connect)

	if err != nil {
		log.Println(err.Error())
	}

	DB = db

	return &SqlHandler{Conn: db}
}

func (sqlHandler *SqlHandler) AutoMigration() error {
	db := sqlHandler.Conn

	if err := db.AutoMigrate(&domain.User{}).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (handler *SqlHandler) Find(arg interface{}) (interface{}, error) {
	result := arg

	if err := handler.Conn.Find(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (handler *SqlHandler) FindById(arg interface{}, id int) (interface{}, error) {
	result := arg

	if err := handler.Conn.First(result, id).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (handler *SqlHandler) Create(arg interface{}) (interface{}, error) {
	result := arg

	if err := handler.Conn.Create(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (handler *SqlHandler) Save(arg interface{}) (interface{}, error) {
	result := arg

	if err := handler.Conn.Save(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (handler *SqlHandler) Delete(arg interface{}) (interface{}, error) {
	result := arg

	if err := handler.Conn.Unscoped().Delete(result).Error; err != nil {
		return nil, err
	}
	
	return result, nil
}

