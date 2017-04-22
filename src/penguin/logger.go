package penguin

import (
	"fmt"
	"time"
)

func logger(level int, logstr string) {
	var l string
	switch level {
	case 0:
		l = "DEBUG"
	case 1:
		l = "INFO"
	case 2:
		l = "WARN"
	case 3:
		l = "ERROR"
	case 4:
		l = "FATAL"
	default:
		l = "DEBUG"
	}
	t := time.Now().Format("Mon Jan _2 15:04:05 2006")
	log := t + " " + l + " " + logstr
	fmt.Println(log)
}
