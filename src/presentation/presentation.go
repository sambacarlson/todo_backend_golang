package presentation

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	lg "github.com/sambacarlson/todo_backend_golang/src/logic"
)

func Route(BASE_URL string, lgc lg.LogicImpl) error {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, time.Now().String())
	})

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "status: ok")
	})

	// user enpoints
	router.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/users", func(ctx *gin.Context) {
		users, err := lgc.GetAllUsers(ctx)
		if err != nil {
			log.Printf("---> error is not nil: %v", err)
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, users)
	})
	router.PATCH("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.DELETE("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})

	// todo enpoints
	router.POST("/todos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/todos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.PATCH("/todos/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.DELETE("/todos/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/todos/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/todos/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})

	// reminder endpoints
	router.POST("/reminders", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/reminders", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.PATCH("/reminders/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.DELETE("/reminders/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/reminders/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})

	// category endpoints
	router.POST("/categories", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/categories", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.PATCH("/categories/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.DELETE("/categories/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})
	router.GET("/categories/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})

	// return
	return router.Run(BASE_URL)
}
