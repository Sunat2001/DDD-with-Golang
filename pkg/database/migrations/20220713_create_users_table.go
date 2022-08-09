package migrations

import (
	"CRUD/pkg/database"
	"CRUD/pkg/domain/users/Entities"
	"log"
)

type UserMigration struct{}

func (m *UserMigration) Up() {
	err := database.Connector.Migrator().AutoMigrate(&Entities.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (m *UserMigration) Down() {
	err := database.Connector.Migrator().DropTable(&Entities.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
