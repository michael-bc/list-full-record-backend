package brand_database

import "github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"

func categoryFromDTO(dto categoryDTO) brand_domain.Category {
	return brand_domain.Category{
		ID:        dto.ID,
		Name:      dto.Name,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
