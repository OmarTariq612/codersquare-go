package main

import (
	"fmt"

	"github.com/OmarTariq612/codersquare-go/handlers"
	"github.com/OmarTariq612/codersquare-go/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// r := gin.Default()

	r := gin.New()
	r.Use(gin.Logger(), middlewares.ErrorMiddleware())

	r.GET("/panic", func(ctx *gin.Context) {
		panic("error occured")
	})

	// users
	{
		r.POST("/api/v1/signup", handlers.SignupHandler) // signup
		r.POST("/api/v1/signin", handlers.SigninHandler) // signin
	}

	r.Use(middlewares.AuthMiddleware())

	// posts
	{
		r.GET("/api/v1/posts", handlers.ListPostsHandler)   // list posts
		r.POST("/api/v1/posts", handlers.CreatePostHandler) // create post
	}

	// r.POST("/testing/:name", func(ctx *gin.Context) {
	// 	type request struct {
	// 		Name string `uri:"name" json:"name" binding:"required"`
	// 		Age  int    `json:"age" binding:"required"`
	// 	}

	// 	req := &request{}

	// 	// name
	// 	if err := ctx.ShouldBindUri(req); err != nil {
	// 		log.Println(err)
	// 	}

	// 	req.Name = ""

	// 	// age
	// 	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
	// 		// log.Println(err)
	// 		panic(err)
	// 	}

	// 	ctx.JSON(http.StatusOK, gin.H{"user": req})
	// })

	if err := r.Run(":5555"); err != nil {
		fmt.Println(err)
	}
}
