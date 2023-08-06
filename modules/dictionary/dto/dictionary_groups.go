package dto

import (
	"time"

	v1 "github.com/atom-apps/dictionary/proto/v1"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type DictionaryGroupForm struct {
	UserID      int64  `form:"user_id" json:"user_id,omitempty"`         // 用户ID
	TenantID    int64  `form:"tenant_id" json:"tenant_id,omitempty"`     // 租户ID
	Name        string `form:"name" json:"name,omitempty"`               // 字典组名称
	Slug        string `form:"slug" json:"slug,omitempty"`               //
	Description string `form:"description" json:"description,omitempty"` // 字典组描述
}

type DictionaryGroupListQueryFilter struct {
	TenantID    int64   `query:"tenant_id" json:"tenant_id,omitempty"`     // 租户ID
	UserID      int64   `query:"user_id" json:"user_id,omitempty"`         // 用户ID
	Name        *string `query:"name" json:"name,omitempty"`               // 字典组名称
	Slug        *string `query:"slug" json:"slug,omitempty"`               //
	Description *string `query:"description" json:"description,omitempty"` // 字典组描述
}

// FromProto
func (m *DictionaryGroupListQueryFilter) FromProto(req *v1.DictionaryService) error {
	return copier.Copy(&m, req)
}

type DictionaryGroupItem struct {
	ID          int64      `json:"id,omitempty"`          //
	CreatedAt   time.Time  `json:"created_at,omitempty"`  //
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`  //
	UserID      int64      `json:"user_id"`               // 用户ID
	TenantID    int64      `json:"tenant_id"`             // 租户ID
	Name        string     `json:"name,omitempty"`        // 字典组名称
	Slug        string     `json:"slug,omitempty"`        //
	Description string     `json:"description,omitempty"` // 字典组描述
	Items       []KeyValue `json:"items,omitempty"`       // items
}

func (m *DictionaryGroupItem) ToProto() *v1.DictionaryService {
	var resp *v1.DictionaryService
	lo.Must0(copier.Copy(&resp, m))
	return resp
}

type KeyValue struct {
	Key   string
	Value string
}