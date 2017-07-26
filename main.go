package main
import (
	"fmt"
	"net/http"
	"os"
	// "html/template"
	// "log"
	)

func main() {
	fmt.Println("Hello dudes! ")
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	fmt.Fprintf(w, err.Error())
	// 	return
	// }
	// t.ExecuteTemplate(w, "index", nil)
	w.Write([]byte("hello blablabla"))
}