package brand_database

import (
	"github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"
	"github.com/samber/lo"
)

func criterionFromDTO(dto criterionDTO) brand_domain.Criterion {
	return brand_domain.Criterion{
		Category: categoryFromDTO(dto.Category),
		Location: locationFromDTO(dto.Location),
		Brands: lo.Map(dto.Brands, func(dto brandDTO, _ int) brand_domain.Brand {
			return BrandFromDTO(dto)
		}),
	}
}
