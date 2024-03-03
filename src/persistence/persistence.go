package persistence

import (
	"context"

	"github.com/sambacarlson/todo_backend_golang/src/models"
	"gorm.io/gorm"
)

type DataAccess interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	// UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error)
	// DeleteUser(ctx context.Context, userID uuid.UUID) error
	// GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)

	// CreateTodo(ctx context.Context, todo models.Todo) (*models.Todo, error)
	// GetAllTodos(ctx context.Context) ([]*models.Todo, error)
	// UpdateTodo(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error)
	// DeleteTodo(ctx context.Context, todoID uuid.UUID) error
	// GetTodoByID(ctx context.Context, todoID uuid.UUID) (*models.Todo, error)
	// GetTodosForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Todo, error)

	// CreateReminder(ctx context.Context, reminder models.Reminder) (*models.Reminder, error)
	// GetAllReminders(ctx context.Context) ([]*models.Reminder, error)
	// UpdateReminder(ctx context.Context, reminderID uuid.UUID, reminder models.Reminder) (*models.Reminder, error)
	// DeleteReminder(ctx context.Context, reminderID uuid.UUID) error
	// GetReminderByID(ctx context.Context, reminderID uuid.UUID) (*models.Reminder, error)

	// CreateCategory(ctx context.Context, category models.Category) (*models.Category, error)
	// GetAllCategories(ctx context.Context) ([]*models.Category, error)
	// UpdateCategory(ctx context.Context, categoryID uuid.UUID, category models.Category) (*models.Category, error)
	// DeleteCategory(ctx context.Context, categoryID uuid.UUID) error
	// GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error)
}

type DBCon struct {
	db *gorm.DB
}

func NewConnection(gormDB *gorm.DB) (*DBCon, error) {
	return &DBCon{db: gormDB}, nil
}

// makes sure that DBCon implements DataAcess interface
var _ DataAccess = &DBCon{}

