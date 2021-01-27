package p022

import (
    "fmt"
    "log"
    "io/ioutil"
    "sort"
    "strings"
)

func main() {
    filename:="p022_names.txt"

    // open file
    // sort
    // score each name
    // report sum

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    names := strings.Split(string(data), ",")
    sort.Strings(names)

    sum:=0
    for i, name := range names {
        s:=0
        for j:=1;j<len(name)-1;j++ {
            n:=int(name[j] - 'A') + 1
            s += n
        }
        sum += s * (i + 1)
    }

    fmt.Printf("%v\n", sum)
}
