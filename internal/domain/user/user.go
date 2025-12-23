package user

import (
	"todo-list-api-roadmap-sh/internal/domain/shared"
	"todo-list-api-roadmap-sh/internal/domain/todo"
)

type User struct {
	shared.BaseEntity
	Name     string      `gorm:"column:name;not null;size:100"`
	Email    string      `gorm:"column:email;not null;size:256"`
	Password string      `gorm:"column:password;not null;size:100"`
	Todos    []todo.Todo `gorm:"foreignKey:UserID"`
}
