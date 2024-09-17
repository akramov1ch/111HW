package main

import (
    "111HW/config"
    "111HW/controllers"
    "111HW/middleware"
    "111HW/models"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()

    models.ConnectDatabase()

    r := gin.Default()

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    auth.POST("/posts", controllers.CreatePost)
    auth.GET("/posts", controllers.GetPosts)
    auth.PUT("/posts/:id", controllers.UpdatePost)
    auth.DELETE("/posts/:id", controllers.DeletePost)

    r.Run(":8080")
}
