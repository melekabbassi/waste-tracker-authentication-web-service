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

// GET /users/:id
// if the user doesn't exist then it returns an error message and status code 500
func GetUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	user := UserDTO{}

	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return c.Status(500).SendString("User doesn't exist")
	}

	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(user)
}

// POST /users
// it accepts a JSON body then creates a new user if the previous id is 2 then the new id will be 3 and so on
// if the email or username already exists then it returns an error message and status code 500
func CreateUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	user := UserDTO{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// check if the email or username already exists
	var email, username string
	err = db.QueryRow("SELECT email, username FROM users WHERE email = ? OR username = ?", user.Email, user.Username).Scan(&email, &username)
	if err == nil {
		return c.Status(500).SendString("Email or Username already exists")
	}

	// get the last id
	var lastID int
	err = db.QueryRow("SELECT MAX(id) FROM users").Scan(&lastID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// insert the new user
	_, err = db.Exec("INSERT INTO users VALUES(?, ?, ?, ?)", lastID+1, user.Email, user.Username, user.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.CloseDB(db)

	return c.SendString("User created successfully")
}

// PUT /users/:id
func UpdateUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	user := UserDTO{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// check if the user exists
	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE id = ?", id).Scan(&userID)
	if err != nil {
		return c.Status(500).SendString("User doesn't exist")
	}

	// update the user
	_, err = db.Exec("UPDATE users SET email = ?, username = ?, password = ? WHERE id = ?", user.Email, user.Username, user.Password, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.CloseDB(db)

	return c.SendString("User updated successfully")
}

// DELETE /users/:id
// when the user is deleted the id of the next user will be the same as the deleted user that means if the id of the deleted user is 2 and the next is 3 the next user will be 2
func DeleteUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	// get the ID of the user to be deleted
	var deletedID int
	err := db.QueryRow("SELECT id FROM users WHERE id = ?", id).Scan(&deletedID)
	if err != nil {
		return c.Status(500).SendString("User doesn't exist")
	}

	// delete the user
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// update the IDs of the remaining users
	_, err = db.Exec("UPDATE users SET id = id - 1 WHERE id > ?", deletedID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.CloseDB(db)

	return c.SendString("User deleted successfully")
}
