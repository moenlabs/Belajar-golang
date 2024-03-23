package database

var connection string

func init() {
	connection = "posgresql"
}

func GetDatabase() string {
	return connection
}
