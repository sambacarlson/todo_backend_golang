package logic

import (
	"context"

	"github.com/google/uuid"
	"github.com/sambacarlson/todo_backend_golang/src/models"
	"github.com/sambacarlson/todo_backend_golang/src/persistence"
)

type Logic interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
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

type LogicImpl struct {
	repo persistence.DataAccess
}

var _ Logic = &LogicImpl{}

func NewLogic(repo persistence.DataAccess) (*LogicImpl, error) {
	return &LogicImpl{
		repo: repo,
	}, nil
}

// user
func (l *LogicImpl) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	panic("unimplemented")
}
func (l *LogicImpl) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users, err := l.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (l *LogicImpl) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error) {
	panic("unimplemented")
}
func (l *LogicImpl) DeleteUser(ctx context.Context, userID uuid.UUID) error { panic("unimplemented") }
func (l *LogicImpl) GetUserByID(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error) {
	panic("unimplemented")
}

// todo
func (l *LogicImpl) CreateTodo(ctx context.Context, todo models.Todo) (*models.Todo, error) {
	panic("unimplemented")
}
func (l *LogicImpl) GetAllTodos(ctx context.Context) ([]*models.Todo, error) { panic("unimplemented") }
func (l *LogicImpl) UpdateTodo(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error) {
	panic("unimplemented")
}
func (l *LogicImpl) DeleteTodo(ctx context.Context, todoID uuid.UUID) error { panic("unimplemented") }
func (l *LogicImpl) GetTodoByID(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error) {
	panic("unimplemented")
}
func (l *LogicImpl) GetTodosForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Todo, error) {
	panic("unimplemented")
}

// reminder
func (l *LogicImpl) CreateReminder(ctx context.Context, reminder models.Reminder) (*models.Reminder, error) {
	panic("unimplemented")
}
func (l *LogicImpl) GetAllReminders(ctx context.Context) (*models.Reminder, error) {
	panic("unimplemented")
}
func (l *LogicImpl) UpdateReminder(ctx context.Context, reminderID uuid.UUID, reminder models.Reminder) (*models.Reminder, error) {
	panic("unimplemented")
}
func (l *LogicImpl) DeleteReminder(ctx context.Context, reminderID uuid.UUID) error {
	panic("unimplemented")
}
func (l *LogicImpl) GetReminderByID(ctx context.Context, reminderID uuid.UUID) (*models.Reminder, error) {
	panic("unimplemented")
}

// category
func (l *LogicImpl) CreateCategory(ctx context.Context, category models.Category) (*models.Category, error) {
	panic("unimplemented")
}
func (l *LogicImpl) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	panic("unimplemented")
}
func (l *LogicImpl) UpdateCategory(ctx context.Context, categoryID uuid.UUID, category models.Category) (*models.Category, error) {
	panic("unimplemented")
}
func (l *LogicImpl) DeleteCategory(ctx context.Context, categoryID uuid.UUID) error {
	panic("unimplemented")
}
func (l *LogicImpl) GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error) {
	panic("unimplemented")
}
