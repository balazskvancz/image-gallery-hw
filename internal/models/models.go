package models

type Image struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type Images = []*Image
