package gosphinx

import (
	"fmt"
	"log"
	"testing"
)

var (
	sc *Client
	//host = "/var/run/searchd.sock"
	host  = "localhost"
	index = "test1"
	words = "test"
)

func init() {
	fmt.Println("Init sphinx client ...")

	// host := "9.138.24.138"
	opts := &Options{
		Host:       host,
		Port:       9312,
		Timeout:    5000,
		MaxMatches: 1000,
		MatchMode:  6,
		Limit:      20,
		GroupSort:  "@group desc",
		Select:     "*",
	}

	sc = NewClient(opts)
	if err := sc.Error(); err != nil {
		log.Fatalf("Init sphinx client > %v\n", err)
	}

	if err := sc.Open(); err != nil {
		log.Fatalf("Init sphinx client > %v\n", err)
	}
}

func TestInitClient(t *testing.T) {

	status, err := sc.Status()
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for _, row := range status {
		fmt.Printf("%20s:\t%s\n", row[0], row[1])
	}

}

func TestQuery(t *testing.T) {
	fmt.Println("Running sphinx Query() test...")
	res, err := sc.Query(words, index, "")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("Query > %#+v\n", res)

	for _, v := range res.Matches {
		t.Logf("id:%d, %#+v\n", v.DocId, v)
	}

	if sc.GetLastWarning() != "" {
		fmt.Printf("Query warning: %s\n", sc.GetLastWarning())
	}

}
