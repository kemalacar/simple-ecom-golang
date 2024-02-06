package models

import "gorm.io/gorm"

type StoreProduct struct {
	gorm.Model
	Price         float64 `gorm:""`
	Quantity      uint64  `gorm:""`
	StoreId       uint64
	ProductId     uint64
	Images        []SProductImage `gorm:"foreignKey:ProductStoreId;constraint:OnDelete:CASCADE"`
	ColorOptionId *uint64         `gorm:"foreignkey:ColorOptionId"`
	SizeOptionId  *uint64         `gorm:"foreignkey:SizeOptionId"`
	ColorOption   ProductOption   `gorm:"foreignKey:ColorOptionId;references:ID"`
	SizeOption    ProductOption   `gorm:"foreignKey:SizeOptionId;references:ID"`
}

type SProductImage struct {
	gorm.Model
	Big            string `gorm:"type:varchar(255)" json:"big"`
	ProductStoreId uint64
	Extension      string `gorm:"-"`
}

func StoreProductsToDto(sp []StoreProduct) (responseList []StoreProductDto) {
	for i := range sp {
		responseList = append(responseList, StoreProductToDto(sp[i]))
	}
	return responseList
}

func StoreProductToDto(p StoreProduct) StoreProductDto {
	return StoreProductDto{
		Id:       p.ID,
		Price:    p.Price,
		StoreId:  p.StoreId,
		Quantity: p.Quantity,
		Images:   SImagesToDto(p.Images),
		Color:    ProductOptionToDto(p.ColorOption),
		Size:     ProductOptionToDto(p.SizeOption),
	}
}

type StoreProductDto struct {
	Id       uint                  `json:"id,omitempty"`
	Price    float64               `json:"price,omitempty"`
	Quantity uint64                `json:"quantity,omitempty"`
	StoreId  uint64                `json:"storeId,omitempty"`
	Images   []ImageDto            `json:"images"`
	Color    ProductOptionResponse `json:"color"`
	Size     ProductOptionResponse `json:"size"`
}

func StoreProductToModel(p StoreProductDto) StoreProduct {
	product := StoreProduct{
		//ID:       p.Id,
		Price:    p.Price,
		StoreId:  p.StoreId,
		Quantity: p.Quantity,
		Images:   SImagesToModel(p.Images),
	}

	if p.Color.Id != 0 {
		product.ColorOptionId = &p.Color.Id
	}

	if p.Size.Id != 0 {
		product.SizeOptionId = &p.Size.Id
	}
	return product
}

func StoreProductsToModel(items []StoreProductDto) (responseList []StoreProduct) {
	for i := range items {
		responseList = append(responseList, StoreProductToModel(items[i]))
	}
	return responseList
}

func SImagesToDto(images []SProductImage) (responseList []ImageDto) {
	for i := range images {
		responseList = append(responseList, SImageToDto(images[i]))
	}
	return responseList
}

func SImageToDto(p SProductImage) ImageDto {
	return ImageDto{
		Big: p.Big,
	}
}

func SImagesToModel(images []ImageDto) (responseList []SProductImage) {
	for i := range images {
		responseList = append(responseList, SImageToModel(images[i]))
	}
	return responseList
}

func SImageToModel(p ImageDto) SProductImage {
	return SProductImage{
		Big:       p.Big,
		Extension: p.Extension,
	}
}
