package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(10)
	f, err := os.Open("urls.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	wg := new(sync.WaitGroup)
	num := 1
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		fmt.Println(line)
		wg.Add(1)
		go func(num int, line string) {
			defer wg.Done()
			_ = "breakpoint"
			resp, err := http.Get(line)
			if err != nil {
				panic(err)
			}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(strconv.Itoa(num)+".html", data, 0755)
		}(num, line)
		num++
	}
	wg.Wait()
}
