package loop

import (
	"fmt"
	"log"
	"testing"
)

func TestFor(t *testing.T) {
	i := 1
	for i < 5 {
		to := i + 1
		for {
			if true {
				to++
				if to > 10 {
					log.Println("have err")
					return
				}
				continue
			}

			i = to
			break
		}

		if to == 8 {
			break
		}
	}

	fmt.Println("finished")
}
