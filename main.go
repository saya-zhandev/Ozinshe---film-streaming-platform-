package main

import (
	"Ozinshe---film-streaming-platform-/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/saya-zhandev/Ozinshe---film-streaming-platform-/tree/main/handlers"
)

func main() {
	r := gin.Default()

	corsConfig := cors.Config{ //cors: lets the browser know that it is safe to access the info from this
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}
	r.Use(cors.New(corsConfig))

	moviesHandler := handlers.NewMovieHandlers()

	r.GET("/movies/:id", moviesHandler.FindById)
	r.GET("/movies", moviesHandler.FindAll)
	r.POST("/movies", moviesHandler.Create)
	r.PUT("/movies/:id", moviesHandler.Update)
	r.DELETE("/movies/:id", moviesHandler.Delete)

	r.Run(":8081")
}
