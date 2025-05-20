package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// ioutil.ReadAll will read data from the r *http.Request or file until a error or  EOF is encountered. It returns data in the form of []byte slice
//
//	Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an [InvalidUnmarshalError].
// the below function was widely used previously but after 1.22 + there are new ways to do so

func ParseBodyOldVersion(r *http.Request, x interface{}) {
	io.ReadAll(r.Body) // ioutil.ReadAll is deprecated: As of Go 1.16, this function simply calls [io.ReadAll]
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// ioutil.ReadAll(r.Body):

//     Reads the entire body into memory first, even if it's large.

//     Can cause performance issues or memory pressure under load.

// json.NewDecoder(r.Body).Decode(&user):

//     Streams the body directly into the struct.
// 	   This makes it more efficient, especially for large payloads, and allows for better resource management.

// Uses less memory and is more efficient, especially for large payloads.
func ParseBody(r *http.Request, x interface{}) error {
	defer r.Body.Close() // Ensure the body is closed after reading
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(x)
	if err != nil {
		return err
	}
	return nil
}
