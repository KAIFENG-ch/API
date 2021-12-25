package Update

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
	var updateData items
	updateID := c.PostForm("ID")
	Dbs.First(&updateData, updateID)
	if updateData.Status == false{
		Dbs.Model(&updateData).Update("status",true)
		c.JSON(http.StatusOK,gin.H{
			"status": http.StatusOK,
			"data": gin.H{
				"item":gin.H{
					"id": updateData.ID,
					"title": updateData.Title,
					"content": updateData.Content,
					"view": updateData.View,
					"status": updateData.Status,
					"created_At": updateData.CreatedAt,
					"start_time": time.Now(),
					"end_time": 0,
				},
				"total": 1,
			},
			"msg": "ok",
			"error": "",
		})
	}else {
		Dbs.Model(&updateData).Update(false)
		c.JSON(http.StatusOK,gin.H{
			"status": http.StatusOK,
			"data": gin.H{
				"item":gin.H{
					"id": updateData.ID,
					"title": updateData.Title,
					"content": updateData.Content,
					"view": updateData.View,
					"status": updateData.Status,
					"created_At": updateData.CreatedAt,
					"start_time": time.Now(),
					"end_time": 0,
				},
				"total": 1,
			},
			"msg": "ok",
			"error": "",
		})
	}
}

func All(c *gin.Context)  {
	var updateData []items
	status := c.Param("status")
	Dbs.Find(&updateData,"status = ?", status)
	if len(updateData) == 0{
		c.JSON(http.StatusNotFound,gin.H{
			"status": http.StatusNotFound,
			"msg": "no",
		})
		return
	}
	for _, data := range updateData{
		if data.Status == false{
			Dbs.Model(&data).Update("status",true)
		}else {
			Dbs.Model(&data).Update("status",false)
		}
		c.JSON(http.StatusOK,gin.H{
			"status": http.StatusOK,
			"data": gin.H{
				"item": gin.H{
					"id":         data.ID,
					"title":      data.Title,
					"content":    data.Content,
					"view":       data.View,
					"status":     data.Status,
					"created_At": data.CreatedAt,
					"start_time": time.Now(),
					"end_time":   0,
				},
				"total": len(updateData),
			},
			"msg":   "ok",
			"error": "",
		})
	}
}
