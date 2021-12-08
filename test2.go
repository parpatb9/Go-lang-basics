
package main

import (
	
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

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
	var checkProperty [1000] string
	dupCount:=0

	


    // declare Iterate for storing the cvr readings to an array
	i:=0
    //For Loop to store the CSV files for comparing later to check duplicates
	for {
		//Iterate for array and later to be use in max iteration for finding duplicates
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
	countLines:=0
	//declare counter for total result lines 
	uniqueLines:=0
   //main for loop for printing the property for test4
	for {
		// Read each property from csv
		countLines++
		property, err := r.Read()
		//Check if the end of line is reached then break the loop
		if err == io.EOF {
			fmt.Printf("No of all entry %d\n",countLines)
			fmt.Printf("No of Duplicate %d\n",dupCount)
			fmt.Printf("Result %d",uniqueLines)
			break
		}
		//Check for error in reading the CSV file
		if err != nil {
			log.Fatal(err)
		}
		//store to propertyAppend the location, town and date
		propertyAppend = property[1] + property[2] + property[3]
		t:=0
		duplicate:=0;

					//loop for checking duplicate
					
					for {
						t++
						//checkDuplicate reader need to be reset
						

						//start comparing if t is morethan i to ensure that it start after the current line
						if t > countLines{
					
						
							if checkProperty[t] == propertyAppend && checkProperty[t] != ""{

								duplicate++
								dupCount++
								//Deleting the other duplicate entries ahead
								checkProperty[t] = ""

							}
						

						}
						//Break the loop if t is equel to i
						if t == i {
							break
						}

					}

			//Print the first duplicate entry and don't print the empty row		
			if propertyAppend != "" {
				fmt.Printf("%s\n", property)
				uniqueLines++
		
			}


	}
}