package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230806_133430CreateDictionaryGroupItems) table() interface{} {
	type DictionaryGroupItems struct {
		gorm.Model
		DictionaryGroupID uint
		Key               string `gorm:"type:varchar(198); not null; comment:字典项键"`
		Value             string `gorm:"type:varchar(198); not null; comment:字典项值"`
		Order             uint   `gorm:"unsigned; not null;default 0;comment:排序"`
	}

	return &DictionaryGroupItems{}
}

func (m *Migration20230806_133430CreateDictionaryGroupItems) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230806_133430CreateDictionaryGroupItems) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230806_133430CreateDictionaryGroupItemsMigration)
}

type Migration20230806_133430CreateDictionaryGroupItems struct {
	id string
}

func New20230806_133430CreateDictionaryGroupItemsMigration() contracts.Migration {
	return &Migration20230806_133430CreateDictionaryGroupItems{id: "20230806_133430_create_dictionary_group_items"}
}

func (m *Migration20230806_133430CreateDictionaryGroupItems) ID() string {
	return m.id
}
