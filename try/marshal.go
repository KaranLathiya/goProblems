package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string `json:"birdType"`
	// we can set the "omitempty" property as part of the JSON tag
	Description string `json:"what it does,omitempty"`
}

// If we want to always ignore a field, we can use the json:"-" struct tag to denote that we never want this field included:

// type Bird struct {
// 	Species     string `json:"-"`
// }

func main() {
	pigeon := Bird{
		Species: "Pigeon",
	}

	data, _ := json.Marshal(pigeon)

	fmt.Println(string(data))
}

// pigeon := &Bird{
// 	Species:     "Pigeon",
// 	Description: "likes to eat seed",
//   }

//   // Now we pass a slice of two pigeons
//   data, _ := json.Marshal([]*Bird{pigeon, pigeon})
//   fmt.Println(string(data))

//   This will give the output:

// [{"birdType":"Pigeon","what it does":"likes to eat seed"},{"birdType":"Pigeon","what it does":"likes to eat seed"}]

// bird := Bird{
// 	Species: "pigeon",
// 	Description: "likes to eat seed",
// }

// // The second parameter is the prefix of each line, and the third parameter
// // is the indentation to use for each level
// data, _ := json.MarshalIndent(bird, "", "  ")
// fmt.Println(string(data))
// The output in this case will be a formatted JSON string, instead of the compressed one-liner response that we’ve been getting so far:

// {
// 	"Species": "pigeon",
// 	"Description": "likes to eat seed"
// }

type Dimensions struct {
	Height int
	Width  int
}

// // marshals a Dimensions struct into a JSON string
// // with format "heightxwidth"
// func (d Dimensions) MarshalJSON() ([]byte, error) {
// 	return []byte(fmt.Sprintf(`"%dx%d"`, d.Height, d.Width)), nil
// }

// func main() {
// 	bird := Bird{
// 		Species: "pigeon",
// 		Dimensions: Dimensions{
// 			Height: 24,
// 			Width:  10,
// 		},
// 	}
// 	birdJson, _ := json.Marshal(bird)
// 	fmt.Println(string(birdJson))
// 	// {"Species":"pigeon","Dimensions":"24x10"}
// }

// func main() {
// 	// The keys need to be strings, the values can be
// 	// any serializable value
// 	birdData := map[string]any{
// 		"birdSounds": map[string]string{
// 			"pigeon": "coo",
// 			"eagle":  "squawk",
// 		},
// 		"total birds": 2,
// 	}

// 	// JSON encoding is done the same way as before
// 	data, _ := json.Marshal(birdData)
// 	fmt.Println(string(data))
// }
