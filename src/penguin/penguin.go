package penguin

import (
	"fmt"
)

const DEBUG, INFO, WARN, ERROR, FATAL int = 0, 1, 2, 3, 4

// 主函数
func Main() {
	logger(DEBUG, "日志")
	hostname := Config("hostname")
	fmt.Println(hostname)
}
