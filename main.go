package socialwebsite

import (
	"SocialWebsite/infrastructure/repository"
	"SocialWebsite/routers"
)

func main() {

	r := routers.User()
	defer repository.Mysqldb.Close()
	r.Run(":8080")
}
