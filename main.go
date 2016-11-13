package main

import (
	"bublserver/handlers"
	"flag"
	"log"
	"net/http"
	"os"
)

//main: initialize database and start server
func main() {
	//db.InitDB()
	//defer db.SQLDB.Close()
	parseFlags()     //Command Line Arguments
	r := getRoutes() //create routes
	log.Println("Starting Jimmy Server..")
	http.ListenAndServe(":8080", r)
}

//getRoutes: create server routes
func getRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	return mux
}

//parseFlags: read command line arguments
func parseFlags() {
	//create flag pointers
	logPtr := flag.Bool("log", false, "Contols writing to log file.")
	flag.Parse() //parse flags
	//Handle flags
	if *logPtr == true {
		err := os.Remove("server.log") //remove local copy
		f, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Error opening logging file..")
		}
		log.SetOutput(f) //set logging to write to file
	}
}
