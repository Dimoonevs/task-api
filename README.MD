# Task-API

## One-shot run
```bash
    git clone https://github.com/yourname/task-api
    cd task-api
    make task-api-run-docker
```
API   — http://localhost:8080  
Mongo — mongodb://localhost:27017/taskdb  
Stop  — `make down`

## API Endpoints

| Method | URL         | Body / Query params                    | Description                         |
|--------|-------------|----------------------------------------|-------------------------------------|
| POST   | /tasks      | JSON: `title`, `description`, `status` | Create a new task                   |
| GET    | /tasks      | `status`, `page`, `pageSize`           | List tasks with filter & pagination |
| GET    | /tasks/{id} | -                                      | Get a task by ID                    |
| PUT    | /tasks/{id} | Full task JSON                         | Fully replace an existing task      |
| PATCH  | /tasks/{id} | Any subset of fields                   | Partially update an existing task   |
| DELETE | /tasks/{id} | –                                      | Delete a task                       |

Allowed `status` values:  
`todo`, `in_progress`, `in_qa`, `ready_for_release`, `done`

