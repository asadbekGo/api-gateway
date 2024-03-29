package models

type Todo struct {
	ID        string `json:"id"`
	Assignee  string `json:"assignee"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateTodo struct {
	Assignee string `json:"assignee"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}

type ListTodos struct {
	Todos []Todo `json:"todos"`
}
