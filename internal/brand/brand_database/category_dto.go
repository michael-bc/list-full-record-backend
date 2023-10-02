package brand_database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type categoryDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c categoryDTO) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *categoryDTO) Scan(data any) error {
	if data == nil {
		return nil
	}

	switch val := data.(type) {
	case []byte:
		return json.Unmarshal(val, &c)
	case string:
		return json.Unmarshal([]byte(val), &c)
	default:
		return fmt.Errorf("unexpected data type - %T", data)
	}
}
