package main

import (
	"fmt"
	"net/http"
	"bufio"
	"os"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

type Message struct {
	Data string `json:"data"`
}

/**
 * As the user for the server to connect to
 */
func main() {

	go openLocalConnection()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Connection IP: ")
	server, _ := reader.ReadString('\n')

	sendMsg(server)

}

/**
 * Open a local connection for people to send POST requests to
 */
func openLocalConnection() {
	http.HandleFunc("/", nil)
	http.ListenAndServe(":7777", logRequest(http.DefaultServeMux))
}

/**
 * Log out any POST request infomation made
 */
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			hah, _ := ioutil.ReadAll(r.Body);
			fmt.Println(string(hah)) // {"data":"sss\n"}

			var msg Message
			json.Unmarshal(hah, &msg)

			fmt.Println(msg)

	})
}

/**
 * Send a message to the other person
 */
func sendMsg(server string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Message: ")
	msg, _ := reader.ReadString('\n')

	values := map[string]string{"data": msg}

	jsonValue, _ := json.Marshal(values)

	http.Post("http://localhost:7777", "application/json", bytes.NewBuffer(jsonValue))

	// fmt.Print(msg)

	sendMsg(server)

}
