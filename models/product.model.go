package models

import "gorm.io/gorm"

// Product ENTITY
type Product struct {
	gorm.Model
	Title         string         `gorm:"type:varchar(255);not null"`
	Detail        string         `gorm:"type:varchar(255);not null"`
	BrandId       uint64         `gorm:""`
	Brand         Brand          `gorm:"foreignKey:BrandId;references:ID"`
	StoreProducts []StoreProduct `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	Images        []ProductImage `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
}

func (Product) TableName() string {
	return "product"
}

type ProductImage struct {
	gorm.Model
	Big       string `gorm:"type:varchar(255)" json:"big"`
	ProductId uint64
	Extension string `gorm:"-"`
}

// ProductResponse RESPONSE
type ProductResponse struct {
	Id            uint              `json:"id,omitempty"`
	Title         string            `json:"title,omitempty"`
	Detail        string            `json:"detail,omitempty"`
	BrandId       uint64            `json:"brandId,omitempty"`
	Brand         BrandResponse     `json:"brand"`
	StoreProducts []StoreProductDto `json:"storeProducts"`
	Images        []ImageDto        `json:"images"`
}
type ImageDto struct {
	Big       string `json:"big"`
	Extension string `json:"extension"`
}

func ProductsToDto(Products []Product) (responseList []ProductResponse) {
	for i := range Products {
		responseList = append(responseList, ProductToDto(Products[i]))
	}
	return responseList
}

func ProductToDto(p Product) ProductResponse {
	return ProductResponse{
		Id:            p.ID,
		Title:         p.Title,
		Detail:        p.Detail,
		Brand:         BrandToDto(p.Brand),
		StoreProducts: StoreProductsToDto(p.StoreProducts),
		Images:        ImagesToDto(p.Images),
	}
}

func ImagesToDto(images []ProductImage) (responseList []ImageDto) {
	for i := range images {
		responseList = append(responseList, ImageToDto(images[i]))
	}
	return responseList
}

func ImageToDto(p ProductImage) ImageDto {
	return ImageDto{
		Big: p.Big,
	}
}

func ImagesToModel(images []ImageDto) (responseList []ProductImage) {
	for i := range images {
		responseList = append(responseList, ImageToModel(images[i]))
	}
	return responseList
}

func ImageToModel(p ImageDto) ProductImage {
	return ProductImage{
		Big:       p.Big,
		Extension: p.Extension,
	}
}

type CreateProductRequest struct {
	Title         string            `json:"title,omitempty"`
	Detail        string            `json:"detail,omitempty"`
	Brand         BrandDto          `json:"brand"`
	StoreProducts []StoreProductDto `json:"storeProducts"`
	Images        []ImageDto        `json:"images"`
}

func ProductToModel(p *CreateProductRequest) Product {
	return Product{
		Title:         p.Title,
		Detail:        p.Detail,
		BrandId:       p.Brand.Id,
		StoreProducts: StoreProductsToModel(p.StoreProducts),
		Images:        ImagesToModel(p.Images),
	}

}

type ImageHolder struct {
	Array []ProductImage
}
