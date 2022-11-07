package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/handlers"
	"github.com/OmarTariq612/codersquare-go/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if err := datastore.CloseDB(); err != nil {
			log.Println(err)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	env := os.Getenv("ENV")
	gin.SetMode(env)

	r := gin.Default()
	// r := gin.New()
	// r.Use(gin.Logger(), middlewares.ErrorMiddleware())

	r.Use(cors.Default())

	// healthz
	r.GET("/api/v1/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "OK"}) })

	r.Use(middlewares.ParseJWTMiddleware(datastore.DB))

	// users
	{
		userHandler := handlers.NewUsersHandler(datastore.DB)
		r.POST("/api/v1/signup", userHandler.Signup)                                   // signup
		r.POST("/api/v1/signin", userHandler.Signin)                                   // signin
		r.GET("/api/v1/users/:id", userHandler.GetUser)                                // get by id
		r.GET("/api/v1/users", middlewares.AuthMiddleware, userHandler.GetCurrentUser) // get by id
	}

	// posts
	{
		postsHandler := handlers.NewPostsHandler(datastore.DB)
		r.GET("/api/v1/posts", postsHandler.List)                                      // list posts
		r.GET("/api/v1/posts/:id", postsHandler.Get)                                   // get post
		r.POST("/api/v1/posts", middlewares.AuthMiddleware, postsHandler.Create)       // create post
		r.DELETE("/api/v1/posts/:id", middlewares.AuthMiddleware, postsHandler.Delete) // delete post
	}

	// likes
	{
		likesHandler := handlers.NewLikesHandler(datastore.DB)
		r.GET("/api/v1/likes/:post_id", likesHandler.List)                                  // list likes  (for a specific post)
		r.POST("/api/v1/likes/:post_id", middlewares.AuthMiddleware, likesHandler.Create)   // create like
		r.DELETE("/api/v1/likes/:post_id", middlewares.AuthMiddleware, likesHandler.Delete) // delete like
	}

	// comments
	{
		commentsHandler := handlers.NewCommentsHandler(datastore.DB)
		r.GET("/api/v1/comments/:post_id", commentsHandler.List)
		r.GET("/api/v1/comments/:post_id/count", commentsHandler.Count)
		r.POST("/api/v1/comments/:post_id", middlewares.AuthMiddleware, commentsHandler.Create)
		r.DELETE("/api/v1/comments/:id", middlewares.AuthMiddleware, commentsHandler.Delete)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil || port == 0 {
		port = 5555
	}

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Println(err)
	}
}
