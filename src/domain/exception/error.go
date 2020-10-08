package exception

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrAlreadyExist not found
var ErrAlreadyExist = errors.New("Email already exists")

