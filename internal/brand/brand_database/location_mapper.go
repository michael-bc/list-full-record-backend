package brand_database

import "github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"

func locationFromDTO(dto locationDTO) brand_domain.Location {
	return brand_domain.Location{
		ID:        dto.ID,
		City:      dto.City,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
