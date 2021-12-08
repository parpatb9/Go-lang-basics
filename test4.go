package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//Check if the property price is below 400K
func checkPrice(price string) bool {
	intPrice, _ := strconv.Atoi(price)
	if intPrice < 400000 {
		return true
	} else {
		return false
	}
}

//function to check if the properties are PL,CRES and AVE
func checkPlace(price string) bool {

	splitPlace := strings.Split(price, " ")
	arrayPlace := len(splitPlace) - 1
	if splitPlace[arrayPlace] == "PL" || splitPlace[arrayPlace] == "CRES" || splitPlace[arrayPlace] == "AVE" {
		return true
	} else {
		return false
	}

}

//function to check the 10th property
func checkCount(count int) bool {

	if count == 0 {
		return true
	} else if (count % 10) == 1 {
		return false
	} else {
		return true
	}

}

func main() {
	// Open the CSV files
	csvfile, err := os.Open("properties.txt")
	csvfile2, err := os.Open("properties.txt")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	dr := csv.NewReader(csvfile2)
	var propertyAppend = ""
	var checkProperty [1000]string
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
	//declare counter for total 10th property lines
	countenth := 0
	//main for loop for printing the property for test4
	for {
		// Read each property from csv
		countLines++
		property, err := r.Read()
		//Check if the end of line is reached then break the loop
		if err == io.EOF {
			fmt.Printf("No of all entry %d\n", countLines)
			fmt.Printf("No of Duplicate %d\n", dupCount)
			fmt.Printf("Result %d", uniqueLines-countenth)
			break
		}
		//Check for error in reading the CSV file
		if err != nil {
			log.Fatal(err)
		}

		//store to propertyAppend the location, town and date
		propertyAppend = property[1] + property[2] + property[3]
		t := 0
		duplicate := 0

		//loop for checking duplicate
		for {
			t++
			//start comparing if t is morethan i to ensure that it start after the current line
			if t > countLines {
				//Compare the checkProperty array to propertAppend if there are duplicates, Check also if there is an empty row
				if checkProperty[t] == propertyAppend && checkProperty[t] != "" {

					//Increment Counters for checking results
					duplicate++
					dupCount++
					checkProperty[t] = ""

				}

				if t == i {
					break
				}

			}
		}

		//Call checkPrice (test 4.1), checkPlace Functions (test 4.2), Remove all duplicates based on Test 3, and remove blank rows
		if checkPrice(property[4]) && checkPlace(property[1]) && duplicate == 0 && propertyAppend != "" {

			//This to cancel the print of the 10th property (Test 4.3) and continue the count of uniqueLines and increment the counTenth to minus the total uniqueLines
			if checkCount(uniqueLines) == false {
				uniqueLines++
				countenth++
			} else { // Print if the Price is below 400k, the proper places and not the 10th property.
				fmt.Printf("%s\n", property)
				uniqueLines++
			}

		}

	}
}
