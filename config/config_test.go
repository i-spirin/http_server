package config_test

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/i-spirin/http_server/config"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// func main() {
//     rand.Seed(time.Now().UnixNano())

//     fmt.Println(randSeq(10))
// }

func TestParse(t *testing.T) {
	filename := "/tmp/test_config" + randSeq(10) + ".yaml"

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Fatalf("Cannot create sample config file: %v", err)
	}

	defer os.Remove(filename)
	defer file.Close()

	file.Write([]byte("bind_host: 1.2.3.4\nbind_port: 1234\n"))

	conf := config.Config{}
	err = conf.Parse(filename)
	if err != nil {
		t.Fatalf("Error parsing config: %v", err)
	}

	if conf.BindHost != "1.2.3.4" {
		t.Fatalf("BindHost not equal to 1.2.3.4, got: %v", conf.BindHost)
	}

	if conf.BindPort != 1234 {
		t.Fatalf("BindPort not equal to 1234, got: %v", conf.BindPort)
	}
}

func TestParse2(t *testing.T) {
	filename := "/tmp/test_config" + randSeq(10) + ".yaml"

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Fatalf("Cannot create sample config file: %v", err)
	}

	defer os.Remove(filename)
	defer file.Close()

	file.Write([]byte("bind_host 1.2.3.4\nbind_port: 1234\n"))

	conf := config.Config{}
	err = conf.Parse(filename)

	if err == nil {
		t.Fatalf("Parse does not returned error during checking YAML syntax")
	}

}

func TestParse3(t *testing.T) {
	conf := config.Config{}
	err := conf.Parse("/tmp/" + randSeq(10))
	if err == nil {
		log.Fatalf("Parse does not returned error while opens non-existing file")
	}
}
