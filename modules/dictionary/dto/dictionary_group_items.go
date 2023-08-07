package dto

type DictionaryGroupItemForm struct {
	Value string `form:"value" json:"value,omitempty"` // 字典项值
	Order *int64 `form:"order" json:"order,omitempty"` // 排序
}

type DictionaryGroupItemListQueryFilter struct {
	DictionaryGroupID *int64  `query:"dictionary_group_id" json:"dictionary_group_id,omitempty"` //
	Value             *string `query:"value" json:"value,omitempty"`                             // 字典项值
	Order             *int64  `query:"order" json:"order,omitempty"`                             // 排序
}

type DictionaryGroupItemItem struct {
	ID    int64  `json:"id,omitempty"`    //
	Value string `json:"value,omitempty"` // 字典项值
	Order int64  `json:"order,omitempty"` // 排序
}
