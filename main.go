package main

import "github.com/gin-gonic/gin"

type Todo struct {
    ID   int    `json:"id"`
    Task string `json:"task"`
}

var todos = []Todo{
    {ID: 1, Task: "Learn Go"},
    {ID: 2, Task: "Build API"},
}

func main() {
    r := gin.Default()

    r.GET("/todos", func(c *gin.Context) {
        c.JSON(200, todos)
    })

    r.POST("/todos", func(c *gin.Context) {
        var newTodo Todo
        if err := c.BindJSON(&newTodo); err != nil {
            return
        }
        todos = append(todos, newTodo)
        c.JSON(200, newTodo)
    })

    r.Run(":8080")
}
