// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package postgres

import (
	"context"

	"github.com/jackc/pgtype"
)

const getLimitedTasksByStatus = `-- name: GetLimitedTasksByStatus :many
SELECT id, name, type, status, priority, payload, created_at, updated_at FROM tasks WHERE status = $1 LIMIT $2
`

type GetLimitedTasksByStatusParams struct {
	Status TaskStatus
	Limit  int32
}

func (q *Queries) GetLimitedTasksByStatus(ctx context.Context, arg GetLimitedTasksByStatusParams) ([]Task, error) {
	rows, err := q.db.Query(ctx, getLimitedTasksByStatus, arg.Status, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Status,
			&i.Priority,
			&i.Payload,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMissedTasks = `-- name: GetMissedTasks :many
SELECT id, name, type, status, priority, payload, created_at, updated_at
FROM tasks
WHERE status = $1 AND updated_at <= now() - ($2 * interval '1 second') LIMIT $3
`

type GetMissedTasksParams struct {
	Status  TaskStatus
	Column2 interface{}
	Limit   int32
}

func (q *Queries) GetMissedTasks(ctx context.Context, arg GetMissedTasksParams) ([]Task, error) {
	rows, err := q.db.Query(ctx, getMissedTasks, arg.Status, arg.Column2, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Status,
			&i.Priority,
			&i.Payload,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, name, type, status, priority, payload, created_at, updated_at FROM tasks WHERE id = $1
`

func (q *Queries) GetTaskByID(ctx context.Context, id int32) (Task, error) {
	row := q.db.QueryRow(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.Status,
		&i.Priority,
		&i.Payload,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTaskStatusChangeHistory = `-- name: GetTaskStatusChangeHistory :many
SELECT id, task_id, old_status, new_status, created_at FROM tasks_status_change_history WHERE task_id = $1
`

func (q *Queries) GetTaskStatusChangeHistory(ctx context.Context, taskID int32) ([]TasksStatusChangeHistory, error) {
	rows, err := q.db.Query(ctx, getTaskStatusChangeHistory, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TasksStatusChangeHistory
	for rows.Next() {
		var i TasksStatusChangeHistory
		if err := rows.Scan(
			&i.ID,
			&i.TaskID,
			&i.OldStatus,
			&i.NewStatus,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTasksByStatus = `-- name: GetTasksByStatus :many
SELECT id, name, type, status, priority, payload, created_at, updated_at FROM tasks WHERE status = $1
`

func (q *Queries) GetTasksByStatus(ctx context.Context, status TaskStatus) ([]Task, error) {
	rows, err := q.db.Query(ctx, getTasksByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Status,
			&i.Priority,
			&i.Payload,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertTask = `-- name: InsertTask :one
INSERT INTO tasks (
    name, type, status, priority, payload
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING id
`

type InsertTaskParams struct {
	Name     string
	Type     TaskType
	Status   TaskStatus
	Priority TaskPriority
	Payload  pgtype.JSON
}

func (q *Queries) InsertTask(ctx context.Context, arg InsertTaskParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertTask,
		arg.Name,
		arg.Type,
		arg.Status,
		arg.Priority,
		arg.Payload,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertTaskStatusChangeHistory = `-- name: InsertTaskStatusChangeHistory :exec
INSERT INTO tasks_status_change_history (
    task_id, old_status, new_status
) VALUES (
             $1, $2, $3
         )
`

type InsertTaskStatusChangeHistoryParams struct {
	TaskID    int32
	OldStatus TaskStatus
	NewStatus TaskStatus
}

func (q *Queries) InsertTaskStatusChangeHistory(ctx context.Context, arg InsertTaskStatusChangeHistoryParams) error {
	_, err := q.db.Exec(ctx, insertTaskStatusChangeHistory, arg.TaskID, arg.OldStatus, arg.NewStatus)
	return err
}

const updateTaskStatus = `-- name: UpdateTaskStatus :exec
UPDATE tasks SET status = $1 WHERE id = $2
`

type UpdateTaskStatusParams struct {
	Status TaskStatus
	ID     int32
}

func (q *Queries) UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) error {
	_, err := q.db.Exec(ctx, updateTaskStatus, arg.Status, arg.ID)
	return err
}
