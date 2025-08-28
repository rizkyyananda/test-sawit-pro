package response

type Coordinate struct {
	X             int `json:"x"`
	Y             int `json:"y"`
	TotalDistance int `json:"total"`
}
