package main

const hash = "$2y$14$7aNuDEs7G6KxyYZLShEHlOpY4cjxV4kizm3noGFNBW11dvJdgtp3G"

type user struct {
	password string
	name     string
}

// GetUsers function to get all available users
func GetUsers() map[string]user {
	return map[string]user{
		"akarrlein": {password: hash, name: "Andre"},
		"pgeissler": {password: hash, name: "Patrick"},
	}
}
