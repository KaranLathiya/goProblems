package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func main() {
	data, err := ioutil.ReadAll(resp.Body)

	if err == nil && data != nil {
		err = json.Unmarshal(data, value)
	}

	err = json.NewDecoder(resp.Body).Decode(value)

	birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird Bird
	if !json.Valid([]byte(birdJson)) {
		// handle the error here
		fmt.Println("invalid JSON string:", birdJson)
		return
	}
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird.Species)
	// {pigeon likes to perch on rocks}
}

// func main() {
// 	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
// 	var result map[string]any
// 	json.Unmarshal([]byte(birdJson), &result)

// 	// The object stored in the "birds" key is also stored as
// 	// a map[string]any type, and its type is asserted from
// 	// the `any` type
// 	birds := result["birds"].(map[string]any)

// 	for key, value := range birds {
// 		// Each value is an `any` type, that is type asserted as a string
// 		fmt.Println(key, value.(string))
// 	}
// 	//pigeon likes to perch on rocks
// 	//eagle bird of prey
// }
