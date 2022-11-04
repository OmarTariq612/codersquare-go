package main

import (
	"fmt"
	"net/http"

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

	r := gin.Default()

	// r := gin.New()
	// r.Use(gin.Logger(), middlewares.ErrorMiddleware())

	// healthz
	r.GET("/api/v1/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "OK"}) })

	r.Use(middlewares.ParseJWTMiddlware)

	// users
	{
		userHandler := handlers.NewUsersHandler(datastore.DB)
		r.POST("/api/v1/signup", userHandler.Signup)                                   // signup
		r.POST("/api/v1/signin", userHandler.Signin)                                   // signin
		r.GET("/api/v1/users/:id", userHandler.GetUser)                                // get by id
		r.GET("/api/v1/users", middlewares.AuthMiddleware, userHandler.GetCurrentUser) // get by id
	}

	// r.Use(middlewares.AuthMiddleware())

	// posts
	{
		postsHandler := handlers.NewPostsHandler(datastore.DB)
		r.GET("/api/v1/posts", postsHandler.List)                                      // list posts
		r.POST("/api/v1/posts", middlewares.AuthMiddleware, postsHandler.Create)       // create post
		r.GET("/api/v1/posts/:id", middlewares.AuthMiddleware, postsHandler.Get)       // get post
		r.DELETE("/api/v1/posts/:id", middlewares.AuthMiddleware, postsHandler.Delete) // delete post
	}

	// likes
	{
		likesHandler := handlers.NewLikesHandler(datastore.DB)
		r.GET("/api/v1/likes/:post_id", likesHandler.List)                                  // list likes  (for a specific post)
		r.POST("/api/v1/likes/:post_id", middlewares.AuthMiddleware, likesHandler.Create)   // create like /////////////////////
		r.DELETE("/api/v1/likes/:post_id", middlewares.AuthMiddleware, likesHandler.Delete) // delete like /////////////////////
	}

	// comments
	{
		commentsHandler := handlers.NewCommentsHandler(datastore.DB)
		r.GET("/api/v1/comments/:post_id", commentsHandler.List)
		r.GET("/api/v1/comments/:post_id/count", commentsHandler.Count)
		r.POST("/api/v1/comments/:post_id", middlewares.AuthMiddleware, commentsHandler.Create)
		r.DELETE("/api/v1/comments/:id", middlewares.AuthMiddleware, commentsHandler.Delete)
	}

	if err := r.Run(":5555"); err != nil {
		fmt.Println(err)
	}
}
