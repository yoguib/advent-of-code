package main

import "fmt"

func main() {
	juju := []string{"bla", "foo"}
	fmt.Println(len(juju))

	for i := range juju {
		fmt.Println("index", i)
		if i == len(juju)-1 {
			fmt.Println("we go to the end", i)
		}
	}
}
