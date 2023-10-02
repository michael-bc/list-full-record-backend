package brand_http

import "github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"

func LocationToDTO(location brand_domain.Location) LocationDTO {
	return LocationDTO{
		ID:        location.ID,
		City:      location.City,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}
}
