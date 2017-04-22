package penguin

import "fmt"

func Run() {
	hostname := Config("hostname")
	fmt.Println(hostname)
}
