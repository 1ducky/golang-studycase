package todolog

type Status string

const (
	StatusCreate Status = "CREATE"
	StatusUpdate Status = "UPDATE"
	StatusDelete Status = "DELETE"
	StatusNull   Status = "NULL"
)

type CreateRequest struct {
	UserID     int
	TaskID     int
	FromStatus Status
	ToStatus   Status
}
