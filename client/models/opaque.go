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

type PageOfOpaqueRecordsModel struct {
	Items                   []OpaqueRecordModel `json:"items,omitempty"`
	Page                    int                 `json:"page,omitempty"`
	PageSize                int                 `json:"pageSize,omitempty"`
	TotalNumberOfPages      int                 `json:"totalNumberOfPages,omitempty"`
	LastToFirst             bool                `json:"lastToFirst,omitempty"`
	LastChangedRecordSerial int64               `json:"lastChangedRecordSerial,omitempty"`
}
