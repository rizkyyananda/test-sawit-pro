package request

type CreateEstateRequest struct {
	Width  int `json:"width" binding:"required,min=1,max=50000"`
	Length int `json:"length" binding:"required,min=1,max=50000"`
}
