package main

import (
	"bufio"
	"fmt"
	"os"
	"crypto/sha256"
)


func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	user, _ := r.ReadBytes('\n')

	user = user[0:len(user) -1]
	fmt.Print("Pass: ")
	pass, _ := r.ReadBytes('\n')
	pass = pass[0:len(pass) - 1] 

	sum := sha256.Sum256(pass)

	output := fmt.Sprintf("%s:%x\n", string(user), sum)

	os.Mkdir("config", 0755)

	f, _ := os.Create("config/passwd")
	f.Write([]byte(output))
	f.Close()
}

