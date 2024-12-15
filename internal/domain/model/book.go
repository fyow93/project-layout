package model

type Book struct {
	ID           string
	Title        string
	Author       string
	IsActive     bool
	CanBeDeleted bool
}
