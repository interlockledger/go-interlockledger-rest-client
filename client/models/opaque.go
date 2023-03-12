package models

import "time"

type OpaqueRecordModel struct {
	Network       string    `json:"network,omitempty"`
	ChainId       string    `json:"chainId,omitempty"`
	Serial        int64     `json:"serial,omitempty"`
	ApplicationId int64     `json:"applicationId,omitempty"`
	PayloadTagId  int64     `json:"payloadTagId,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
}
