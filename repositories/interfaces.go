package repositories

import "github.com/caiomarinello/ninjaGateway/components"

// Registrar is the interface that wraps the basic Register method.
// Register is supposed to take a component and save it into a database.
type Registrar[T interface{}] interface {
	Register(component T) error
}

// UserFetcher is the interface that wraps the method FetchUserByEmail.
type UserFetcher interface {
	FetchUserByEmail(email string) (*components.User, error)
}
