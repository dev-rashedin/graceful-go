/* Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

*/

// Map is similar to dictionary in python and object in javaScript, hashMap in java

package main

import "fmt"


func main(){
	fmt.Println("-----Maps in go-----")
	fmt.Println(" ")
	// Initialize the map in go
	var myMap map[string]string; // declare the map
	
	myMap = make(map[string]string );

	fmt.Println("Map value is: ", myMap)

	// Declare and initialize a map
	countryAndCapital := make(map[string]string)
  countryAndCapital["Bangladesh"] = "Dhaka"
	countryAndCapital["India"] = "New Delhi"
	countryAndCapital["USA"] = "Washington, D.C."
	countryAndCapital["France"] = "Paris"

	// Extract the values
  for country := range countryAndCapital {
		fmt.Println("The capital of", country, "is", countryAndCapital[country])
	}

	// Extract the keys and values
	for country, capital := range countryAndCapital {
		fmt.Println("The capital of", country, "is", capital)
	}

	// Check if the key exists
 capital, ok := countryAndCapital["Japan"]

 if ok {
 	fmt.Println("The capital of Japan is", capital)
 } else {
 	fmt.Println("Japan is not in the map")
 }

//  Delete an element from the map
delete(countryAndCapital, "India")

fmt.Println("Entry for India is deleted")

// Update the map state
fmt.Println("Updated Map: ")

countryAndCapital["USA"] = "Alaska"

for country, capital := range countryAndCapital{
	fmt.Println("The capital of", country, "is", capital)
}
}