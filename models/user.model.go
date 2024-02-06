package models

import "gorm.io/gorm"

// User ENTITY
type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Type      string `gorm:"not null"`
	StoreId   uint64 `gorm:"not null"`
}

// SignUpInput REQUEST
type SignUpInput struct {
	FirstName       string `json:"firstName" binding:"required"`
	LastName        string `json:"lastName" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	StoreId         uint64 `json:"storeId" binding:""`
	Type            string `json:"type" binding:""`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// UserResponse RESPONSE
type UserResponse struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	StoreId   uint64 `json:"storeId,omitempty"`
	Type      string `json:"type,omitempty"`
}

func UsersToDto(users []User) (responseList []UserResponse) {
	for i := range users {
		responseList = append(responseList, UserToDto(users[i]))
	}
	return responseList
}

func UserToDto(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		StoreId:   user.StoreId,
		Type:      user.Type,
	}
}
