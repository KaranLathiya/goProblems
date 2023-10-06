// Go program to illustrate the
// modification concept in map

package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	fmt.Println("Original map: ", m_a_p)

	// Assigned the map into a new variable
	new_map := m_a_p

	// Perform modification in new_map
	new_map[96] = "Parrot"
	new_map[98] = "Pig"
	// Deleting keys
	// Using delete function
	delete(m_a_p, 90)
	delete(m_a_p, 93)
	// Checking the key is available
	// or not in the m_a_p map
	pet_name, ok := m_a_p[90]
	fmt.Println("\nKey present or not:", ok)
	fmt.Println("Value:", pet_name)

	// Display after modification
	fmt.Println("New map: ", new_map)
	fmt.Println("\nModification done in old map:\n", m_a_p)
}
