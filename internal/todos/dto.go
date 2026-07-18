package todos

type result[T any] struct {
	data T
	err  error
}
type Error struct {
	Code    string `json:"code"` //frontend code translate
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// input contract
type CreateRequest struct {
	Task string `json:"task"`
}
type UpdateRequest struct {
	ID     int
	Task   string
	IsDone bool
}

type DeleteRequest struct {
	ID int
}

type TodosMeta struct {
	NextCursor int  `json:"nextCursor,omitempty"`
	Limit      int  `json:"limit,omitempty"`
	HasNext    bool `json:"hasNext"`
}
