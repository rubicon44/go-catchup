package main

import (
	"app/config"
	"app/infra"
	"app/interface/handler"
	"app/usecase"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	taskRepository := infra.NewTaskRepository(config.NewDB())
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	r := gin.Default()
	handler.InitRouting(r, taskHandler)
	r.Run(":8080")
}
