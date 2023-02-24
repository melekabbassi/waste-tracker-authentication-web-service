package main

import (
	"example/waste-tracker-authentication-web-service/database"
	"example/waste-tracker-authentication-web-service/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	loadENV()

	//Initialize Fiber app
	app := generateApp()

	db := database.OpenDB()
	defer database.CloseDB(db)

	//port := os.Getenv("PORT")

	app.Listen(":8081")

	//menu()

	// Define a route that handles GET requests to "/users"
	// http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	// Call the GetAllUsers function from the handlers package
	// 	users, err := handlers.GetAllUsers()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	// Encode the users slice to JSON format
	// 	jsonData, err := json.Marshal(users)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	// Set the Content-Type header to application/json
	// 	w.Header().Set("Content-Type", "application/json")

	// 	// Write the JSON data to the response
	// 	w.Write(jsonData)
	// })

	// routers.HandleAllUsers()

	// Start the server
	// log.Fatal(http.ListenAndServe(":8081", nil))
	/***********************************************************************/

}

func loadENV() error {
	goENV := os.Getenv("GO_ENV")
	if goENV == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func generateApp() *fiber.App {
	app := fiber.New()

	// create healthcheck route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create the dumptruck group and routes
	userGroup := app.Group("/users")
	userGroup.Get("/", handlers.GetUsers)

	return app
}

// func menu() {
// 	// create a menu to choose what to do use switch case
// 	var choice int
// 	for choice != 6 {
// 		fmt.Println("1. Get all users")
// 		fmt.Println("2. Get user by id")
// 		fmt.Println("3. Insert user")
// 		fmt.Println("4. Update user")
// 		fmt.Println("5. Delete user")
// 		fmt.Println("6. Exit")

// 		fmt.Println("Enter your choice: ")
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			users, err := handlers.GetAllUsers()
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println(users)
// 		case 2:
// 			var id int
// 			user, err := handlers.GetUserByID(id)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println(user)
// 		case 3:
// 			var email, username, password string
// 			fmt.Println("Enter email: ")
// 			fmt.Scanln(&email)
// 			fmt.Println("Enter username: ")
// 			fmt.Scanln(&username)
// 			fmt.Println("Enter password: ")
// 			fmt.Scanln(&password)

// 			// create a user
// 			user := handlers.UserDTO{Email: email, Username: username, Password: password}

// 			// insert the user into the database
// 			err := handlers.InsertUser(user)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		case 4:
// 			var id int
// 			var email, username, password string
// 			fmt.Println("Enter id: ")
// 			fmt.Scanln(&id)
// 			fmt.Println("Enter email: ")
// 			fmt.Scanln(&email)
// 			fmt.Println("Enter username: ")
// 			fmt.Scanln(&username)
// 			fmt.Println("Enter password: ")
// 			fmt.Scanln(&password)

// 			// create a user
// 			user := handlers.UserDTO{ID: id, Email: email, Username: username, Password: password}

// 			// update the user in the database
// 			err := handlers.UpdateUser(user)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		case 5:
// 			var id int
// 			fmt.Println("Enter id: ")
// 			fmt.Scanln(&id)

// 			// delete the user in the database
// 			err := handlers.DeleteUser(id)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		case 6:
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Invalid choice")
// 		}
// 	}
// }
