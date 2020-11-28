package main

import "github.com/dacharat/rlp-api/apps"

func main() {
	router := apps.GenerateRouter()
	router.Run()
}
