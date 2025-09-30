package main

import (
	"auth/db"
	"auth/register"
	"fmt"
	"net/http"
)


func main(){

	db := db.ConnectionDB()
	defer db.Close()
	fmt.Println("Database Connected")


	http.HandleFunc("/register", RegisterHandler)
	http.ListenAndServe(":8081", nil)

}


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "faqat post ruxsat", http.StatusMethodNotAllowed)
	}

	user, err := register.Register(r)
	if err != nil {
		http.Error(w, "xato: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Foydalanuvchi royhatdan otdi: %s", user.Username)
}
