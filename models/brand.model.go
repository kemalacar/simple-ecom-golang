package models

import (
	"gorm.io/gorm"
)

// Brand  ENTITY
type Brand struct {
	gorm.Model
	LogoUrl string `gorm:"type:varchar(255);not null"`
	Name    string `gorm:"type:varchar(255);not null"`
}

func (Brand) TableName() string {
	return "brand"
}

// BrandResponse RESPONSE
type BrandResponse struct {
	Id      uint   `json:"id,omitempty"`
	LogoUrl string `json:"logoUrl,omitempty"`
	Name    string `json:"name,omitempty"`
}

func (brand *Brand) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

type BrandDto struct {
	Id      uint64 `json:"id" `
	LogoUrl string `json:"logoUrl" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

func BrandsToDto(brands []Brand) (responseList []BrandResponse) {
	for i := range brands {
		responseList = append(responseList, BrandToDto(brands[i]))
	}
	return responseList
}

func BrandToDto(brand Brand) BrandResponse {
	return BrandResponse{
		Id:      brand.ID,
		Name:    brand.Name,
		LogoUrl: brand.LogoUrl,
	}
}

func BrandToModel(cbr *BrandDto) (brand Brand) {
	return Brand{
		Name:    cbr.Name,
		LogoUrl: cbr.LogoUrl,
	}
}
