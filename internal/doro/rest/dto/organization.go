package dto

type CreateOrganizationRequestDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ImportOrganizationRequestDto struct {
	Source string `json:"source" binding:"required"`
}
