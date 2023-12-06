package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lipaysamart.com/build-bookstore-crud/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {

	dsn := "root:YXq8e2w8BnatCK@tcp(172.16.10.99:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动迁移表数据
	db.AutoMigrate(&models.Book{})

	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	r := gin.Default()

	r.POST("/book/", func(c *gin.Context) {

		var newBook models.Book
		// 将查询绑定到结构体 models.Book
		err := c.ShouldBindJSON(&newBook)
		if err != nil {
			c.JSON(200, gin.H{
				"data": gin.H{},
				"code": 404,
				"msg":  "数据绑定失败",
			})
		} else {
			db.Create(&newBook)
			c.JSON(200, gin.H{
				"data": newBook,
				"code": 200,
				"msg":  "数据绑定成功",
			})
		}
	})

	r.GET("/book/", func(c *gin.Context) {

		var getAllBooks []models.Book
		// 查找所有书籍
		db.Find(&getAllBooks)

		if getAllBooks == nil {
			c.JSON(200, gin.H{
				"data": gin.H{},
				"code": 404,
				"msg":  "数据查询失败",
			})
		} else {
			c.JSON(200, gin.H{
				"data": gin.H{
					"dataList": getAllBooks,
				},
				"code": 200,
				"msg":  "已返回全部数据",
			})
		}
	})

	r.GET("/book/:id", func(c *gin.Context) {

		var getBook []models.Book

		id := c.Param("id")
		db.Where("ID = ?", id).Find(&getBook)
		if len(getBook) == 0 {
			c.JSON(200, gin.H{
				"data": "",
				"code": 404,
				"msg":  "数据查询失败",
			})
		} else {
			c.JSON(200, gin.H{
				"data": getBook,
				"code": "200",
				"msg":  "数据查询成功",
			})
		}
	})

	r.PUT("/book/:id", func(c *gin.Context) {

		var updateBook models.Book

		id := c.Param("id")
		err := db.Where("ID = ?", id).Find(&updateBook)
		if err != nil {
			c.JSON(200, gin.H{
				"data": "",
				"code": 404,
				"msg":  "此用户ID不存在",
			})
		} else {
			err := c.ShouldBindJSON(&updateBook)
			if err != nil {
				c.JSON(200, gin.H{
					"data": "",
					"code": 404,
					"msg":  "书籍更新失败",
				})
			} else {
				db.Where("id = ?", id).Updates(&updateBook)
				c.JSON(200, gin.H{
					"data": updateBook,
					"code": 200,
					"msg":  "书籍更新成功",
				})
			}
		}
	})
	// 删除一本书籍
	r.DELETE("/book/:id", func(c *gin.Context) {

		var deleteBook []models.Book

		id := c.Param("id")
		db.Where("ID = ?", id).Find(&deleteBook)
		if len(deleteBook) == 0 {
			c.String(404, "此用户 ID 不存在")
		} else {
			db.Delete(&deleteBook)
			c.String(200, "成功删除一本书")
		}
	})

	// 监听 8090 端口
	err = r.Run(":8090")
	if err != nil {
		log.Fatal(err)
	}
}
