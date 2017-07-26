package main
import (
	"fmt"
	"net/http"
	"os"
	"html/template"
	)

func main() {
	fmt.Println("Hello dudes! ")
	http.HandleFunc("/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "index", nil)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}