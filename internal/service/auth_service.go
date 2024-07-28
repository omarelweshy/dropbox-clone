package service

import "dropbox-clone/internal/repository"

type AuthService struct {
	AuthRepository repository.UserRepository
}
