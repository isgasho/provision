package midlayer

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestStaticFiles(t *testing.T) {

	hh := ServeStatic(":3235235", ".", log.New(os.Stderr, "", log.LstdFlags))
	if hh != nil {
		if hh.Error() != "listen tcp: address 3235235: invalid port" {
			t.Errorf("Expected a different error: %v", hh.Error())
		}
	} else {
		t.Errorf("Should have returned an error")
	}

	go ServeStatic(":32134", ".", log.New(os.Stderr, "", log.LstdFlags))

	response, err := http.Get("http://127.0.0.1:32134/dhcp.go")
	count := 0
	if err != nil && count < 10 {
		t.Logf("Failed to get file: %v", err)
		time.Sleep(1 * time.Second)
		count++
	}
	if count == 10 {
		t.Errorf("Should have served the file: missing content")
	}
	buf, _ := ioutil.ReadAll(response.Body)
	if !strings.HasPrefix(string(buf), "package midlayer") {
		t.Errorf("Should have served the file: missing content")
	}

}