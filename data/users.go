package data

import (
	"encoding/json"
	"fmt"
	"os"
)

// User represents a user's data
type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const usersFilePath = "./data/users.json"

func TestUsers() {
	// Example usage
	AddUser(User{Email: "user1@example.com", Username: "user1", Password: "password1"})
	AddUser(User{Email: "user2@example.com", Username: "user2", Password: "password2"})
	AddUser(User{Email: "user3@example.com", Username: "user3", Password: "password3"})

	user, err := GetUserByUsername("user2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User found:", user)
	}

	allUsers, err := GetAllUsers()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All users:", allUsers)
	}
}

// AddUser adds a new user to the JSON file
func AddUser(newUser User) error {
	users, err := GetAllUsers()
	if err != nil {
		return err
	}

	users = append(users, newUser)

	jsonBytes, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}

	file, err := os.Create(usersFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(username string) (User, error) {
	users, err := GetAllUsers()
	if err != nil {
		return User{}, err
	}

	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("user not found with username: %s", username)
}
// GetAllUsers retrieves all users from the JSON file
func GetAllUsers() ([]User, error) {
	file, err := os.Open(usersFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}


// RemoveUserByUsername removes a user by their username from the JSON file
func RemoveUserByUsername(usernameToRemove string) error {
	users, err := GetAllUsers()
	if err != nil {
		return err
	}

	var updatedUsers []User
	userFound := false
	for _, user := range users {
		if user.Username == usernameToRemove {
			userFound = true
		} else {
			updatedUsers = append(updatedUsers, user)
		}
	}

	if !userFound {
		return fmt.Errorf("user with username %s not found", usernameToRemove)
	}

	file, err := os.Create(usersFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonEncoder := json.NewEncoder(file)
	err = jsonEncoder.Encode(updatedUsers)
	if err != nil {
		return err
	}

	return nil
}