package meta

import "time"

type CommonMetadata struct {
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

