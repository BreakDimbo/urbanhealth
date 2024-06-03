package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// if err := db.Init(); err != nil {
	// 	panic(fmt.Sprintf("mysql init failed with %+v", err))
	// }

	r := SetupRouter()
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
