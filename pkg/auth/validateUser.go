package auth

func ValidateUser(username, password string) bool {
	users := GetUsers()
	if pass, ok := users[username]; ok {
		return pass == password
	}
	return false
}
