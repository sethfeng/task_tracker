# Task Tracker

[EN](README.md) | [中文](README_cn.md)

A Go-based command-line task management tool that supports basic task operations, status management, and categorized task listing, with data persisted in a JSON file.

Practice Project: https://roadmap.sh/projects/task-tracker

## Features
- Basic task operations: Add, update, and delete tasks
- Status management: Mark tasks as pending, in progress, or done
- Categorized listing: View all tasks, completed tasks, in-progress tasks, or incomplete tasks
- Data persistence: Automatically stores tasks in a `tasks.json` file

## Installation & Execution
1. Ensure Go environment is installed (version ≥1.21)
2. Clone/download the project to local (directory should contain `main.go` and `go.mod`)
3. Run with command:
   ```bash
   go run main.go [command] [arguments]
   ```

## Command Reference
### 1. Add Task (add)
- Format: `task-tracker add <title> <description>`
- Example: `go run main.go add "Learn Go" "Complete basic syntax chapter"`
- Notes: Automatically generates incremental task ID with initial status "pending"

### 2. Update Task (update)
- Format: `task-tracker update <task_id> <new_title> <new_description>`
- Example: `go run main.go update 1 "Advanced Go Learning" "Complete concurrency chapter"`
- Notes: Updates task by ID

### 3. Delete Task (delete)
- Format: `task-tracker delete <task_id>`
- Example: `go run main.go delete 1`
- Notes: Deletes specified task (irreversible)

### 4. Mark Status (mark)
- Format: `task-tracker mark <task_id> <status>`
- Valid statuses: `pending` (to-do), `in progress` (ongoing), `done` (completed)
- Example: `go run main.go mark 1 "in progress"`
- Notes: Changes task status

### 5. List Tasks (list)
- Basic format: `task-tracker list` (view all tasks)
- Filter formats:
  - `task-tracker list done` (view completed tasks)
  - `task-tracker list in-progress` (view in-progress tasks)
  - `task-tracker list not-done` (view incomplete tasks, i.e., non-done statuses)
- Example: `go run main.go list done`

## Data Storage
Tasks are automatically stored in the `tasks.json` file in the same directory, with format:
```json
[
  {
    "id": 1,
    "title": "Learn JSON",
    "description": "Master JSON serialization",
    "status": "pending"
  },
  {
    "id": 2,
    "title": "Learn Go",
    "description": "Complete Go from beginner to advanced",
    "status": "in progress"
  }
]
```

## Notes
- First run will automatically create an empty `tasks.json` file
- All operations (add/update/delete/status change) automatically save to JSON file
- Invalid arguments (non-numeric ID, invalid status) will show error messages
```
