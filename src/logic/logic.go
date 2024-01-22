package logic

import (
	"context"

	"github.com/google/uuid"
	"github.com/sambacarlson/todo_backend_golang/src/models"
)

type Logic interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUserByID(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error)

	CreateTodo(ctx context.Context, todo models.Todo) (*models.Todo, error)
	GetAllTodos(ctx context.Context) ([]*models.Todo, error)
	UpdateTodo(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error)
	DeleteTodo(ctx context.Context, todoID uuid.UUID) error
	GetTodoByID(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error)
	GetTodosForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Todo, error)

	CreateReminder(ctx context.Context, reminder models.Reminder) (*models.Reminder, error)
	GetAllReminders(ctx context.Context) (*models.Reminder, error)
	UpdateReminder(ctx context.Context, reminderID uuid.UUID, reminder models.Reminder) (*models.Reminder, error)
	DeleteReminder(ctx context.Context, reminderID uuid.UUID) error
	GetReminderByID(ctx context.Context, reminderID uuid.UUID) (*models.Reminder, error)

	CreateCategory(ctx context.Context, category models.Category) (*models.Category, error)
	GetAllCategories(ctx context.Context) ([]*models.Category, error)
	UpdateCategory(ctx context.Context, categoryID uuid.UUID, category models.Category) (*models.Category, error)
	DeleteCategory(ctx context.Context, categoryID uuid.UUID) error
	GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error)
}
