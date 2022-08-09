package Repositories

import (
	"CRUD/pkg/database"
	"CRUD/pkg/domain/users/Entities"
)

func Create(user Entities.User) Entities.User {
	user.HashPassword()
	database.Connector.Create(&user)
	return user
}

func FetchAll() []Entities.User {
	var users []Entities.User
	database.Connector.Find(&users)
	return users
}

func FetchById(id int) Entities.User {
	var user Entities.User
	database.Connector.First(&user, id)

	user.HidePassword()
	return user
}

func Update(id int, user Entities.User) {
	database.Connector.Model(&user).Where("id = ?", id).Updates(&user)
}

func Delete(id int) {
	database.Connector.Model(&Entities.User{}).Where("id = ?", id).Delete(&Entities.User{})
}

func FetchByEmail(email string, user Entities.User) Entities.User {
	database.Connector.Model(&user).Where("email = ?", email).First(&user)
	return user
}
