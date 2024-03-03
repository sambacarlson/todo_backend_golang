package presentation

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	lg "github.com/sambacarlson/todo_backend_golang/src/logic"
	"github.com/sambacarlson/todo_backend_golang/src/models"
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
		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		createdUser, err := lgc.CreateUser(ctx, user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, createdUser)
	})
	router.GET("/users", func(ctx *gin.Context) {
		users, err := lgc.GetAllUsers(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, users)
	})
	// router.PATCH("/users/:id", func(ctx *gin.Context) {
	// 	parsedUserID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	var user models.User
	// 	if err := ctx.BindJSON(&user); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	updatedUser, err := lgc.UpdateUser(ctx, parsedUserID, user)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, updatedUser)
	// })
	// router.DELETE("/users/:id", func(ctx *gin.Context) {
	// 	parsedUserID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	err = lgc.DeleteUser(ctx, parsedUserID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, nil)
	// })
	// router.GET("/users/:id", func(ctx *gin.Context) {
	// 	parsedUserID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	user, err := lgc.GetUserByID(ctx, parsedUserID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, user)
	// })

	// // todo enpoints
	// router.POST("/todos", func(ctx *gin.Context) {
	// 	var todo models.Todo
	// 	if err := ctx.BindJSON(&todo); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	createdTodo, err := lgc.CreateTodo(ctx, todo)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusCreated, createdTodo)
	// })
	// router.GET("/todos", func(ctx *gin.Context) {
	// 	todos, err := lgc.GetAllTodos(ctx)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, todos)
	// })
	// router.PATCH("/todos/:id", func(ctx *gin.Context) {
	// 	parsedTodoID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	var todo models.Todo
	// 	if err := ctx.BindJSON(&todo); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	updatedTodo, err := lgc.UpdateTodo(ctx, parsedTodoID, todo)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, updatedTodo)
	// })
	// router.DELETE("/todos/:id", func(ctx *gin.Context) {
	// 	parsedTodoID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	err = lgc.DeleteTodo(ctx, parsedTodoID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, nil)
	// })
	// router.GET("/todos/:id", func(ctx *gin.Context) {
	// 	parsedTodoID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	todo, err := lgc.GetTodoByID(ctx, parsedTodoID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, todo)
	// })
	// router.GET("/todos/users/:id", func(ctx *gin.Context) {
	// 	parsedUserID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	todos, err := lgc.GetTodosForOneUser(ctx, parsedUserID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, todos)
	// })

	// // reminder endpoints
	// router.POST("/reminders", func(ctx *gin.Context) {
	// 	var reminder models.Reminder
	// 	if err := ctx.BindJSON(&reminder); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	createdReminder, err := lgc.CreateReminder(ctx, reminder)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusCreated, createdReminder)
	// })
	// router.GET("/reminders", func(ctx *gin.Context) {
	// 	reminders, err := lgc.GetAllReminders(ctx)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, reminders)
	// })
	// router.PATCH("/reminders/:id", func(ctx *gin.Context) {
	// 	parsedReminderID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	var reminder models.Reminder
	// 	if err := ctx.BindJSON(&reminder); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	updatedReminder, err := lgc.UpdateReminder(ctx, parsedReminderID, reminder)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, updatedReminder)
	// })
	// router.DELETE("/reminders/:id", func(ctx *gin.Context) {
	// 	parsedReminderID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	err = lgc.DeleteReminder(ctx, parsedReminderID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, nil)
	// })
	// router.GET("/reminders/:id", func(ctx *gin.Context) {
	// 	parsedReminderID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	reminder, err := lgc.GetReminderByID(ctx, parsedReminderID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, reminder)
	// })
	// router.GET("/reminders/users/:id", func(ctx *gin.Context) {
	// 	parsedUserID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	reminders, err := lgc.GetRemindersForOneUser(ctx, parsedUserID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, reminders)
	// })

	// // category endpoints
	// router.POST("/categories", func(ctx *gin.Context) {
	// 	var category models.Category
	// 	if err := ctx.BindJSON(&category); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	createdCategory, err := lgc.CreateCategory(ctx, category)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusCreated, createdCategory)
	// })
	// router.GET("/categories", func(ctx *gin.Context) {
	// 	categories, err := lgc.GetAllCategories(ctx)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, categories)
	// })
	// router.PATCH("/categories/:id", func(ctx *gin.Context) {
	// 	parsedCategoryID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	var category models.Category
	// 	if err := ctx.BindJSON(&category); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	updatedCategory, err := lgc.UpdateCategory(ctx, parsedCategoryID, category)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, updatedCategory)
	// })
	// router.DELETE("/categories/:id", func(ctx *gin.Context) {
	// 	parsedCategoryID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	err = lgc.DeleteCategory(ctx, parsedCategoryID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, nil)
	// })
	// router.GET("/categories/:id", func(ctx *gin.Context) {
	// 	parsedCategoryID, err := utils.ParseUUID(ctx.Param("id"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid UUID param: %s", err.Error())})
	// 		return
	// 	}
	// 	category, err := lgc.GetCategoryByID(ctx, parsedCategoryID)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, category)
	// })

	// return
	return router.Run(BASE_URL)
}
