package models

import (
	"gorm.io/gorm"
)

// Store ENTITY
type Store struct {
	gorm.Model
	Title  string `gorm:"type:varchar(255);not null"`
	Email  string `gorm:"type:varchar(255);not null"`
	Status Status `gorm:"type:varchar(255);not null"`
}

func (Store) TableName() string {
	return "store"
}

// CreateStoreRequest REQUEST
type CreateStoreRequest struct {
	Title  string `json:"title,omitempty"`
	Email  string `json:"email,omitempty"`
	Status Status `json:"status,omitempty"`
}

// StoreResponse RESPONSE
type StoreResponse struct {
	Id     uint   `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Email  string `json:"email,omitempty"`
	Status Status `json:"status,omitempty"`
}

func StoresToDto(stores []Store) (responseList []StoreResponse) {
	for i := range stores {
		responseList = append(responseList, StoreToDto(stores[i]))
	}
	return responseList
}

func StoreToDto(brand Store) StoreResponse {
	return StoreResponse{
		Id:     brand.ID,
		Title:  brand.Title,
		Email:  brand.Email,
		Status: brand.Status,
	}
}

func StoreToModel(cbr *CreateStoreRequest) (brand Store) {
	return Store{
		Title:  cbr.Title,
		Email:  cbr.Email,
		Status: cbr.Status,
	}
}

type Status string

const (
	ACTIVE  Status = "ACTIVE"  // EnumIndex = 1
	PASSIVE        = "PASSIVE" // EnumIndex = 2
)
