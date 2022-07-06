package repositories

import (
	"CRUD/pkg/database"
	"CRUD/pkg/domain/users/entities"
)

func Create(user entities.User) {
	database.Connector.Create(&user)
}

func FetchAll() []entities.User {
	var users []entities.User
	database.Connector.Find(&users)
	return users
}

func FetchById(id int) entities.User {
	var user entities.User
	database.Connector.First(&user, id)
	return user
}

func Update(id int, user entities.User) {
	database.Connector.Model(&user).Where("id = ?", id).Updates(&user)
}

func Delete(id int) {
	database.Connector.Model(&entities.User{}).Where("id = ?", id).Delete(&entities.User{})
}
