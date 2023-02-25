package handlers

import (
	"example/waste-tracker-authentication-web-service/database"

	"github.com/gofiber/fiber/v2"
)

type UserDTO struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GET /users
func GetUsers(c *fiber.Ctx) error {
	db := database.OpenDB()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	users := make([]UserDTO, 0)

	for rows.Next() {
		user := UserDTO{}
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
		if err != nil {
			return err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(users)
}

// func GetAllUsers() ([]UserDTO, error) {
// 	db := database.OpenDB()
// 	defer database.CloseDB(db)

// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	database.CloseDB(db)

// 	var users []UserDTO // create a slice to store the users

// 	for rows.Next() {
// 		var user UserDTO
// 		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		//fmt.Println("id: " + string(user.ID) + "email: " + user.Email + ", username: " + user.Username + ", password: " + user.Password)
// 		users = append(users, user) // append each user to the slice
// 	}
// 	return users, nil // return the slice of users
// }

// func GetUserByID() (UserDTO, error) {
// 	// db := database.OpenDB()
// 	// defer database.CloseDB(db)

// 	// var user UserDTO

// 	// // ask the user for an id
// 	// fmt.Println("Enter the id of the user you want to get: ")
// 	// fmt.Scanln(&id)

// 	// err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
// 	// if err != nil {
// 	// 	panic(err.Error() + "\nUser not found")
// 	// }
// 	// return user, nil
// 	db := database.OpenDB()
// 	defer database.CloseDB(db)

// 	var user UserDTO

// 	err := db.QueryRow("SELECT * FROM users WHERE id = ?", user.ID).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// // insert a user into the database
// func InsertUser(user UserDTO) error {
// 	db := database.OpenDB()
// 	defer database.CloseDB(db)

// 	insert, err := db.Query("INSERT INTO users VALUES (?, ?, ?, ?)", user.ID, user.Email, user.Username, user.Password)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer insert.Close()
// 	return nil
// }

// // update a user in the database
// func UpdateUser(user UserDTO) error {
// 	db := database.OpenDB()
// 	defer database.CloseDB(db)

// 	update, err := db.Query("UPDATE users SET email = ?, username = ?, password = ? WHERE id = ?", user.Email, user.Username, user.Password, user.ID)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer update.Close()
// 	return nil
// }

// // delete a user from the database
// func DeleteUser(id int) error {
// 	db := database.OpenDB()
// 	defer database.CloseDB(db)

// 	delete, err := db.Query("DELETE FROM users WHERE id = ?", id)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer delete.Close()
// 	return nil
// }
