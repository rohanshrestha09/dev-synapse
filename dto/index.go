package dto

import (
	"mime/multipart"
)

type RegisterDTO struct {
	Name            string `form:"name" binding:"required"`
	Email           string `form:"email" binding:"required,email" validate:"email"`
	Password        string `form:"password" binding:"required" validate:"gte=8"`
	ConfirmPassword string `form:"confirmPassword" binding:"required" validate:"eqfield=Password"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProfileUpdateDTO struct {
	Name  string                `form:"name"`
	Bio   string                `form:"bio"`
	Image *multipart.FileHeader `form:"image"`
}

type ProjectCreateDTO struct {
	Name              string                `form:"name" binding:"required"`
	Description       string                `form:"description" binding:"required"`
	Published         bool                  `form:"published"`
	Image             *multipart.FileHeader `form:"image" binding:"required"`
	EstimatedDuration uint                  `form:"estimatedDuration" binding:"required"`
}

type ProjectUpdateDTO struct {
	Name              string                `form:"name"`
	Description       string                `form:"description"`
	Image             *multipart.FileHeader `form:"image"`
	EstimatedDuration uint                  `form:"estimatedDuration"`
}
