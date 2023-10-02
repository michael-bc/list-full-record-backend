package brand_http

import "github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"

func CategoryToDTO(category brand_domain.Category) CategoryDTO {
	return CategoryDTO{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
