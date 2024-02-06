package models

import (
	"gorm.io/gorm"
)

// ProductOption ENTITY
type ProductOption struct {
	gorm.Model
	Value      string            `gorm:"type:varchar(255);not null"`
	Attr1      string            `gorm:"type:varchar(255);not null"`
	OptionType ProductOptionType `gorm:"type:varchar(255);not null"`
}

func (ProductOption) TableName() string {
	return "product_option"
}

// CreateProductOptionRequest REQUEST
type CreateProductOptionRequest struct {
	Value      string            `json:"value,omitempty"`
	Attr1      string            `json:"attr1,omitempty"`
	OptionType ProductOptionType `json:"optionType,omitempty"`
}

// ProductOptionResponse RESPONSE
type ProductOptionResponse struct {
	Id         uint64            `json:"id,omitempty"`
	Value      string            `json:"value,omitempty"`
	Attr1      string            `json:"attr1,omitempty"`
	OptionType ProductOptionType `json:"optionType,omitempty"`
}

func ProductOptionsToDto(ProductOptions []ProductOption) (responseList []ProductOptionResponse) {
	for i := range ProductOptions {
		responseList = append(responseList, ProductOptionToDto(ProductOptions[i]))
	}
	return responseList
}

func ProductOptionToDto(po ProductOption) ProductOptionResponse {
	return ProductOptionResponse{
		Id:         uint64(po.ID),
		Value:      po.Value,
		Attr1:      po.Attr1,
		OptionType: po.OptionType,
	}
}

func ProductOptionToModel(cbr *CreateProductOptionRequest) (brand ProductOption) {
	return ProductOption{
		Value:      cbr.Value,
		Attr1:      cbr.Attr1,
		OptionType: cbr.OptionType,
	}
}

type ProductOptionType string

const (
	COLOR Status = "COLOR" // EnumIndex = 1
	SIZE         = "SIZE"  // EnumIndex = 2
)
