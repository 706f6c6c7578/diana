package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Process the command-line parameters
	textPtr := flag.String("t", "", "The text to be encrypted or decrypted")
	keyPtr := flag.String("k", "", "The key to be used for encryption or decryption")
	decryptPtr := flag.Bool("d", false, "Set this parameter to switch to decryption mode")
	flag.Parse()

	// Check if the text and the key were provided
	if *textPtr == "" || *keyPtr == "" {
		fmt.Println("Error: You must provide both the text (-t) and the key (-k).")
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
		output = decryptWithTrigraph(*textPtr, *keyPtr, trigraph)
	} else {
		output = encryptWithTrigraph(*textPtr, *keyPtr, trigraph)
	}
	fmt.Printf("\n            *FOR EDUCATIONAL PURPOSES*\n*DESTROY THE PAD SECURELY AFTER USING THE PROGRAM*\n\nInput: %s\nKey: %s\nOutput: %s\n", *textPtr, *keyPtr, output)
}

func encryptWithTrigraph(plaintext string, key string, trigraph map[string]rune) string {
	// Make sure the plaintext and the key have the same length
	if len(plaintext) != len(key) {
		fmt.Println("Error: The plaintext and the key must have the same length.")
		return ""
	}

	// Convert the plaintext and the key to uppercase
	plaintext = strings.ToUpper(plaintext)
	key = strings.ToUpper(key)

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
	// Make sure the ciphertext and the key have the same length
	if len(ciphertext) != len(key) {
		fmt.Println("Error: The ciphertext and the key must have the same length.")
		return ""
	}

	// Convert the ciphertext and the key to uppercase
	ciphertext = strings.ToUpper(ciphertext)
	key = strings.ToUpper(key)

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

