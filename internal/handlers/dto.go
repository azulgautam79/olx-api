package handlers

import "time"

type CreateListingRequest struct {
	Title       string `json:"title" example:"Iphone 18 Pro"`
	Description string `json:"description" example:"Latest Iphone from Apple with fold also available"`
	Price       int64  `json:"price" example:"300000"`
	City        string `json:"city" example:"KTM"`
}

type CreateListingResponse struct {
	ID          string    `json:"id" example:"2323423423423423"`
	Title       string    `json:"title" example:"Iphone 18 Pro"`
	Description string    `json:"description" example:"Latest Iphone from Apple with fold also available"`
	Price       int64     `json:"price" example:"300000"`
	City        string    `json:"city" example:"KTM"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UpdateListingRequest struct {
	ID          string `json:"id" example:"2323423423423423"`
	Title       string `json:"title" example:"Iphone 18 Pro"`
	Description string `json:"description" example:"Latest Iphone from Apple with fold also available"`
	Price       int64  `json:"price" example:"300000"`
	City        string `json:"city" example:"KTM"`
}
