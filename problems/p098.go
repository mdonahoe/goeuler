package main

/*
By replacing each of the letters in the word CARE with 1, 2, 9, and 6 respectively, we form a square number: 1296 = 36^2.
What is remarkable is that, by using the same digital substitutions, the anagram, RACE, also forms a square number: 9216 = 96^2.
We shall call CARE (and RACE) a square anagram word pair and specify further that leading zeroes are not permitted,
neither may a different letter have the same digital value as another letter.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words,
find all the square anagram word pairs (a palindromic word is NOT considered to be an anagram of itself).

What is the largest square number formed by any member of such a pair?
*/

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
    fmt.Println("Permutations", items, n)
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
            remainingItems := items
            fmt.Println("remain = ", remainingItems, "items = ", items)
            remainingItems = CopyRemoveIndex(remainingItems, i)
            fmt.Println("afterremain = ", remainingItems, "items = ", items)
            fmt.Println("range ", i, item, remainingItems, items, "n=", n)
            ps := Permutations(remainingItems, n - 1)
            for _, p := range ps {
                p = append([]int{item}, p...)
                perms = append(perms, p)
            }
        }
    }

    return perms
}

func CopyRemoveIndex(s []int, index int) []int {
    items := make([]int, 0)
    items = append(items, s[:index]...)
    items = append(items, s[index+1:]...)
    return items
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
    fmt.Println(keysToKeep)

    // for each key, assign a value for each number
    // we need to permute all the numbers 0-9
    // ps := Permutations([]int{9,8,7,6,5,4,3,2,1,0}, 9)
    // fmt.Println(ps)


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
