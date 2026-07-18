package http

import (
	"net/http"
	"restApi/internal/query"
	"strconv"
)

var DefaultValues query.GenericQuery = query.GenericQuery{
	Cursor:   0,
	Limit:    10,
	Search:   "",
	SortedBy: query.SortedByDesc,
	Order:    query.OrderByCreatedAt,
}

// ParseGenericRequest membaca query params dari http.Request,
// mem-parsing ke GenericRequest, dan mengisi default value jika kosong/invalid.
//
// Contoh query: /orders?cursor=5&limit=20&search=john&sorted_by=ASC&order=updated_at
func ParseGenericRequest(r *http.Request) query.GenericQuery {
	q := r.URL.Query()

	req := query.GenericQuery{
		Cursor:   parseIntOrDefault(q.Get("cursor"), DefaultValues.Cursor),
		Limit:    parseIntOrDefault(q.Get("limit"), DefaultValues.Limit),
		Search:   parseStringOrDefault(q.Get("search"), DefaultValues.Search),
		SortedBy: parseSortedOrDefault(q.Get("sorted_by")),
		Order:    parseOrderOrDefault(q.Get("order")),
	}

	return req
}

func parseIntOrDefault(val string, def int) int {
	if val == "" {
		return def
	}
	n, err := strconv.Atoi(val)
	if err != nil || n < 0 {
		return def
	}
	return n
}

func parseStringOrDefault(val string, def string) string {
	if val == "" {
		return def
	}
	return val
}

func parseSortedOrDefault(val string) query.GenericSorted {
	switch query.GenericSorted(val) {
	case query.SortedByAsc:
		return query.SortedByAsc
	case query.SortedByDesc:
		return query.SortedByDesc
	default:
		return DefaultValues.SortedBy
	}
}

func parseOrderOrDefault(val string) query.GenericOrder {
	switch query.GenericOrder(val) {
	case query.OrderByCreatedAt:
		return query.OrderByCreatedAt
	case query.OrderByUpdatedAt:
		return query.OrderByUpdatedAt
	default:
		return DefaultValues.Order
	}
}
