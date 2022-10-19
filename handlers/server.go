// app.go

package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/goccy/go-json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	Router *chi.Mux
	DB     *gorm.DB
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (server *Server) Initialize(username, password, hostname, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbname)

	var err error

	server.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	server.Router = chi.NewRouter()
	server.Router.Use(httprate.Limit(40, // requests
		10*time.Second, // per duration
		httprate.WithKeyFuncs(httprate.KeyByIP),
	), cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},

		// AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) initializeRoutes() {
	server.Router.Post("/signin", server.signIn)
	server.Router.Get("/users", server.getUsers)
	server.Router.Get("/userssearch", server.getUsersWithSearch)
	server.Router.Get("/listingssearch", server.getListingsWithSearch)
	server.Router.Get("/listing/{listing_id:[a-zA-z0-9]+}", server.getListing)
	server.Router.Get("/feed", server.getAllDefaultFeed)

	server.Router.Get("/searchbyimage", server.searchByImage)
	server.Router.Get("/searchbyimage2", server.searchByImage2)
	server.Router.Post("/user", server.createUser)
	server.Router.Post("/addtocart/{listing_id:[a-zA-z0-9]+}", server.addToCart)
	server.Router.Get("/cart/{user_id:[a-zA-z0-9]+}", server.getUserCart)
	server.Router.Delete("/removefromcart/{listing_id:[a-zA-z0-9]+}", server.removeFromCart)

	server.Router.Get("/user/{username:[a-zA-z0-9]+}", server.getUser)
	server.Router.Get("/user/{username:[a-zA-z0-9]+}/images", server.getAllUserImages)
	server.Router.Get("/user/{username:[a-zA-z0-9]+}/listings", server.getAllUserListings)

	server.Router.Get("/user/{username:[a-zA-z0-9]+}/images", server.getUserImagesByPage)

	server.Router.Put("/user", server.updateUser)
	server.Router.Delete("/user/{username:[a-zA-z0-9]+}", server.deleteUser)
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	return origin == "http://127.0.0.1:7070"

}

func respondWithError(w http.ResponseWriter, code int, message string) {

	payload := &ErrorResponse{Error: message}

	respondWithJSON(w, code, payload)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	response, err := json.Marshal(payload)
	if err != nil {
		log.Println("Marshalling JSON failed:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	w.Write(response)
	// log.Printf("\nResponse:%+v\n\n", payload)
}
