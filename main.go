package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"task-manager/model"
	"task-manager/router"
	S "task-manager/services"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.ConnectDB()
	model.DB.AutoMigrate(&model.User{}, &model.Task{})

	delay, _ := strconv.Atoi(os.Getenv("TASK_AUTO_COMPLETE_MINUTES"))
	go S.StartAutoCompleteWorker(time.Duration(delay) * time.Minute)

	r := router.SetupRouter()
	r.Run(":8080")
}
