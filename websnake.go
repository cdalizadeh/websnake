package main

import (
    "net/http"
	"fmt"
	"os"
	"github.com/cdalizadeh/websnake/snake"
)

var PORT1 = ":8039"
var PORT2 = ":8040"

func main() {
	go createSocketServer(PORT2)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/js/glue.js", handleGlue)
	http.HandleFunc("/js/snake.js", handleSnake)
	fmt.Println("Now Serving.")
	http.ListenAndServe(PORT1, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "public/index.html")
	default:
		// Give an error message.
	}
}

func handleStranger(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "public/stranger.html")
	case "POST":
		// Return game page
	default:
		// Give an error message.
	}
}

func handleFriend(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "public/friend.html")
	case "POST":
		fmt.Println("recieve")
	default:
		// Give an error message.
	}
}

func testGame(){
	fmt.Println(string(3))
	sn := snake.CreateSnake()
	sn.Step()
	sn.Step()
	sn.Move(1, 0)
	sn.Step()
	sn.Step()
	sn.Move(1, 3)
	sn.Step()
	sn.Move(1, 2)
	sn.Step()
	sn.Step()
	sn.Step()
	sn.Step()
	sn.PrintState()
	os.Exit(3)
}

func handleGlue(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/js/glue.js")
}

func handleSnake(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/js/snake.js")
}