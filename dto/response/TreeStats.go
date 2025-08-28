package response

type TreeStats struct {
	Count  int `json:"count"`
	Max    int `json:"max"`
	Min    int `json:"min"`
	Median int `json:"median"`
}
