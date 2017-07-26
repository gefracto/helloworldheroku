package main
import (
	"fmt"
	"net/http"
	"os"
	)


var counter = 0

func main() {
	fmt.Println("Hello dudes! ")
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	counter++
	
	w.Write([]byte(fmt.Sprintf("Hello %d times.", counter)))
}