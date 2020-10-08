package main

import "SocialWebsite/routers"

func main() {

	r := routers.Account()
	r.Run(":8080")
}