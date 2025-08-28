package request

type AddTreeRequest struct {
	X      int `json:"x" binding:"required,min=1"`
	Y      int `json:"y" binding:"required,min=1"`
	Height int `json:"height" binding:"required,min=1,max=30"`
}
