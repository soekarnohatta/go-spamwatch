// Package example provides example to the user.
package example

import (
	"fmt"
	"github.com/SoekarnoHatta/go-spamwatch/spamwatch"
	"os"
	"strconv"
)

func main() {
	// Initialize token
	token := "PUT YOUR TOKEN HERE"

	// Initialize New Client
	client, _ := spamwatch.NewClient("", token)

	// Start requesting to the API
	bannedUser, _ := client.GetBan(123456789)
	if bannedUser != nil {
		fmt.Println("UserID:" + strconv.Itoa(bannedUser.Id))
		fmt.Println("Reason:" + bannedUser.Reason)
		fmt.Println("Message" + bannedUser.Message)
	}
}
