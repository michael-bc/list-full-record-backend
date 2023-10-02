package brand_database

import (
	"github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"
)

func BrandFromDTO(dto brandDTO) brand_domain.Brand {
	return brand_domain.Brand{
		ID:            dto.ID,
		Name:          dto.Name,
		Image:         dto.Image,
		Location:      locationFromDTO(dto.Location),
		ProductsCount: dto.ProductsCount,
		AverageRating: dto.AverageRating,
	}
}
