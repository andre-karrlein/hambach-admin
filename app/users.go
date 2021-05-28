package main

const hash = "$2y$14$7aNuDEs7G6KxyYZLShEHlOpY4cjxV4kizm3noGFNBW11dvJdgtp3G"

// GetUsers function to get all available users with password
func GetUsers() map[string]string {
	return map[string]string{
		"akarrlein": hash,
		"pgeissler": hash,
	}
}

func GetNameOfUser() map[string]string {
	return map[string]string{
		"akarrlein": "Andre",
		"pgeissler": "Patrick",
	}
}
