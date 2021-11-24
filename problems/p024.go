package p024
/*
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

import "fmt"


func permute(prefix string, choices string, out []string) []string {
    for i, c := range choices {
        // Append to prefix
        new_prefix := prefix + string(c)

        // Remove c from the choices
        remaining_choices := ""
        for j, other := range choices {
            if j != i {
                remaining_choices += string(other)
            }
        }

        if len(remaining_choices) == 0 {
            // base case
            out = append(out, new_prefix)
        } else {
            // recursive case
            out = permute(new_prefix, remaining_choices, out)
        }
    }
    return out
}

func main() {
    // Create a list of characters.
    chars := "0123456789"

    // Permute them.
    permutations := make([]string, 0)
    permutations = permute("", chars, permutations)

    // Print the millionth item.
    for i, item := range permutations {
        if i == 999999 {
            fmt.Printf("%s\n", item)
        }
    }
}
