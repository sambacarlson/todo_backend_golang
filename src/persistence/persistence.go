package persistence

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sambacarlson/todo_backend_golang/src/models"
)

type DataAccess interface {
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

type DBCon struct {
	db sqlx.ExtContext
}

func NewConnection(db *sql.DB) (*DBCon, error) {
	pgdb := sqlx.NewDb(db, "postgres")
	return &DBCon{db: pgdb}, nil
}

var _ DataAccess = &DBCon{}

// user
func (d *DBCon) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	panic("unimplemented")
}

func (d *DBCon) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users := []*models.User{}
	err := sqlx.GetContext(ctx, d.db, users, "SELECT * FROM public.users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d *DBCon) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error) {
	panic("unimplemented")
}

func (d *DBCon) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	panic("unimplemented")
}
func (d *DBCon) GetUserByID(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error) {
	panic("unimplemented")
}

// todo
func (d *DBCon) CreateTodo(ctx context.Context, todo models.Todo) (*models.Todo, error) {
	panic("unimplemented")
}
func (d *DBCon) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	panic("unimplemented")
}
func (d *DBCon) UpdateTodo(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error) {
	panic("unimplemented")
}
func (d *DBCon) DeleteTodo(ctx context.Context, todoID uuid.UUID) error {
	panic("unimplemented")
}
func (d *DBCon) GetTodoByID(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error) {
	panic("unimplemented")
}
func (d *DBCon) GetTodosForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Todo, error) {
	panic("unimplemented")
}

// reminder
func (d *DBCon) CreateReminder(ctx context.Context, reminder models.Reminder) (*models.Reminder, error) {
	panic("unimplemented")
}
func (d *DBCon) GetAllReminders(ctx context.Context) (*models.Reminder, error) {
	panic("unimplemented")
}
func (d *DBCon) UpdateReminder(ctx context.Context, reminderID uuid.UUID, reminder models.Reminder) (*models.Reminder, error) {
	panic("unimplemented")
}
func (d *DBCon) DeleteReminder(ctx context.Context, reminderID uuid.UUID) error {
	panic("unimplemented")
}
func (d *DBCon) GetReminderByID(ctx context.Context, reminderID uuid.UUID) (*models.Reminder, error) {
	panic("unimplemented")
}

// category
func (d *DBCon) CreateCategory(ctx context.Context, category models.Category) (*models.Category, error) {
	panic("unimplemented")
}
func (d *DBCon) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	panic("unimplemented")
}
func (d *DBCon) UpdateCategory(ctx context.Context, categoryID uuid.UUID, category models.Category) (*models.Category, error) {
	panic("unimplemented")
}
func (d *DBCon) DeleteCategory(ctx context.Context, categoryID uuid.UUID) error {
	panic("unimplemented")
}
func (d *DBCon) GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error) {
	panic("unimplemented")
}
