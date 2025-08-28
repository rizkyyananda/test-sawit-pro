package entity

type Tree struct {
	ID       string `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	EstateID string `json:"estate_id" gorm:"type:uuid;not null;index"`
	X        int    `json:"x_coordinate" gorm:"not null;check:x > 0"`
	Y        int    `json:"y_coordinate" gorm:"not null;check:y > 0"`
	Height   int    `json:"height" gorm:"not null;check:height >= 1 and height <= 30"`
	Estate   Estate `gorm:"foreignKey:EstateID;references:ID;constraint:OnDelete:CASCADE"`
	BaseAuditEntry
}

func (a *Tree) TableName() string {
	return "t_tree"
}
