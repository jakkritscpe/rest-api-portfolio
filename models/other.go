package models

type Tools struct {
	CommonFields
	Name          string        `json:"name,omitempty"`
	Urlimg        string        `json:"urlimg,omitempty"`
	CategoryID    int           `json:"category_id,omitempty"`
	CategoryTools CategoryTools `gorm:"foreignKey:CategoryID;references:ID"`
}

type CategoryTools struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Projects struct {
	CommonFields
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Urlimg      string `json:"urlimg,omitempty"`
	SkillID     int    `json:"category_id,omitempty"`
	Skills      Skills `gorm:"foreignKey:SkillID;references:ID"`
}

type Skills struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
