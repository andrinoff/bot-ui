package main

import (
	discord_bot "bot-ui/discord"
	get_conf "bot-ui/func"
	read_message "bot-ui/message"
	telegram_bot "bot-ui/telegram"
	"fmt"
	"os"
)


func main() {
	get_conf.Read()
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path_to_markdown_file>")
		os.Exit(1)
	}

	// The file path is the second argument.
	msgPath := os.Args[1]
	imgPath := os.Args[2]

	// --- 2. Read File Content into a String ---
	// Call our new function to get the file content.
	contentString, err := read_message.Read(msgPath)
	if err != nil {
		// If the function returns an error, print it and exit.
		fmt.Println(err)
		os.Exit(1)
	}

	// --- 3. Print the Entire String ---
	// Print a header and then the entire file content.
	fmt.Printf("--- Content of %s ---\n", msgPath)
	fmt.Println(contentString)
	fmt.Println("--- End of File ---")

	// telegram_bot.Send(contentString)
	discord_bot.Send(contentString, imgPath)
	telegram_bot.Send(contentString, imgPath)
	
}