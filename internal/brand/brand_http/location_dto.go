package brand_http

import "time"

type LocationDTO struct {
	ID        int64     `json:"id"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
