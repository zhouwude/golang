// twitter_status.go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* these structs will house the unmarshalled response.
   they should be hierarchically shaped like the XML
   but can omit irrelevant data. */
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

// var user User

func main() {
	// perform an HTTP request for the twitter status of user: Googland
	resp, _ := http.Get("http://twitter.com/users/Googland.xml")
	// initialize the structure of the XML response
	user := User{xml.Name{"", "user"}, Status{""}}
	// unmarshal the XML into our structures
	defer resp.Body.Close() //关掉
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Printf("%s ---", body)
		xml.Unmarshal(body, &user)
	}
	fmt.Printf("name: %s ", user.XMLName)
	fmt.Printf("status: %s", user.Status.Text)
}
