package Add

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)


var Dbs *gorm.DB

type items struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	View int `json:"view"`
	Status bool `json:"status"`
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	deletedAt *time.Time
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

func Add(c *gin.Context)  {
	data := items{
		Title: c.PostForm("title"),
		Content: c.PostForm("content"),
		View: 0,
		Status: false,
	}
	Dbs.Save(&data)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"data": gin.H{
				"item": gin.H{
					"id": data.ID,
					"title" : data.Title,
					"content" : data.Content,
					"view": data.View,
					"status": data.Status,
					"created_time": data.CreatedAt,
					"start_time": data.UpdatedAt,
					"end_time": 0,
				},
				"total": 1,
			},
			"msg" : "ok",
			"error" : "",
		},
	)
}