// user
func (d *DBCon) CreateUser(ctx context.Context, user models.User)(*models.User, error) {
	createdUser := models.User{}
	err := d.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func (d *DBCon) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := d.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// func (d *DBCon) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) (*models.User, error) {
// 	var updatedUser models.User
// 	err := sqlx.GetContext(ctx, d.db, &updatedUser, "UPDATE public.users SET username = $1 WHERE id = $2 RETURNING *", user.Username, userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &updatedUser, nil
// }

// func (d *DBCon) DeleteUser(ctx context.Context, userID uuid.UUID) error {
// 	_, err := d.db.ExecContext(ctx, "DELETE FROM public.users WHERE id = $1", userID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (d *DBCon) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
// 	var selectedUser models.User
// 	err := sqlx.GetContext(ctx, d.db, &selectedUser, "SELECT * FROM public.users WHERE id = $1", userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &selectedUser, nil
// }

// // todo
// func (d *DBCon) CreateTodo(ctx context.Context, todo models.Todo) (*models.Todo, error) {
// 	createdTodo := models.Todo{}
// 	err := sqlx.GetContext(ctx, d.db, &createdTodo, "INSERT INTO public.todos (heading, body, user_id, category_id) VALUES ($1, $2, $3, $4) RETURNING *", todo.Heading, todo.Body, todo.UserID, todo.CategoryID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &createdTodo, nil
// }
// func (d *DBCon) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
// 	var todos []*models.Todo
// 	err := sqlx.SelectContext(ctx, d.db, &todos, "SELECT * FROM public.todos")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return todos, nil
// }
// func (d *DBCon) UpdateTodo(ctx context.Context, todoID uuid.UUID, todo models.Todo) (*models.Todo, error) {
// 	var updatedTodo models.Todo
// 	err := sqlx.GetContext(ctx, d.db, &updatedTodo, "UPDATE public.todos SET user_id = $1, heading = $2, body = $3, category_id = $4  WHERE id = $5 RETURNING *", todo.UserID, todo.Heading, todo.Body, todo.CategoryID, todoID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &updatedTodo, nil
// }
// func (d *DBCon) DeleteTodo(ctx context.Context, todoID uuid.UUID) error {
// 	_, err := d.db.ExecContext(ctx, "DELETE FROM public.todos WHERE id = $1", todoID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (d *DBCon) GetTodoByID(ctx context.Context, todoID uuid.UUID) (*models.Todo, error) {
// 	var selectedTodo models.Todo
// 	err := sqlx.GetContext(ctx, d.db, &selectedTodo, "SELECT * FROM public.todos WHERE id = $1", todoID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &selectedTodo, nil
// }
// func (d *DBCon) GetTodosForOneUser(ctx context.Context, userID uuid.UUID) ([]*models.Todo, error) {
// 	var selectedUsers []*models.Todo
// 	err := sqlx.SelectContext(ctx, d.db, &selectedUsers, "SELECT * FROM public.todos WHERE user_id = $1", userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return selectedUsers, nil
// }

// // reminder
// func (d *DBCon) CreateReminder(ctx context.Context, reminder models.Reminder) (*models.Reminder, error) {
// 	createdReminder := models.Reminder{}
// 	err := sqlx.GetContext(ctx, d.db, &createdReminder, "INSERT INTO public.reminders(todo_id, description, remind_time, snooze, snooze_timeout) VALUES($1, $2, $3, $4, $5) RETURNING *", reminder.TodoID, reminder.Description, reminder.RemindTime, reminder.Snooze, reminder.SnoozeTimeout)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &createdReminder, nil
// }
// func (d *DBCon) GetAllReminders(ctx context.Context) ([]*models.Reminder, error) {
// 	var reminders []*models.Reminder
// 	err := sqlx.SelectContext(ctx, d.db, &reminders, "SELECT * FROM public.reminders")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return reminders, nil
// }
// func (d *DBCon) UpdateReminder(ctx context.Context, reminderID uuid.UUID, reminder models.Reminder) (*models.Reminder, error) {
// 	var updatedReminder models.Reminder
// 	err := sqlx.GetContext(ctx, d.db, &updatedReminder, "UPDATE public.reminders SET todo_id = $1, snooze = $2, snooze_timeout = $3, remind_time = $4, description = $5 WHERE id = $6 RETURNING *", reminder.TodoID, reminder.Snooze, reminder.SnoozeTimeout, reminder.RemindTime, reminder.Description, reminderID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &updatedReminder, nil
// }
// func (d *DBCon) DeleteReminder(ctx context.Context, reminderID uuid.UUID) error {
// 	_, err := d.db.ExecContext(ctx, "DELETE FROM public.reminders WHERE id = $1", reminderID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (d *DBCon) GetReminderByID(ctx context.Context, reminderID uuid.UUID) (*models.Reminder, error) {
// 	var selectedReminder models.Reminder
// 	err := sqlx.GetContext(ctx, d.db, &selectedReminder, "SELECT * FROM public.reminders WHERE id = $1", reminderID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &selectedReminder, nil
// }

// // category
// func (d *DBCon) CreateCategory(ctx context.Context, category models.Category) (*models.Category, error) {
// 	createdCategory := models.Category{}
// 	err := sqlx.GetContext(ctx, d.db, &createdCategory, "INSERT INTO public.categories(title, description) VALUES ($1, $2) RETURNING *", category.Title, category.Description)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &createdCategory, nil
// }
// func (d *DBCon) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
// 	var categories []*models.Category
// 	err := sqlx.SelectContext(ctx, d.db, &categories, "SELECT * FROM public.categories")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return categories, nil
// }
// func (d *DBCon) UpdateCategory(ctx context.Context, categoryID uuid.UUID, category models.Category) (*models.Category, error) {
// 	var updatedCategory models.Category
// 	err := sqlx.GetContext(ctx, d.db, &updatedCategory, "UPDATE public.categories SET title = $1, description = $2 WHERE id = $3 RETURNING *", category.Title, category.Description, categoryID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &updatedCategory, nil
// }
// func (d *DBCon) DeleteCategory(ctx context.Context, categoryID uuid.UUID) error {
// 	_, err := d.db.ExecContext(ctx, "DELETE FROM public.categories WHERE id = $1", categoryID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (d *DBCon) GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error) {
// 	var selectedCategory models.Category
// 	err := sqlx.GetContext(ctx, d.db, &selectedCategory, "SELECT * FROM public.categories WHERE id = $1", categoryID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &selectedCategory, nil
// }
