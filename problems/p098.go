package main


import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
    "sort"
)

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func Permutations(items []int, n int) [][]int {
    // permute from 9-0 for size times
    perms := make([][]int, 0)
    if n == 0 {
        return perms
    }
    if n == 1 {
        for _, item := range items {
            p := make([]int, 0)
            p = append(p, item)
            perms = append(perms, p)
        }
    } else {
        for i, item := range items {
            // copying the items and remove the current index
            remainingItems := RemoveIndex(items, i)
            ps := Permutations(remainingItems, n - 1)
            for _, p := range ps {
                p = append([]int{item}, p...)
                perms = append(perms, p)
            }
        }
    }

    return perms
}

func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}


func main() {
	filename := "p098_words.txt"

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(data), ",")

    // find all the anagrams in words
    anagrams := make(map[string][]string)
	for _, word := range words {
        trimmed := strings.Trim(word, "\"")
        letters := SortString(trimmed)
        anagrams[letters] = append(anagrams[letters], trimmed)
	}

    // Drop words without anagrams
    groups := make([][]string, 0)
    keysToKeep := make([]string, 0)
    for key, pairs := range anagrams {
        if len(pairs) > 1 {
            groups = append(groups, pairs)
            keysToKeep = append(keysToKeep, key)
        }
    }
    fmt.Println(keysToKeep)

    // Sort anagrams by length
    sort.Slice(keysToKeep, func(i, j int) bool {
        return len(keysToKeep[i]) > len(keysToKeep[j])
    })

    // for each key, assign a value for each number
    // we need to permute all the numbers 0-9
    ps := Permutations([]int{0,1,2}, 2)
    fmt.Println(ps)


    /*
    for _, group := range groups {
        // fmt.Println(group)
        for _, letter := range group {


        }
    }
    */

    // For each group, assign a letter mapping and see if perfect-squares are formed.
    // TODO

}
