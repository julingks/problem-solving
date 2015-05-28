package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

type Rank struct {
	Name  string
	Score int
}

var (
	mux   *sync.Mutex = &sync.Mutex{}
	words []string    = []string{"sun", "hello", "world"}
	ranks []Rank
)

type ByScore []Rank

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

func handleConnect(conn net.Conn) {
	rand.Seed(time.Now().Unix())
	life := 20
	score := 0

	w := bufio.NewWriter(conn)
	r := bufio.NewReader(conn)
	w.WriteString("Welcome to hangman\n")
	w.WriteString("*******************************\n")
	w.WriteString("What is your name ? ")
	w.Flush()

	// read name
	name, _ := r.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println(name)
	w.WriteString(fmt.Sprintln("Hello, ", name))

	// word & tword init
	word := words[rand.Intn(len(words))]
	tword := strings.Repeat("_", len(word))

	for life > 0 {
		w.WriteString(fmt.Sprintln("Word: ", tword))
		w.WriteString(fmt.Sprintf("(LIFE: %d SCORE: %d ) Your guess > ", life, score))
		w.Flush()
		str, _ := r.ReadString('\n')
		str = strings.TrimSpace(str)
		if len(str) > 1 {
			w.WriteString(fmt.Sprintln("Input only one character!"))
			w.Flush()
			continue
		}

		if strings.Contains(word, str) && !strings.Contains(tword, str) {
			for i, s := range word {
				if string(s) == str {
					tword = tword[:i] + string(s) + tword[i+1:]
				}
			}
		} else {
			w.WriteString(fmt.Sprintf("'%s' is not in the word T_T\n", str))
			life--
		}

		if word == tword {
			w.WriteString(fmt.Sprintln("Congratulations! The word is ", word))
			w.WriteString(fmt.Sprintln("Start new game:)"))
			score += len(word)
			word = words[rand.Intn(len(words))]
			tword = strings.Repeat("_", len(word))
		}
	}
	w.WriteString(fmt.Sprintf("You get %d points\n", score))
	w.WriteString(fmt.Sprintln("********* R A N K I N G ************"))

	mux.Lock()

	tempRanks := append(ranks, Rank{name, score})
	sort.Sort(ByScore(tempRanks))
	if len(tempRanks) <= 5 {
		ranks = tempRanks
	} else {
		ranks = tempRanks[:5]
	}

	mux.Unlock()

	for i, v := range ranks {
		if i > 4 {
			break
		}
		w.WriteString(fmt.Sprintf("Rank %d : ( %d pt ) %s\n", i+1, v.Score, v.Name))
	}
	w.Flush()
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8887")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnect(conn)
	}
}
