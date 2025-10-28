package errmsg

import "errors"

//Custom errors
var ErrUserNotFound error = errors.New("user not found")
var ErrNotEnoughMoney error = errors.New("not enough money")
var ErrIncorrectAmount error = errors.New("incorrect amount")
