package services

import "github.com/retry19/challenge-hacktiv8/06-web-server-html/entities"

func FindUserByEmail(email string) *entities.User {
	for _, user := range entities.Users {
		if user.Email == email {
			return user
		}
	}
	return nil
}
