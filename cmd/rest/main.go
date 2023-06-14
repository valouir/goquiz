package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/valouir/goquiz/packages/controllers"
)

func main() {

	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())

	ginEngine.GET("/health", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, true)
	})

	ginEngine.GET("/questions", controllers.GetQuestions)
	ginEngine.POST("/answers", controllers.SubmitAnswers)

	// Start the server
	err := ginEngine.Run(fmt.Sprintf(":%v", 6379))
	if err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
