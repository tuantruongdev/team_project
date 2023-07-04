package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"team_project/controllers"
)

const (
	dsn  = "root:123@tcp(127.0.0.1:3306)/cocatech?charset=utf8mb4&parseTime=True&loc=Local"
	port = 801
)

func main() {
	router := gin.Default()
	//router.Static("/", "./statics2")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}
	log.Println("Connected to MySQL:", db)
	router.Group("/api/feedback").POST("/", controllers.PostFeedback(db))
	router.POST("/api/point", controllers.UpdatePoints(db))
	router.POST("/api/point/device", controllers.GetPoints(db))
	router.POST("/api/point/leaderboard", controllers.GetLeaderboard(db))
	router.StaticFS("/", http.Dir("static"))
	router.Run(":" + strconv.Itoa(port))
}
