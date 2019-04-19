package entity_objects

import "errors"

var ErrNotFound = errors.New("Not found")

//ErrCannotBeDeleted bookmark cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")
