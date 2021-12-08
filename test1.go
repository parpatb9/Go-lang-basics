package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("properties.txt")
	csvfile2, err := os.Open("properties.txt")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	dr := csv.NewReader(csvfile2)
	var propertyAppend = ""
	var checkProperty [300]string
	dupCount := 0

	// declare Iterate for storing the cvr readings to an array
	i := 0
	//For Loop to store the CSV files for comparing later to check duplicates
	for {
		i++
		checkDuplicate, err2 := dr.Read()
		if err2 == io.EOF {
			break

		}
		if err2 != nil {
			log.Fatal(err)
		}
		//store to checkProperty the location, town and date. This is for comparing later to detect duplicates
		checkProperty[i] = checkDuplicate[1] + checkDuplicate[2] + checkDuplicate[3]

	}

	//declare counter for total read lines
	countLines := 0
	//declare counter for total result lines
	uniqueLines := 0
	for {
		//count all the iteration of for loop
		countLines++
		// Read each property from csv
		property, err := r.Read()
		if err == io.EOF {
			fmt.Printf("No of all entry %d\n", countLines)
			fmt.Printf("No of Duplicate %d\n", dupCount)
			fmt.Printf("Result %d", uniqueLines)
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		propertyAppend = property[1] + property[2] + property[3]
		t := 0
		duplicate := 0

		//loop for checking duplicate

		for {
			t++
			//checkDuplicate reader need to be reset

			//start comparing if t is morethan countLines to ensure that it start after the current line
			if t > countLines {

				if checkProperty[t] == propertyAppend && checkProperty[t] != "" {

					duplicate++
					dupCount++

				}

			}

			if t == i {
				break
			}

		}

		//if no duplicate print property and print the last entry of duplication
		if duplicate == 0 {
			fmt.Printf("%s\n", property)
			uniqueLines++
		}

	}
}
