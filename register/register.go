package register

import (
	"auth/db"
	"auth/hash"
	"encoding/json"
	"log"
	"net/http"
)




type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
}

func Register(r *http.Request) (*User, error){
	var u User

	err :=  json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal("json read error ", err)
	}
	passwordhash, err := hash.HashPassword(u.Password)
	if err != nil {
		log.Fatal("password hashing error", err)
	}
	u.Password = passwordhash
   
	conn := db.ConnectionDB()
	defer conn.Close()
_, err = conn.Exec(
    "INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
    u.Username, u.Email, u.Password,
)
  
if err != nil {
    return nil, err
}

	return &u, nil


}