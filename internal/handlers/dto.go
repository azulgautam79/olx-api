package handlers

import (
	"fmt"
	"strings"
	"time"
)

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

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

func (req CreateListingRequest) Validate() error {
    if strings.TrimSpace(req.Title) == "" {
        return &ValidationError{Field: "title", Msg: "must not be empty"}
    }

    if strings.TrimSpace(req.Description) == "" {
        return &ValidationError{Field: "description", Msg: "must not be empty"}
    }

    if req.Price <= 0 {
        return &ValidationError{Field: "price", Msg: "must be greater than zero"}
    }

    if strings.TrimSpace(req.City) == "" {
        return &ValidationError{Field: "city", Msg: "must not be empty"}
    }

    return nil
}
