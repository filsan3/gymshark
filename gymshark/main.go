package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// calculatePacks takes the number of items ordered
// and returns a map indicating the packs needed for each pack size.
func calculatePacks(itemsOrdered int) map[int]int {
	// List of available pack sizes.
	packSizes := []int{5000, 2000, 1000, 500, 250}
	// Map to store the count of each pack size needed.
	packsNeeded := make(map[int]int)

	// Iterate through each pack size and calculate the number of packs needed.
	for _, packSize := range packSizes {
		// Calculate how many full packs of the current size are needed.
		fullPacks := itemsOrdered / packSize
		if fullPacks > 0 {
			packsNeeded[packSize] = fullPacks
			itemsOrdered -= fullPacks * packSize
		}
	}

	// If there are still any remaining items, we need an additional pack of size 250.
	if itemsOrdered > 0 {
		packsNeeded[250] = 1
	}

	return packsNeeded
}


// indexHandler handles requests to the root ("/") and serves the index.html file.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// calculateHandler handles requests to "/calculate" and calculates the packs needed
// based on the number of items provided in the request. It then sends the result back to the client.
func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the number of items from the form data.
	itemsStr := r.FormValue("items")
	items, err := strconv.Atoi(itemsStr)

	// Check for errors in parsing the input.
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Calculate the packs needed.
	packsNeeded := calculatePacks(items)

	// Display the result on the client's browser.
	fmt.Fprintf(w, "Items ordered: %d\n", items)
	fmt.Fprintf(w, "Packs needed: %d\n", packsNeeded)
}

func main() {
	// Register handlers for different routes.
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calculate", calculateHandler)

	// Start the server and print a message indicating its location.
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
