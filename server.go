package main

import (
	"fmt"

	"github.com/OmarTariq612/codersquare-go/datastore"
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

	// users
	{
		userHandler := handlers.NewUserHandler(datastore.DB)
		r.POST("/api/v1/signup", userHandler.Signup)                                     // signup
		r.POST("/api/v1/signin", userHandler.Signin)                                     // signin
		r.GET("/api/v1/users/:id", userHandler.GetUser)                                  // get by id
		r.GET("/api/v1/users", middlewares.AuthMiddleware(), userHandler.GetCurrentUser) // get by id
	}

	// r.Use(middlewares.AuthMiddleware())

	// posts
	{
		postsHandler := handlers.NewPostsHandler(datastore.DB)
		r.GET("/api/v1/posts", postsHandler.List)                                  // list posts
		r.POST("/api/v1/posts", middlewares.AuthMiddleware(), postsHandler.Create) // create post
		r.GET("/api/v1/posts/:id", middlewares.AuthMiddleware(), postsHandler.Get) // get post
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
