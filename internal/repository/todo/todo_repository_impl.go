package repository

import (
	"database/sql"
	"encoding/json"
	"todo-restful-api/internal/model/domain"
)

type TodoRepositoryImpl struct {
	DB *sql.DB
}

// FindById implements TodoRepository.
func (repository *TodoRepositoryImpl) FindById(todoId int) (domain.Todo, error) {
	query := `
		SELECT todo.id, todo.title, todo.status, todo.created_at, todo.updated_at
		FROM todos AS todo
		JOIN users AS user ON (todo.user_id = user.id)
		WHERE todo.id = (?)
	`
	var todo domain.Todo

	row := repository.DB.QueryRow(query, todoId)

	if err := row.Scan(&todo.Id, &todo.Title, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return todo, err
	}

	return todo, nil
}

// Create implements TodoRepository.
func (repository TodoRepositoryImpl) Create(todo domain.Todo) domain.Todo {
	query := `INSERT INTO todos(title, user_id) VALUES (?, ?)`

	result, err := repository.DB.Exec(query, todo.Title, todo.UserId)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	todo.Id = int(id)

	return todo
}

// Delete implements TodoRepository.
func (repository TodoRepositoryImpl) Delete(todoId int, userId int) {
	query := `delete from todos where id = (?) and user_id = (?)`
	if _, err := repository.DB.Exec(query, todoId, userId); err != nil {
		panic(err)
	}
}

// FindAll implements TodoRepository.
func (repository TodoRepositoryImpl) FindAll(todoId int) []domain.Todo {
	query := `
		SELECT todo.id, todo.title, todo.status, json_object('id', user.id, 'email', user.email ,'username', user.username) AS user,
			 todo.created_at ,todo.updated_at
		FROM todos AS todo
		JOIN users AS user ON (todo.user_id = user.id)
		WHERE user.id = (?)
	`
	rows, err := repository.DB.Query(query, todoId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var todos []domain.Todo
	var userJSON string

	for rows.Next() {
		var todo domain.Todo

		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Status, &userJSON, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			panic(err)
		}
		if err := json.Unmarshal([]byte(userJSON), &todo.User); err != nil {
			panic(err)
		}

		todos = append(todos, todo)
	}
	return todos
}

// Update implements TodoRepository.
func (repository TodoRepositoryImpl) Update(todo domain.Todo) domain.Todo {
	query := `
		UPDATE todos
		SET title = (?), status = (?), updated_at = current_timestamp
		WHERE id = (?) AND user_id = (?)`
	if _, err := repository.DB.Exec(query, todo.Title, todo.Status, todo.Id, todo.UserId); err != nil {
		panic(err)
	}
	return todo
}

func NewTodoRepository(DB *sql.DB) TodoRepository {
	return &TodoRepositoryImpl{
		DB: DB,
	}
}
