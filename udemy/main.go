// for readability and simplicity move to one file architecture
package main

import (
	"fmt"
    "time"
)

// simple broadcaster broadcast message
func msg(name string) <-chan string {
	c := make(chan string)
	// start go routine which send message for ever
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d\n", name, i)
		}
	}()
	return c
}

func fnl(c1 <-chan string, c2 <-chan string, done chan bool) <-chan string {
	c := make(chan string)
	go func() {
        t := time.After(1 * time.Second)
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
            case <- t:
                close(c)
			    done <- true
                return
			}
		}
	}()
	return c
}

func main() {
	// create chan-func
	ma := msg("Mashka")
	sa := msg("Slavka")
	done := make(chan bool)
	res := fnl(ma, sa, done)
    for i := 0; ; i++ {
        select {
        case <- done:
            // can not break out of select, but can goto label.
            goto label
        default:
		fmt.Printf("%d\t%s", i, <-res)
    }
	}
    // out of for loop
    label:
    fmt.Println("Finished!")
}
