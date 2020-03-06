// +build prod

package config

const (
	MONGODB_HOST            = "mongodb://127.0.0.1:27017"
	MONGODB_DATABASE        = "apidb"
	MONGODB_CONNECTION_POOL = 50
	API_PORT                = 8080
)
