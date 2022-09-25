package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:@tcp(127.0.0.1:3306)/developer?charset=utf8mb4&parseTime=True&loc=Local"
)

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) statusColName() string {
	return "status"
}
func (ToDoItem) TableName() string { return "todo_items" }

func main() {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)

	router := gin.Default()
	v1 := router.Group("/v1/items")
	{
		v1.GET("", getTasks(db)).
			POST("", createTask(db)).
			GET("/:id", getTasksById(db)).
			DELETE("/:id", deleteTaskById(db)).
			PATCH("/:id", markTaskById(db))
	}

	router.Run()
}

func protector() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "error": "you dont have permission"})
	}
}

func countWithId(db *gorm.DB, count *int64, id int) {
	db.Model(&ToDoItem{}).Where("id =?", id).Count(count)
}

func createTask(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var tempTask = ToDoItem{}
		if err := context.ShouldBind(&tempTask); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//remove blank space
		tempTask.Title = strings.TrimSpace(tempTask.Title)
		if tempTask.Title == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "task title cannot be blank"})
			return
		}

		tempTask.Status = "doing"
		if err := db.Create(&tempTask).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "ok", "data": tempTask})

		//fmt.Println(context.Params.Get("title"))
	}
}

func getTasks(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var taskItem []ToDoItem
		//	if err := db.Find(&taskItem).Error; err != nil {
		if err := db.Where("id > 0").Find(&taskItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"status": "error", "messenger": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "ok", "data": taskItem})
	}
}

func getTasksById(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var taskItem ToDoItem
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"status": "id error"})
			return
		}

		if err := db.Where("id = ?", id).First(&taskItem).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"status": "no id match"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "ok", "data": taskItem})
	}
}

func deleteTaskById(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var count int64 = 0

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "error": err})
			return
		}
		countWithId(db, &count, id)
		if count < 1 {
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "no record match"})
			return
		}
		if err := db.Delete(&ToDoItem{}, id).Error; err != nil {
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "error": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func markTaskById(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var task ToDoItem
		var count int64 = 0
		context.ShouldBind(&task)
		fmt.Println(task)
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
			return
		}

		if len(task.Status) == 0 || (task.Status != "Doing" && task.Status != "Finished") {
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "no acceptable query"})
			return
		}

		countWithId(db, &count, id)
		if count < 1 {
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "no record match"})
			return
		}
		db.Model(&ToDoItem{}).Where("id =?", id).Update(ToDoItem{}.statusColName(), task.Status)
		context.JSON(http.StatusAccepted, gin.H{"status": "ok", "message": "updated"})
		return
	}
}
