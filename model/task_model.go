package model

import (
	"fmt"
	"time"
)

type Task struct {
	ID          uint64    `db:"id"`
	Title       string    `db:"title"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt	time.Time `db:"updated_at"`
	IsDone      bool      `db:"is_done"`
	Description string    `db:"description"`
}

func Index() ([]Task, error) {
	var tasks []Task
	var err = db.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func Show(id int) (Task, error) {
	var task Task
	var err = db.Get(&task, "SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func Create(title string, description string) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", title, description)
	if err != nil {
		return -1, err
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return -1, nil
	}
	return taskID, nil
}

func Update(id int, title string, isDone bool, description string) (string, error) {
	result, err := db.Exec("UPDATE tasks SET title = ?, is_done = ?, description = ? WHERE id = ?", title, isDone, description, id)
	if err != nil {
		return "", err
	}

	path := "/list"
	n, err := result.RowsAffected();
	fmt.Println(n)
	if n == 1 && err == nil {
		path = fmt.Sprintf("/task/%d", id) // 正常にIDを取得できた場合は /task/<id> へ戻る
	}
	return path, nil
}

func Delete(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}