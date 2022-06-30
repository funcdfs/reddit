package mysql

import "errors"

var (
	ErrorUserExists   = errors.New("user already exists: ")
	ErrorUserNotExist = errors.New("user does not exist: ")
	ErrorPassword     = errors.New("password error: ")
	ErrorInsertFailed = errors.New("insert failed: ")
	ErrorInValidID    = errors.New("invalid id: ")
	ErrorQueryError   = errors.New("query error: ")
)
