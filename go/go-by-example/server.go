// run a server on http://localhost:8080
package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func infoHandler(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	escapedPath := html.EscapeString(path)
	method := request.Method
	// ignore errors
	body, _ := ioutil.ReadAll(request.Body)
	request.ParseForm() // need to call this before calling request.Form
	formValues := request.Form
	// %q puts quotes around the string
	// end with \n so curl doesn't show weird % char at the end
	if method == "POST" || method == "PUT" {
		fmt.Fprintf(writer, "%s %s\nBody\n%s\n", method, escapedPath, body)
	} else {
		// Only print "?" near end of URL if there's a query string
		if len(formValues) > 0 {
			fmt.Fprintf(writer, "%s %s?%s\n", method, escapedPath, formValues.Encode())
		} else {
			fmt.Fprintf(writer, "%s %s\n", method, escapedPath)
		}
	}
}

func iLikeHandler(writer http.ResponseWriter, request *http.Request) {
	pathWithoutLeadingSlash := request.URL.Path[1:]
	fmt.Fprintf(writer, "I like %s\n", pathWithoutLeadingSlash)
}

func main() {
	http.HandleFunc("/", iLikeHandler)
	// http.HandleFunc("/", infoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
