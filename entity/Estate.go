package entity

type Estate struct {
	ID     string `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Width  int    `json:"width" gorm:"not null;check:width > 0"`
	Length int    `json:"length" gorm:"not null;check:length > 0"`

	Trees []Tree `json:"trees" gorm:"foreignKey:EstateID;constraint:OnDelete:CASCADE"`
}

func (a *Estate) TableName() string {
	return "t_estate"
}
