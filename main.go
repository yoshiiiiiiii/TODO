package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	_ "net/http"
	"strconv"
	"time"
)

type Plan struct {
	gorm.Model
	Plan string `validate:"required"`
	Detail string

}

func db_init() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}

	db.AutoMigrate(&Plan{})
	db.Close()
}
func create(plan string, detail string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}
	db.Create(&Plan{Plan: plan, Detail: detail})
	db.Close()
}
func get_all() []Plan {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}
	var plans []Plan
	db.Find(&plans)
	db.Close()
	return plans
}

func delete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database\n")
	}
	var plan Plan
	db.First(&plan, id)
	db.Delete(&plan)
	db.Close()

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	db_init()
	r.GET("/", func(c *gin.Context) {
		plans := get_all()
		t := time.Now()
		const layout = "2006年01月02日"
		c.HTML(200, "index.tmpl", gin.H{
			"plans": plans,
			"time": t.Format(layout),
		})
	})

	r.GET("/delete/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		delete(id)
		c.Redirect(302, "/")
	})

	r.POST("/new", func(c *gin.Context) {
		plan := c.PostForm("plan")
		detail := c.PostForm("detail")
		create(plan, detail)
		c.Redirect(302, "/")
	})
	r.Run()
}