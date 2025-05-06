// Lets implement a simple  HTTP server using the built in net/http package

package server

// We need to import the net/http package for this, but lets worry about that later
// Importing a logging utility would be a good idea too for debugging our outputs

import (
	"net/http" // This is the package that provides HTTP client and server implementations, we will worry about this later
)

// Lets assume for now that all our server will do take an http request and return a string that greets the user e.g. "Hello, {request.Name}!"
// Lets start by thinking of an interface that will define the methods that our server will implement

type HelloServer interface {
	// This method will take a string and return a string
	// It will be used to greet the user
	Greet(name string) string
	// This method will take an http.Request and return an http.Response
	// It will be used to handle the http request and return a response
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

// Interfaces are a powerful way to define the behavior of our server
// They allow us to define a contract that our server must adhere to
// We can define a struct that implements this interface

type SimpleHelloServer struct {
	// This struct will implement the HelloServer interface
	// We can add any fields that we want here
	// Structs are a way to group related data together
	// For example, we can add a field for the server name
	ServerName string
}

// Now lets implement our first method, Greet
// Since this method is part of the SimpleHelloServer struct, we need to use a pointer receiver
// This allows us to modify the struct fields if we want to

// Try implementing the method and see if it works by running go test
// Greet should takes a name and returns a greeting string that includes the server name as follows:

// PASS: Greet("Ethan") should return "Hello Ethan from SimpleHelloServer!"
// HINT: You can use the fmt package to format the string, but you'll need to import it
// Bonus points if you include error checking in your implementation
func (s *SimpleHelloServer) Greet(name string) string {
	return ""
}
