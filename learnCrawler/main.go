package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	//"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"

	"golang.org/x/text/transform"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//resp, err := http.Get("http://www.newsmth.net/about/index.html")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("err:", resp.StatusCode)
	}
//gbk转化为utf8方法一
	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//gbk转化为utf8方法2
	e := determineEncoding(resp.Body)
	utf8Reader = transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	//all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("###%s\n###", all)
	printAllcity(all)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		panic(err)
	}

	e,_,_ := charset.DetermineEncoding(bytes, "")
	fmt.Printf("=====%+v\n", e)
	return e
}

func printAllcity(content[] byte)  {
	res := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]*)"[^>]*([^<]+)</a>`)
	match := res.FindAllSubmatch(content, -1)

	for _, m := range match {
		fmt.Printf("city:%s, url:%s\n", m[2], m[1])
	}
	fmt.Printf("matched found %d", len(match))
}