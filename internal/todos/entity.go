package todos

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	IsDone    bool      `json:"isDone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
