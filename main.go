package main

import (
    t "./trie"

    f "fmt"
    "io/ioutil"
    s "strings"
)

func main() {
    dat, err := ioutil.ReadFile("dict.txt")
    if err != nil {
        panic(err)
    }
    str := string(dat)
    words := s.Split(str, "\n")

    trie := t.MakeTrie(words)

    for _, word := range words {
        if !trie.Find(word) {
            f.Println("couldn't find", word)
        }
    }
    f.Println("finished")
}
