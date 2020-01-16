package main

import (
	ApiwithGoChi "ApiWithGoChi.com"
	"fmt"
)

var mh *ApiwithGoChi.MongoHandler

func main(){
	mongoDbConnection := "mongodb://localhost:27017"
	mh  := ApiwithGoChi.NewHandler(mongoDbConnection)
	if mh != nil {
		fmt.Println("Success")
	}

}
