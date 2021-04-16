package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

var errMatrixEmpty = errors.New("invalid input: empty matrix\n")
var errMatrixNotSquare = errors.New("invalid input: matrix is not square\n")
var errMatrixInvalidNumber = errors.New("invalid input: matrix has invalid number format\n")

//A Matrix could have any dimension where the number of rows are equal to the number of columns (square), and each value is an integer.
type Matrix struct {
	records [][]string
}

//isValid will check the member "records":
//It cannot be empty.
//It must be square matrix.
//All elements must be integer and not exceed MaxInt.
func (m *Matrix) isValid() error {
	if m.records == nil || len(m.records) == 0 {
		return errMatrixEmpty
	}
	var matrixSize int = len(m.records)
	for _, row := range m.records {
		var rowSize int = len(row)
		if rowSize != matrixSize {
			return errMatrixNotSquare
		}
		for _, element := range row {
			if _, err := strconv.Atoi(strings.Trim(element, " ")); err != nil {
				errMsg := errMatrixInvalidNumber.Error() + "\tCaused by: " + err.Error() + "\n"
				return errors.New(errMsg)
			}
		}
	}
	return nil
}

//Return the matrix as a string in matrix format where the columns and rows are inverted
func (m *Matrix) invert() string {
	matrixSize := len(m.records)
	rds := make([][]string, matrixSize)
	for i := 0; i < matrixSize; i++ {
		rds[i] = make([]string, matrixSize)
	}
	for i, row := range m.records {
		for j, element := range row {
			rds[j][i] = strings.Trim(element, " ")
		}
	}
	var result string
	for _, row := range rds {
		result = fmt.Sprintf("%s%s\n", result, strings.Join(row, ","))
	}
	return result
}

//Return the matrix as a 1 line string, with values separated by commas.
func (m *Matrix) flatten() string {
	var result string
	for _, row := range m.records {
		result = fmt.Sprintf("%s%s,", result, strings.Join(row, ","))
	}
	return result[:len(result)-1] + "\n"
}

//Return the sum of the integers in the matrix
func (m *Matrix) sum() string {
	result := big.NewInt(0)
	for _, row := range m.records {
		for _, num := range row {
			n, _ := strconv.ParseInt(strings.Trim(num, " "), 10, 64)
			result.Add(result, big.NewInt(n))
		}
	}
	return result.String() + "\n"
}

//Return the product of the integers in the matrix
func (m *Matrix) multiply() string {
	result := big.NewInt(1)
	for _, row := range m.records {
		for _, num := range row {
			n, err := strconv.ParseInt(strings.Trim(num, " "), 10, 64)
			if err != nil {
				fmt.Print(err.Error())
			}
			result.Mul(result, big.NewInt(n))
		}
	}
	return result.String() + "\n"
}

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}

		var response string
		m := Matrix{records: records}
		err = m.isValid()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err.Error())))
			return
		}
		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
		fmt.Fprint(w, response)
	})

	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		m := Matrix{records: records}
		err = m.isValid()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err.Error())))
			return
		}
		response = m.invert()
		fmt.Fprint(w, response)
	})

	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		m := Matrix{records: records}
		err = m.isValid()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err.Error())))
			return
		}
		response = m.flatten()
		fmt.Fprint(w, response)
	})

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		m := Matrix{records: records}
		err = m.isValid()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err.Error())))
			return
		}
		response = m.sum()
		fmt.Fprint(w, response)
	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		m := Matrix{records: records}
		err = m.isValid()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err.Error())))
			return
		}
		response = m.multiply()
		fmt.Fprint(w, response)
	})

	http.ListenAndServe(":8080", nil)
}
