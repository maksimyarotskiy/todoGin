package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inuhu/todo-app/pkg/db"
	"github.com/inuhu/todo-app/pkg/todo"
)

func main() {
	db.Init()

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/todos", todo.RenderTodos)
	//r.GET("/todos", todo.GetTodos)
	r.GET("/todo/:id", todo.GetTodo)
	r.POST("/todos", todo.CreateTodo)
	r.POST("/delete/:id", todo.DeleteTodo)
	r.Run(":8080")
}
