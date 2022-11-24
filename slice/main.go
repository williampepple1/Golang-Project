package main

import "fmt"

func main() {
    months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
    quarter1 := months[0:4]
    quarter2 := months[4: 8]
    fmt.Println(months)
    fmt.Println(quarter1)
    fmt.Println(quarter2)
    fmt.Println()
    fmt.Println("Length:", len(months))
    fmt.Println("Capacity:", cap(months))
}


/*

package main

import "fmt"

func main() {
    letters := []string{"A", "B", "C", "D", "E"}
    remove := 2

	if remove < len(letters) {

		fmt.Println("Before", letters, "Remove ", letters[remove])

		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After", letters)
	}

}

*/
