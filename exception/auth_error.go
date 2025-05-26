package exception

type AuthError struct {
	Error string
}

func NewAuthError(err string) AuthError {
	return AuthError{Error: err}
}
