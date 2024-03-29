package models

type PaginationMetadata struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}
