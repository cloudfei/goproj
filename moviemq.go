package main

import (
	"fmt"
	logging "github.com/op/go-logging"
	//"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func f(data string) {
	fmt.Println(data)
}

func heartbreakerHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	data := req.URL.String()
	go f(data)
	//io.WriteString(w, data)
}

var LISTENING_PORT = 8000
var logger = logging.MustGetLogger("hugmachine.log")

func main() {
	logging.NewLogBackend(os.Stderr, "", 0)
	http.HandleFunc("/heartbreaker", heartbreakerHandler)
	logger.Infof("Listening on port %d", LISTENING_PORT)
	err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(LISTENING_PORT), nil)
	if err != nil {
		logger.Fatal("ListenAndServe: " + err.Error())
	}
}
