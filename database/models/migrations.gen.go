// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"gorm.io/gorm/schema"
)

const TableNameMigration = "migrations"

// Migration mapped from table <migrations>
type Migration struct {
	ID string `gorm:"column:id;type:character varying(255);primaryKey" json:"id"`
}

func (*Migration) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameMigration
	}
	return namer.TableName(TableNameMigration)
}