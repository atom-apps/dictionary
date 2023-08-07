package seeders

import (
	"log"

	"github.com/atom-apps/dictionary/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

type DictionaryGroupsSeeder struct{}

func NewDictionaryGroupsSeeder() contracts.Seeder {
	return &DictionaryGroupsSeeder{}
}

func (s *DictionaryGroupsSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	times := 10
	for i := 0; i < times; i++ {
		data := s.Generate(faker, i)
		if i == 0 {
			stmt := &gorm.Statement{DB: db}
			_ = stmt.Parse(&data)
			log.Printf("seeding %s for %d times", stmt.Schema.Table, times)
		}
		db.Create(&data)
	}
}

func (s *DictionaryGroupsSeeder) Generate(faker *gofakeit.Faker, idx int) models.DictionaryGroup {
	return models.DictionaryGroup{
		UserID:      int64(faker.RandomUint([]uint{0, 1, 2})),
		TenantID:    int64(faker.RandomUint([]uint{0, 1, 2})),
		Name:        faker.Name(),
		Slug:        faker.UUID(),
		Description: faker.Sentence(10),
	}
}
