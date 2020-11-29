package core

import "time"

type Model struct {
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt DeletedAt `gorm:"index"`
	// DeletedAt DeletedAt `sql:"index"`

	ID        int        `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
