package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	insert_user := fmt.Sprintf("%s:%x\n", string(user), sum)

	os.Mkdir("config", 0755)

	existing := fmt.Sprintf("%s:", user)

	new_user := true

	var output string
	m, err := os.Open("config/passwd")
	if err == nil {
		r = bufio.NewReader(m)
		for {
			bytes, err := r.ReadBytes('\n')
			if err != nil { break }

			line := string(bytes)
			if strings.HasPrefix(line, existing) {
				output = output + insert_user
				new_user = false
			} else {
				output = output + line
			}
		}

		m.Close()
	}

	if (new_user) {
		output += insert_user
	}

	f, _ := os.Create("config/passwd")
	f.Write([]byte(output))
	f.Close()
}

