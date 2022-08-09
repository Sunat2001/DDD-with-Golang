package main

import "CRUD/pkg/database/migrations"

func main() {
	migration := migrations.UserMigration{}
	migration.Down()
}
