package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func getPosts(w http.ResponseWriter, request *http.Request) {

	fmt.Fprint(w, "Posts are")
}

func testMux() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	http.ListenAndServe(":8000", router)
}

var mySignedKey = []byte("mySuperSecretPhrase")

func generateJWTToken() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = "Mudassar"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySignedKey)

	return tokenString, err
}

func main() {

	fmt.Println("starting day 4")

	// connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	//abc.TestingAgain()

	token, err := generateJWTToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	//setupDatabase()
}
