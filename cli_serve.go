package main

import (
	"github.com/gin-gonic/gin"
)

func Serve(bind string) (err error) {
	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/check"),
		gin.Recovery(),
	)

	router.GET("/health/check", healthCheck)

	// start server
	log.Info("Start listening on " + bind)
	err = router.Run(bind)
	if err != nil {
		return
	}
	return nil
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"result":  "ok",
		"version": appVersion,
	})
}
