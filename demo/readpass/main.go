package main

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

func main() {
	main2()
}
func main3() {
	fmt.Printf("Enter silent password: ")
	silentPassword,_ := gopass.GetPasswd() // Silent

	// input is in byte and need to convert to string
	// for storing and comparison
	fmt.Println(string(silentPassword))

	fmt.Printf("Enter masked password: ")
	maskedPassword,_ := gopass.GetPasswdMasked() // Masked
	fmt.Println(string(maskedPassword))

}

func main2() {
	username, password := credentials()
	fmt.Printf("Username: %s, Password: %s\n", username, password)
}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}