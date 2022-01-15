package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"test2/Add"
	"test2/Delete"
	"test2/Query"
	"test2/Register"
	"test2/Update"
)

var Dbs *gorm.DB

type items struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	View int `json:"view"`
	Status bool `json:"status"`
}

func (items)TableName() string{
	return "todos"
}

func init()  {
	var err error
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "localhost", 3306, "GOLANG")
	Dbs, err = gorm.Open("mysql", dsn)
	if err != nil{
		panic(err.Error())
	}
	Dbs.AutoMigrate(&items{})
}

func main()  {
	r := gin.Default()
	v1 := r.Group("/api/v1/todo")
	{
		v1.POST("/", Add.Add)
		v1.DELETE("/id", Delete.One)
		v1.DELETE("/status", Delete.Status)
		v1.DELETE("/all", Delete.All)
		v1.PUT("/one", Update.One)
		v1.PUT("/all", Update.All)
		v1.GET("/all",Query.All)
		v1.GET("/status",Query.Status)
		v1.GET("/content",Query.KeyWord)
		v1.POST("/sign",Register.Sign)
		v1.POST("/welcome",Register.Welcome)
		v1.POST("refresh",Register.Refresh)
	}
	err := r.Run(":8000")
	if err != nil	 {
		return
	}
}
