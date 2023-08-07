package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230806_133418CreateDictionaryGroup) table() interface{} {
	type DictionaryGroup struct {
		gorm.Model
		UserID      uint   `gorm:"unsigned; not null;default 0;comment:用户ID"`
		TenantID    uint   `gorm:"unsigned; not null;default 0;comment:租户ID"`
		Name        string `gorm:"type:varchar(198);not null;comment:字典组名称"`
		Slug        string `gorm:"type:varchar(120);not null;comment:字典组别名"`
		Description string `gorm:"type:varchar(198);not null;comment:字典组描述"`
	}

	return &DictionaryGroup{}
}

func (m *Migration20230806_133418CreateDictionaryGroup) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230806_133418CreateDictionaryGroup) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
	// return tx.Migrator().DropColumn(m.table(), "input_column_name")
}

// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
func init() {
	Migrations = append(Migrations, New20230806_133418CreateDictionaryGroupMigration)
}

type Migration20230806_133418CreateDictionaryGroup struct {
	id string
}

func New20230806_133418CreateDictionaryGroupMigration() contracts.Migration {
	return &Migration20230806_133418CreateDictionaryGroup{id: "20230806_133418_create_dictionary_group"}
}

func (m *Migration20230806_133418CreateDictionaryGroup) ID() string {
	return m.id
}
