package rpkg

import (
	"database/sql"
	jsonparse "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Router *mux.Router
	Server *rpc.Server
	DB     *sql.DB
	at     *int
}

// Args holds arguments passed to JSON RPC service
type Args struct {
	Id string
}

// Book struct holds Book JSON structure
type Book struct {
	Id     string `"json:string,omitempty"`
	Name   string `"json:name,omitempty"`
	Author string `"json:author,omitempty"`
}
type JSONServer struct{}

// GiveBookDetail
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	// Read JSON file and load data
	raw := []byte(`
[
	{
		"id": "1234",
		"name": "In the sunburned country",
		"author": "Bill Bryson"
	},
	{
		"id":"2345",
		"name": "The picture of Dorian Gray",
		"author": "Oscar Wilde"
	}
]
`)

	// Unmarshal JSON raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// Iterate over each book to find the given book
	for _, book := range books {
		if book.Id == args.Id {
			// If book found, fill reply with it
			*reply = book
			break
		}
	}
	return nil
}

func (a *App) Initilize() {

	// Create a new RPC server
	a.Server = rpc.NewServer() // Register the type of data requested as JSON
	a.Server.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	a.Server.RegisterService(new(JSONServer), "")
	a.Router = mux.NewRouter()
	a.Router.Handle("/rpc", a.Server)

}

func (a *App) Run(addr string, writeTimeout int, readTimeout int) {

	srv := &http.Server{
		Handler: a.Router,
		Addr:    fmt.Sprintf(":%s", addr),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

