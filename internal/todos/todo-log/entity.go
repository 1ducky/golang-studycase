package todolog

type TodoLog struct {
	ID         int    `json:"id"`
	TaskId     int    `json:"taskId"`
	UserId     int    `json:"UserId"`
	FromStatus Status `json:"fromStatus"`
	ToStatus   Status `json:"toStatus"`
}
