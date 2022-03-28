package ddseries

import (
	"encoding/json"
)

// Body is representation of the response body
type Body struct {
	Responses []Response `json:"responses"`
}

// Response is representation of the response
type Response struct {
	Status      string      `json:"status"`
	RespVersion int64       `json:"resp_version"`
	Series      []Series    `json:"series"`
	ToDate      int64       `json:"to_date"`
	Timing      string      `json:"timing"`
	Query       string      `json:"query"`
	Message     string      `json:"message"`
	ResType     string      `json:"res_type"`
	Interval    int64       `json:"interval"`
	Times       []float64   `json:"times"`
	FromDate    int64       `json:"from_date"`
	GroupBy     []any       `json:"group_by"`
	Values      [][]float64 `json:"values"`
}

// Series is representation of the series
type Series struct {
	Unit       []Unit `json:"unit"`
	QueryIndex int64  `json:"query_index"`
	Aggr       string `json:"aggr"`
	Scope      string `json:"scope"`
	Metric     string `json:"metric"`
	Expression string `json:"expression"`
	TagSet     []any  `json:"tag_set"`
}

// Unit is representation of the unit
type Unit struct {
	Family      string  `json:"family"`
	ScaleFactor float64 `json:"scale_factor"`
	Name        string  `json:"name"`
	ShortName   string  `json:"short_name"`
	Plural      string  `json:"plural"`
	ID          int64   `json:"id"`
}

// Unmarshall unmarshalls json
func (b *Body) Unmarshall(rawBody []byte) error {
	err := json.Unmarshal(rawBody, &b)
	if err != nil {
		return err
	}
	return nil
}
