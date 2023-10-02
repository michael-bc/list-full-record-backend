package brand_database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type locationDTO struct {
	ID        int64     `json:"id"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (l locationDTO) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *locationDTO) Scan(data any) error {
	if data == nil {
		return nil
	}

	switch val := data.(type) {
	case []byte:
		return json.Unmarshal(val, &l)
	case string:
		return json.Unmarshal([]byte(val), &l)
	default:
		return fmt.Errorf("unexpected data type - %T", data)
	}
}
