package config

type Redis struct {
	Host string
	Port string

	// Server username
	// Optional. Default is ""
	Username string

	// Server password
	// Optional. Default is ""
	Password string

	// SelectDB to be selected after connecting to the server.
	// Optional. Default is 0
	SelectDB int
}
