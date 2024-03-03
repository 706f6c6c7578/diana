package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Process the command-line parameters
	textFilePtr := flag.String("t", "", "The file containing the text to be encrypted or decrypted")
	keyFilePtr := flag.String("k", "", "The file containing the key to be used for encryption or decryption")
	decryptPtr := flag.Bool("d", false, "Set this parameter to switch to decryption mode")
	flag.Parse()

	// Check if the text file and the key file were provided
	if *textFilePtr == "" || *keyFilePtr == "" {
		fmt.Println("Error: You must provide both the text file (-t) and the key file (-k).")
		return
	}

	// Read the text and the key from the files
	textBytes, err := ioutil.ReadFile(*textFilePtr)
	if err != nil {
		fmt.Printf("Error: Could not read the text file: %v\n", err)
		return
	}
	keyBytes, err := ioutil.ReadFile(*keyFilePtr)
	if err != nil {
		fmt.Printf("Error: Could not read the key file: %v\n", err)
		return
	}

	// Remove any trailing newline characters
	text := strings.TrimRight(string(textBytes), "\r\n")
	key := strings.TrimRight(string(keyBytes), "\r\n")

	// Check if the key is shorter than the text
	if len(key) < len(text) {
		fmt.Println("Error: The key must not be shorter than the text.")
		return
	}

	// Create an empty map
	trigraph := make(map[string]rune)

	// Initialize the characters
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Fill the map with the positions
	for i := 0; i < len(characters); i++ {
		for j := 0; j < len(characters); j++ {
			key := string(characters[i]) + string(characters[j])
			value := 'A' + ('Z'-'A') - rune(i+j)%26
			trigraph[key] = value
		}
	}

	// Example of using the trigraph for encryption or decryption
	var output string
	if *decryptPtr {
		output = decryptWithTrigraph(text[5:], key[5:], trigraph)
		output = key[:5] + output
	} else {
		output = encryptWithTrigraph(text, key[5:], trigraph)
		output = key[:5] + output
	}
	fmt.Printf("%s\n", output)
}

func encryptWithTrigraph(plaintext string, key string, trigraph map[string]rune) string {
	// Initialize the encrypted text
	var ciphertext strings.Builder

	// Encrypt each character in the plaintext
	for i := 0; i < len(plaintext); i++ {
		// Create the key for the trigraph
		trigraphKey := string(plaintext[i]) + string(key[i])

		// Get the encrypted character from the trigraph
		encryptedChar, exists := trigraph[trigraphKey]
		if !exists {
			fmt.Printf("Error: The key '%s' does not exist in the trigraph.\n", trigraphKey)
			return ""
		}

		// Add the encrypted character to the ciphertext
		ciphertext.WriteRune(encryptedChar)
	}

	// Return the encrypted text
	return ciphertext.String()
}

func decryptWithTrigraph(ciphertext string, key string, trigraph map[string]rune) string {
	// Initialize the plaintext
	var plaintext strings.Builder

	// Decrypt each character in the ciphertext
	for i := 0; i < len(ciphertext); i++ {
		// Create the key for the trigraph
		trigraphKey := string(ciphertext[i]) + string(key[i])

		// Get the decrypted character from the trigraph
		decryptedChar, exists := trigraph[trigraphKey]
		if !exists {
			fmt.Printf("Error: The key '%s' does not exist in the trigraph.\n", trigraphKey)
			return ""
		}

		// Add the decrypted character to the plaintext
		plaintext.WriteRune(decryptedChar)
	}

	// Return the plaintext
	return plaintext.String()
}

