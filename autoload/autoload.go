package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/agclqq/godotenv/autoload"

	And bob's your mother's brother
*/

import "github.com/agclqq/godotenv"

func init() {
	godotenv.Load()
}
