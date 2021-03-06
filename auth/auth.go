package auth
// as always a work in progres...

import(
	"fmt"
        "bufio"
	"strings"
	"os"
        "time"
	"sync"
	"crypto/sha256"
)

const PASSWD_FILE = "config/passwd"

type User struct {
        uid int32
        username string
        password string
        name string
}

type Auth struct {
        Users map[string]User
        last_update int64
}

func New(system string) (*Auth) {
	this := new(Auth)
       
	if _, err := os.Stat(PASSWD_FILE); err != nil {
        	fmt.Printf("%s!\n", err)
                os.Exit(0)
        }

	if system != "generic" { }
        this.Users = make(map[string]User)
        this.LoadFromFile()

	return this
}

func (self *Auth) LoadFromFile() (bool) {
        f, err := os.Open(PASSWD_FILE)
        if err != nil {
                fmt.Printf("%s!\n", err)
                os.Exit(0)
        }

        defer f.Close()

        r := bufio.NewReader(f)

        for {
                bytes, err := r.ReadBytes('\n')
                if err != nil { break }

                if bytes[0] == '#' { continue }

                line := string(bytes)

                eou := strings.Index(line, ":")
                tmp_user := line[0:eou]

                eop := strings.Index(line, "\n")
                tmp_hash := line[eou + 1:eop];
                var tmp User = User{}
                tmp.username = tmp_user
                tmp.password = tmp_hash
                self.Users[tmp_user] = tmp
        }

        self.last_update = time.Now().Unix()

        return true
}

func (self *Auth) WatchConfigFile() {
	fi, _ := os.Stat(PASSWD_FILE);

	mtime := int64(fi.ModTime().Unix())
	if mtime > self.last_update {
		var mutex = &sync.Mutex{}
		mutex.Lock()
		for k := range self.Users {
			delete(self.Users, k)
		}
		self.LoadFromFile()
		mutex.Unlock()
	}	
}


func (self *Auth) Check(user_guess string, pass_guess string) (bool) {
	self.WatchConfigFile()

	var mutex = &sync.Mutex{}
	mutex.Lock()

        if self.Users[user_guess].username != user_guess {
		return false
        }

	pass_hash := fmt.Sprintf("%x", sha256.Sum256([]byte(pass_guess)))
	if self.Users[user_guess].password != pass_hash  {
		return false
	}

	mutex.Unlock()

        /* Works! */
	return true
}

