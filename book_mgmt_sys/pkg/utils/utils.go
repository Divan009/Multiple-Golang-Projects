package utils

import (
	"encoding/json"
	"io"
	//"io/ioutil"
	"net/http"
)

//code to unmarshall json request which we receive from request to pass it to db

//func ParseBody(r *http.Request, x interface{}) {
//	if body, err :=  io.ReadAll(r.Body); err == nil{
//		if err := json.Unmarshal([]byte(body), x): err != nil{
//			return
//		}
//	}
//}

// ParseBody reads the body from an HTTP request and unmarshals it into the provided interface
func ParseBody(r *http.Request, x interface{}) error {
	// Read the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err // Return the error if reading fails
	}

	// Unmarshal the JSON body into the provided interface
	if err := json.Unmarshal(body, x); err != nil {
		return err // Return the error if unmarshalling fails
	}

	return nil // Return nil if everything is successful
}
