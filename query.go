package ddseries

import "fmt"

// Query holds data for query
type Query struct {
	Q        string
	From     int64
	To       int64
	Interval int64
}

// String generates query string
func (q *Query) String(index int) string {
	return fmt.Sprintf(
		"requests[%d][q]=%s&requests[%d][from]=%d&requests[%d][to]=%d&requests[%d][interval]=%d",
		index,
		q.Q,
		index,
		q.From,
		index,
		q.To,
		index,
		q.Interval,
	)
}
