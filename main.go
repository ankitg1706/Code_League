package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Entry point for the web server
func main() {

	http.HandleFunc("/echo", handleEcho)
	http.HandleFunc("/invert", handleInvert)
	http.HandleFunc("/flatten", handleFlatten)
	http.HandleFunc("/sum", handleSum)
	http.HandleFunc("/multiply", handleMultiply)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

// Handle the /echo endpoint
func handleEcho(w http.ResponseWriter, r *http.Request) {
	matrix, err := parseCSVFile(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse CSV file: %v", err), http.StatusBadRequest)
		return
	}
	writeMatrixResponse(w, matrix)
}

// Handle the /invert endpoint
func handleInvert(w http.ResponseWriter, r *http.Request) {
	matrix, err := parseCSVFile(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse CSV file: %v", err), http.StatusBadRequest)
		return
	}
	inverted := invertMatrix(matrix)
	writeMatrixResponse(w, inverted)
}

// Handle the /flatten endpoint
func handleFlatten(w http.ResponseWriter, r *http.Request) {
	matrix, err := parseCSVFile(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse CSV file: %v", err), http.StatusBadRequest)
		return
	}
	var flattened []string
	for _, row := range matrix {
		flattened = append(flattened, row...)
	}
	fmt.Fprint(w, strings.Join(flattened, ","))
}

// Handle the /sum endpoint
func handleSum(w http.ResponseWriter, r *http.Request) {
	matrix, err := parseCSVFile(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse CSV file: %v", err), http.StatusBadRequest)
		return
	}
	sum, err := sumMatrix(matrix)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to calculate sum: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, sum)
}

// Handle the /multiply endpoint
func handleMultiply(w http.ResponseWriter, r *http.Request) {
	matrix, err := parseCSVFile(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse CSV file: %v", err), http.StatusBadRequest)
		return
	}
	product, err := multiplyMatrix(matrix)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to calculate product: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, product)
}

// Parse the uploaded CSV file into a 2D slice of strings
func parseCSVFile(r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve file: %w", err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}
	return records, nil
}

// Invert the matrix by transposing it
func invertMatrix(matrix [][]string) [][]string {
	n := len(matrix)
	inverted := make([][]string, n)
	for i := range inverted {
		inverted[i] = make([]string, n)
		for j := range inverted[i] {
			inverted[i][j] = matrix[j][i]
		}
	}
	return inverted
}

// Sum the elements of the matrix
func sumMatrix(matrix [][]string) (int, error) {
	sum := 0
	for _, row := range matrix {
		for _, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				return 0, fmt.Errorf("invalid integer value: %v", err)
			}
			sum += num
		}
	}
	return sum, nil
}

// Multiply the elements of the matrix
func multiplyMatrix(matrix [][]string) (int, error) {
	product := 1
	for _, row := range matrix {
		for _, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				return 0, fmt.Errorf("invalid integer value: %v", err)
			}
			product *= num
		}
	}
	return product, nil
}

// Write a matrix response as a string in matrix format
func writeMatrixResponse(w http.ResponseWriter, matrix [][]string) {
	var response strings.Builder
	for _, row := range matrix {
		response.WriteString(strings.Join(row, ","))
		response.WriteString("\n")
	}
	fmt.Fprint(w, response.String())
}
