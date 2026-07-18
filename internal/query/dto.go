package query

type GenericOrder string

const (
	OrderByCreatedAt GenericOrder = "created_at"
	OrderByUpdatedAt GenericOrder = "updated_at"
)

type GenericSorted string

const (
	SortedByDesc GenericSorted = "DESC"
	SortedByAsc  GenericSorted = "ASC"
)

type GenericQuery struct {
	Cursor   int
	Limit    int
	Search   string
	SortedBy GenericSorted
	Order    GenericOrder
}
