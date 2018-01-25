package main

import (
	"fmt"
	"net/http"
	"bufio"
	"os"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"flag"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

/**
 * My data strucutre {data: 'string'}
 */
type Message struct {
	Data string
}

/**
 * As the user for the server to connect to
 */
func main() {

	// Grab the set flag port
	port := flag.Int("port", 7777, "Please enter a port")
	flag.Parse()

	// Convert port to a string
	var portStr = strconv.Itoa(*port);

	// Run in parallel - Open the server
	go openLocalConnection(portStr)

	// Tell the user who they are
	// @question why cant this sit inside the go function above ?
	color.HiGreen("My Address: http://localhost:" + portStr)

	// Input partners address
	color.HiRed("Partners Address:")
	reader := bufio.NewReader(os.Stdin)
	partner, _ := reader.ReadString('\n')

	// Send the msg
	sendMsg(partner)

}

/**
 * Open a local connection for people to send POST requests to
 */
func openLocalConnection(port string) {
	http.HandleFunc("/", nil)
	http.ListenAndServe(string(":" + port), logRequest(http.DefaultServeMux))
}

/**
 * Log out any POST request infomation made
 */
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Ready the body
			body, _ := ioutil.ReadAll(r.Body);

			// It it though my struct
			var msg Message
			json.Unmarshal(body, &msg)

			// Display incoming message
			fmt.Println("INCOMING: " + msg.Data)

	})
}

/**
 * Send a message to the other person
 */
func sendMsg(partner string) {

	// Get the message from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Message: ")
	msg, _ := reader.ReadString('\n')

	// Run msg through struct
	values := Message{
		Data: msg,
	}
	jsonValue, _ := json.Marshal(values)

	// Send message via post // @TODO working with string as other IP here
	http.Post(strings.TrimSuffix(partner, "\n"), "application/json", bytes.NewBuffer(jsonValue))

	// Get ready to send the next message
	sendMsg(partner)

}
