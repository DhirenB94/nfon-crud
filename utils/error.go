package utils

import "errors"

var ErrNoItemByID = errors.New("no item found for given id")
var ErrNoItemByName = errors.New("no items found for given name")
