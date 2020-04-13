package fetch

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func Fetch(url string) ([]byte, error)  {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("fetcher err %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e;
}