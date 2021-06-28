package addcheck

import "fmt"

func shouldFailToAddCheck() {
	test := 2 + 1
	fmt.Println(test)
}
