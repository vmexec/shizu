package auth

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func GetUsers() map[string]string {
	file, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var users []User
	json.Unmarshal(byteValue, &users)

	userMap := make(map[string]string)
	for _, user := range users {
		userMap[user.Username] = user.Password
	}

	return userMap
}
