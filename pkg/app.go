package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/levigross/grequests"
	"io"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	at     *int
}

func (a *App) Initilize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.getRoot).Methods("GET")
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/auth", a.getAuth).Methods("GET")
	a.Router.HandleFunc("/auth", a.getAuth).Methods("POST")
	a.Router.HandleFunc("/upload", a.receiveFile).Methods("POST")
	// a.Router.HandleFunc("/product", a.createProduct).Methods("POST")

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

func (a *App) getRoot(w http.ResponseWriter, r *http.Request) {

	log.Printf("get Root")
	products, err := getRoot(a.DB, 0, 5)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) receiveFile(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("File name %s\n", header.Filename)
	// Copy the file data to my buffer
	io.Copy(&Buf, file)
	// do something with the contents...
	// I normally have a struct defined and unmarshal into a struct, but this will
	// work as an example
	contents := Buf.String()
	fmt.Println(contents)
	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects
	Buf.Reset()
	// do something else
	// etc write header
	return
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {

	products, err := getProducts(a.DB, 0, 5)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) getAuth(w http.ResponseWriter, r *http.Request) {

	log.Printf("We are in getAuth\n")
	log.Printf("Method: %v\n", r.Method)
	log.Printf("Header: %v\n", r.Header)

	vals := r.URL.Query()
	code, ok := vals["code"]
	if ok {
		log.Printf("code: %v\n", code[0])

		ro := grequests.RequestOptions{}

		headers := map[string]string{}
		headers["Content-Type"] = "application/json"
		ro.Headers = headers
		ro.JSON = []byte(`{"num":6.13,"strs":["a","b"]}`)

		url := "http://httpbin.org/post"

		ResponseCode("code", ro, url)
		fmt.Fprint(w, "code")

	}
}

func ResponseCode(code string, ro grequests.RequestOptions,
	url string) *grequests.Response {

	response, err := grequests.Post(url, &ro)

	if err != nil {
		log.Printf("err: %v\n", err)
		return nil
	}
	log.Printf("out: %v\n", response)
	return response

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
