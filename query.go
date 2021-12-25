package Query

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
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

const pageSize = 20

func init()  {
	var err error
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "localhost", 3306, "GOLANG")
	Dbs, err = gorm.Open("mysql", dsn)
	if err != nil{
		panic(err.Error())
	}
	Dbs.AutoMigrate(&items{})
}

func All(c *gin.Context) {
	var queryData []items
	Dbs.Find(&queryData)
	var total =len(queryData)
	pageNum := total / pageSize
	page,_ := strconv.Atoi(c.Param("page"))
	if page - 1 > pageNum{
		c.JSON(http.StatusNotFound,gin.H{
			"status": http.StatusNotFound,
			"msg": "no",
		})
		return
	}
	if page - 1 == pageNum{
		limit := total - ((page - 1) * pageSize)
		if err := Dbs.Limit(limit).Offset((page - 1) * pageSize).Find(&queryData).Error;err!= nil{
			c.JSON(http.StatusUnauthorized,gin.H{
				"err": err.Error(),
			})
		}
	}else {
		if err := Dbs.Limit(pageSize).Offset((page - 1) * pageSize).Find(&queryData).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
		}
	}
	for _, data := range queryData {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": gin.H{
				"item": gin.H{
					"id":         data.ID,
					"title":      data.Title,
					"content":    data.Content,
					"view":       data.View,
					"status":     data.Status,
					"created_At": data.CreatedAt,
					"start_time": data.UpdatedAt,
					"end_time":   0,
				},
				"total": len(queryData),
			},
			"msg":   "ok",
			"error": "",
		})
	}
}

func Status(c *gin.Context)  {
	var queryData []items
	status := c.Param("status")
	Dbs.Find(&queryData,"status = ?", status)
	var total =len(queryData)
	pageNum := total / pageSize
	page,_ := strconv.Atoi(c.Param("page"))
	if page - 1 > pageNum{
		c.JSON(http.StatusNotFound,gin.H{
			"status": http.StatusNotFound,
			"msg": "no",
		})
		return
	}
	if page - 1 == pageNum{
		limit := total - ((page - 1) * pageSize)
		if err := Dbs.Limit(limit).Offset((page - 1) * pageSize).Find(&queryData).Error;err!= nil{
			c.JSON(http.StatusUnauthorized,gin.H{
				"err": err.Error(),
			})
		}
	}else {
		if err := Dbs.Limit(pageSize).Offset((page - 1) * pageSize).Find(&queryData).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
		}
	}
	for _, data := range queryData {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": gin.H{
				"item": gin.H{
					"id":         data.ID,
					"title":      data.Title,
					"content":    data.Content,
					"view":       data.View,
					"status":     data.Status,
					"created_At": data.CreatedAt,
					"start_time": data.UpdatedAt,
					"end_time":   0,
				},
				"total": len(queryData),
			},
			"msg":   "ok",
			"error": "",
		})
	}
}

func KeyWord(c *gin.Context)  {
	var queryData []items
	keyWords := c.Param("keyword")
	Dbs.Where("content LIKE ?", "%"+keyWords+"%").Find(&queryData)
	var total =len(queryData)
	pageNum := total / pageSize
	page,_ := strconv.Atoi(c.Param("page"))
	if page - 1 > pageNum{
		c.JSON(http.StatusNotFound,gin.H{
			"status": http.StatusNotFound,
			"msg": "no",
		})
		return
	}
	if page - 1 == pageNum{
		limit := total - ((page - 1) * pageSize)
		if err := Dbs.Limit(limit).Offset((page - 1) * pageSize).Where("content LIKE ?", "%"+keyWords+"%").Find(&queryData).Error;err!= nil{
			c.JSON(http.StatusUnauthorized,gin.H{
				"err": err.Error(),
			})
		}
	}else {
		if err := Dbs.Limit(pageSize).Offset((page - 1) * pageSize).Where("content LIKE ?", "%"+keyWords+"%").Find(&queryData).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
		}
	}
	for _, data := range queryData {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": gin.H{
				"item": gin.H{
					"id":         data.ID,
					"title":      data.Title,
					"content":    data.Content,
					"view":       data.View,
					"status":     data.Status,
					"created_At": data.CreatedAt,
					"start_time": data.UpdatedAt,
					"end_time":   0,
				},
				"total": len(queryData),
			},
			"msg":   "ok",
			"error": "",
		})
	}
}