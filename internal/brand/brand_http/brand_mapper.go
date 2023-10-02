package brand_http

import "github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"

func BrandToDTO(brand brand_domain.Brand) BrandDTO {
	return BrandDTO{
		ID:            brand.ID,
		Name:          brand.Name,
		Image:         brand.Image,
		Location:      LocationToDTO(brand.Location),
		ProductsCount: brand.ProductsCount,
		AverageRating: brand.AverageRating,
		CreatedAt:     brand.CreatedAt,
		UpdatedAt:     brand.UpdatedAt,
	}
}

func BrandsToDTOs(brands []brand_domain.Brand) []BrandDTO {
	var dtos []BrandDTO
	for _, brand := range brands {
		dtos = append(dtos, BrandToDTO(brand))
	}
	return dtos
}
