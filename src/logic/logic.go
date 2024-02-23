package logic

import (
	"context"

	"github.com/google/uuid"
	"github.com/sambacarlson/todo_backend_golang/src/models"
	"github.com/sambacarlson/todo_backend_golang/src/persistence"
)

type Logic interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)

	CreateTodo(ctx context.Context, todo models.Todo) (*models.Todo, error)
	GetAllTodos(ctx context.Context) ([]*models.Todo, error)
	UpdateTodo(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error)
	DeleteTodo(ctx context.Context, todoID uuid.UUID) error
	GetTodoByID(ctx context.Context, todoID uuid.UUID) (*models.Todo, error)
	GetTodosForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Todo, error)

	CreateReminder(ctx context.Context, reminder models.Reminder) (*models.Reminder, error)
	GetAllReminders(ctx context.Context) ([]*models.Reminder, error)
	UpdateReminder(ctx context.Context, reminderID uuid.UUID, reminder models.Reminder) (*models.Reminder, error)
	DeleteReminder(ctx context.Context, reminderID uuid.UUID) error
	GetReminderByID(ctx context.Context, reminderID uuid.UUID) (*models.Reminder, error)
	GetRemindersForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Reminder, error)

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
	createdUser, err := l.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
func (l *LogicImpl) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users, err := l.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (l *LogicImpl) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error) {
	userUpdated, err := l.repo.UpdateUser(ctx, userID, user)
	if err != nil {
		return nil, err
	}
	return userUpdated, nil
}
func (l *LogicImpl) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	err := l.repo.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
func (l *LogicImpl) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	selectedUser, err := l.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return selectedUser, nil
}

// todo
func (l *LogicImpl) CreateTodo(ctx context.Context, todo models.Todo) (*models.Todo, error) {
	createdTodo, err := l.repo.CreateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}
	return createdTodo, nil
}
func (l *LogicImpl) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	todos, err := l.repo.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
func (l *LogicImpl) UpdateTodo(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error) {
	updatedTodo, err := l.repo.UpdateTodo(ctx, todoID, todo)
	if err != nil {
		return nil, err
	}
	return updatedTodo, nil
}
func (l *LogicImpl) DeleteTodo(ctx context.Context, todoID uuid.UUID) error {
	err := l.repo.DeleteTodo(ctx, todoID)
	if err != nil {
		return err
	}
	return nil
}
func (l *LogicImpl) GetTodoByID(ctx context.Context, todoID uuid.UUID) (*models.Todo, error) {
	selectedTodo, err := l.repo.GetTodoByID(ctx, todoID)
	if err != nil {
		return nil, err
	}
	return selectedTodo, nil
}
func (l *LogicImpl) GetTodosForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Todo, error) {
	selectedTodos, err := l.repo.GetTodosForOneUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return selectedTodos, nil
}

// reminder
func (l *LogicImpl) CreateReminder(ctx context.Context, reminder models.Reminder) (*models.Reminder, error) {
	createdReminder, err := l.repo.CreateReminder(ctx, reminder)
	if err != nil {
		return nil, err
	}
	return createdReminder, nil
}
func (l *LogicImpl) GetAllReminders(ctx context.Context) ([]*models.Reminder, error) {
	reminders, err := l.repo.GetAllReminders(ctx)
	if err != nil {
		return nil, err
	}
	return reminders, nil
}
func (l *LogicImpl) UpdateReminder(ctx context.Context, reminderID uuid.UUID, reminder models.Reminder) (*models.Reminder, error) {
	updatedReminder, err := l.repo.UpdateReminder(ctx, reminderID, reminder)
	if err != nil {
		return nil, err
	}
	return updatedReminder, nil
}
func (l *LogicImpl) DeleteReminder(ctx context.Context, reminderID uuid.UUID) error {
	err := l.repo.DeleteReminder(ctx, reminderID)
	if err != nil {
		return err
	}
	return nil
}
func (l *LogicImpl) GetReminderByID(ctx context.Context, reminderID uuid.UUID) (*models.Reminder, error) {
	selectedReminder, err := l.repo.GetReminderByID(ctx, reminderID)
	if err != nil {
		return nil, err
	}
	return selectedReminder, nil
}
func (l *LogicImpl) GetRemindersForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Reminder, error) {
	selectedReminders, err := l.repo.GetAllReminders(ctx)
	if err != nil {
		return nil, err
	}
	return selectedReminders, nil
}

// category
func (l *LogicImpl) CreateCategory(ctx context.Context, category models.Category) (*models.Category, error) {
	createdCategory, err := l.repo.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}
func (l *LogicImpl) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	categories, err := l.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
func (l *LogicImpl) UpdateCategory(ctx context.Context, categoryID uuid.UUID, category models.Category) (*models.Category, error) {
	updatedCategory, err := l.repo.UpdateCategory(ctx, categoryID, category)
	if err != nil {
		return nil, err
	}
	return updatedCategory, nil
}
func (l *LogicImpl) DeleteCategory(ctx context.Context, categoryID uuid.UUID) error {
	err := l.repo.DeleteCategory(ctx, categoryID)
	if err != nil {
		return err
	}
	return nil
}
func (l *LogicImpl) GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error) {
	selectedCategory, err := l.repo.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	return selectedCategory, nil
}
