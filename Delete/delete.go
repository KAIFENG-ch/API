package Delete

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

func One(c *gin.Context)  {
	var deleteData items
	deleteId := c.PostForm("id")
	Dbs.First(&deleteData, deleteId)
	if deleteData.ID == 0{
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"message" : "no",
		})
		return
	}
	Dbs.Delete(&deleteData)
	c.JSON(http.StatusOK,gin.H{
		"status":  http.StatusOK,
		"data": gin.H{
			"item": gin.H{
				"id": deleteData.ID,
				"title" : deleteData.Title,
				"content" : deleteData.Content,
				"view": deleteData.View,
				"status": deleteData.Status,
				"created_time": deleteData.CreatedAt,
				"start_time": deleteData.UpdatedAt,
				"end_time": time.Now(),
			},
			"total": 1,
		},
		"msg" : "ok",
		"error" : "",
	})
}

func Status(c *gin.Context)  {
	var delData []items
	delStatus := c.PostForm("status")
	Dbs.Find(&delData,"status = ?", delStatus)
	if len(delData) <= 0{
		c.JSON(http.StatusOK,gin.H{
			"status": http.StatusNotFound,
			"message" : "no data",
		})
		return
	}
	for _,del := range delData{
		Dbs.Delete(&del)
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"data": gin.H{
				"items": gin.H{
					"id":         del.ID,
					"title":      del.Title,
					"content":    del.Content,
					"view":       del.View,
					"status":     del.Status,
					"created_At": del.CreatedAt,
					"start_Time": del.UpdatedAt,
					"end_Time":   time.Now(),
				},
				"total" : len(delData),
			},
			"msg": "ok",
			"error":"",
		})
	}
}

func All(c *gin.Context)  {
	var delItems []items
	Dbs.Find(&delItems)
	if len(delItems) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNotFound,
			"message": "no data",
		})
		return
	}
	for _,del := range delItems{
		Dbs.Delete(&del)
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"data": gin.H{
				"items": gin.H{
					"id":         del.ID,
					"title":      del.Title,
					"content":    del.Content,
					"view":       del.View,
					"status":     del.Status,
					"created_At": del.CreatedAt,
					"start_Time": del.UpdatedAt,
					"end_Time":   time.Now(),
				},
				"total" : len(delItems),
			},
			"msg": "ok",
			"error":"",
		})
	}
}
