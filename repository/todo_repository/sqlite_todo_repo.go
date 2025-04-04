package repository

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteTodoRepo struct {
    db *sql.DB
}

func filterQueryBuilder(query string,filter QueryFilter) (string, []interface{}) {
    args := []interface{}{}

    if filter.Completed != nil {
        query += " AND completed = ?"
        args = append(args, *filter.Completed)
    }
    if filter.DueBefore != nil {
        query += " AND due < ?"
        args = append(args, filter.DueBefore.Format(time.RFC3339))
    }
    if filter.DueAfter != nil {
        query += " AND due > ?"
        args = append(args, filter.DueAfter.Format(time.RFC3339))
    }
    if filter.NameLike != nil {
        query += " AND name LIKE ?"
        args = append(args, "%"+*filter.NameLike+"%")
    }

    return query, args
}

func NewSQLiteTodoRepo(dbPath string) (ITodoRepository, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    createTableQuery := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        due DATETIME NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT 0
    );`
    _, err = db.Exec(createTableQuery)
    if err != nil {
        return nil, err
    }

    return &SQLiteTodoRepo{db: db}, nil
}

func (r *SQLiteTodoRepo) AddTodo(name string, due *time.Time) (int, error) {
    if due == nil {
        defaultDue := time.Now().AddDate(0, 0, 1)
        due = &defaultDue
    }

    result, err := r.db.Exec("INSERT INTO todos (name, due, completed) VALUES (?, ?, ?)", name, due.Format(time.RFC3339), false)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}

func (r *SQLiteTodoRepo) DeleteTodo(id int) error {
    result, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("todo with corresponding id not found")
    }

    return nil
}

func (r *SQLiteTodoRepo) CompleteTodo(id int) error {
    result, err := r.db.Exec("UPDATE todos SET completed = ? WHERE id = ?", true, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("todo with corresponding id not found")
    }

    return nil
}

func (r *SQLiteTodoRepo) GetTodos(filter QueryFilter) ([]Todo, error) {
	 query, args := filterQueryBuilder("SELECT id, name, due, completed FROM todos WHERE 1=1", filter)
    rows, err := r.db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []Todo
    for rows.Next() {
        var todo Todo
        var due string
        err := rows.Scan(&todo.Id, &todo.Name, &due, &todo.Completed)
        if err != nil {
            return nil, err
        }

        todo.Due, err = time.Parse(time.RFC3339, due)
        if err != nil {
            return nil, err
        }

        todos = append(todos, todo)
    }

    return todos, nil
}

func (r *SQLiteTodoRepo) ModifyTodo(id int, name *string, due *time.Time) error {
    todo, err := r.getTodoById(id)
    if err != nil {
        return err
    }

    if name != nil {
        todo.Name = *name
    }
    if due != nil {
        todo.Due = *due
    }

    result, err := r.db.Exec("UPDATE todos SET name = ?, due = ? WHERE id = ?", todo.Name, todo.Due.Format(time.RFC3339), id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("todo with corresponding id not found")
    }

    return nil
}

func (r *SQLiteTodoRepo) getTodoById(id int) (*Todo, error) {
    row := r.db.QueryRow("SELECT id, name, due, completed FROM todos WHERE id = ?", id)

    var todo Todo
    var due string
    err := row.Scan(&todo.Id, &todo.Name, &due, &todo.Completed)
    if err == sql.ErrNoRows {
        return nil, errors.New("todo with corresponding id not found")
    } else if err != nil {
        return nil, err
    }

    todo.Due, err = time.Parse(time.RFC3339, due)
    if err != nil {
        return nil, err
    }

    return &todo, nil
}