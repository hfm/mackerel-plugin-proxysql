package mpproxysql

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Config for ProxySQL
type Config struct {
	adminCred string
}

// ReadConfig reads the file `f` and parses its content to create
// a new Schema object
func ReadConfig(f string) Config {
	c, err := os.Open(f)
	if err != nil {
		log.Fatalln("Cannot open: ", err)
	}
	defer c.Close()

	return Parse(c)
}

// Parse parses into `in` and returns its content into create a new Config object
func Parse(in io.Reader) Config {
	c := Config{}
	s := bufio.NewScanner(in)
	var l string
	for s.Scan() {
		l = s.Text()
		if strings.Contains(l, "admin_credentials") {
			reg := regexp.MustCompile("admin_credentials\\s*=\\s*\"([^:]*:[^:]*)(;.+:.+)?\"")
			match := reg.FindStringSubmatch(l)
			c.adminCred = match[1]
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalln("Cannot parse: ", err)

	}

	return c
}
