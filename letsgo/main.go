package main

import (
	"crypto/md5"
	"crypto/rc4"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	_ "unsafe"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	http.HandleFunc("/exec", execHandler)
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/readfile", readFile)
	http.ListenAndServe(":8080", nil)
}


func execHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Output: %s", out)
}


func queryHandler(w http.ResponseWriter, r *http.Request) {
    var password = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"
	user := r.URL.Query().Get("user")
	query := fmt.Sprintf("SELECT * FROM users WHERE user='%s'", user)
	db, err := sql.Open("mysql", "user:"+password+"@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process rows
}

func vCrypto() {
	hash := md5.New()
	data := []byte("password")
	hash.Write(data)
	fmt.Printf("MD5 hash: %x\n", hash.Sum(nil))

	cipher, err := rc4.NewCipher(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RC4 cipher:", cipher)
}

func readFile(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Query().Get("file")
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "File content: %s", data)
}
