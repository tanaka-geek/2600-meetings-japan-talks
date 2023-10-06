package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Define the target binary sequence you want to search for.
	// This stands for "Attribute VBA_Name=\""
	targetHex := "417474726962757400652056425F4E616D0065203D20"
	endHex := "00000000"

	// Convert the target hex string to a byte slice.
	targetBytes, err := hex.DecodeString(targetHex)
	if err != nil {
		log.Fatalf("Error decoding hex: %v", err)
	}

	// Convert the end hex string to a byte slice.
	endBytes, err := hex.DecodeString(endHex)
	if err != nil {
		log.Fatalf("Error decoding hex: %v", err)
	}

	// Specify the path to the binary file you want to search in.
	filePath := "./cs.doc"

	// Read the binary file.
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Search for the targetBytes in the fileData.
	firstIndex := searchBinary(fileData, targetBytes)

	if firstIndex >= 0 {
		// Find the second occurrence of the targetBytes after the first match.
		secondIndex := searchBinary(fileData[firstIndex+len(targetBytes):], targetBytes)

		if secondIndex >= 0 {
			// Adjust the secondIndex to include the length of the targetBytes and offset by the first match.
			secondIndex += firstIndex + len(targetBytes)

			// Search for the endBytes starting from the position after the second match.
			endIndex := searchBinary(fileData[secondIndex+len(targetBytes):], endBytes)

			if endIndex >= 0 {
				// Adjust the endIndex to include the length of the targetBytes and offset by the second match.
				endIndex += secondIndex + len(targetBytes)

				// Replace the matched content with null bytes (0).
				for i := secondIndex; i < endIndex; i++ {
					fileData[i] = 0
				}

				// Create a new file with the result.
				outputFilePath := "./cs_modified.doc"
				err := ioutil.WriteFile(outputFilePath, fileData, 0644)
				if err != nil {
					log.Fatalf("Error writing modified file: %v", err)
				}

				fmt.Printf("Replaced the matched content with null bytes and saved to %s\n", outputFilePath)
			} else {
				fmt.Println("End binary sequence not found after the second match.")
			}
		} else {
			fmt.Println("Second target binary sequence not found in the file.")
		}
	} else {
		fmt.Println("Target binary sequence not found in the file.")
	}
}

func searchBinary(data []byte, target []byte) int {
	for i := 0; i <= len(data)-len(target); i++ {
		if bytesEqual(data[i:i+len(target)], target) {
			return i
		}
	}
	return -1
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
