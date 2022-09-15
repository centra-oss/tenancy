package models

import "go.h4n.io/centra/tenancy/models/meta"

type Tenant struct {
    Id string `json:"id"`
    Name string `json:"name"`
    DisplayName string `json:"displayName"`

    meta.CommonMetadata
}

