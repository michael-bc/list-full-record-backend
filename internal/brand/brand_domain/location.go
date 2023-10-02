package brand_domain

import "time"

type Location struct {
	ID        int64
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
