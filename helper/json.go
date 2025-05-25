package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseToJson(write http.ResponseWriter, response any) {
	// add content type header response
	write.Header().Add("Content-Type", "application/json")
	// create encoder to encode apiResponse variable, then return into http response
	encode := json.NewEncoder(write)
	if err := encode.Encode(response); err != nil {
		panic(err)
	}
}
