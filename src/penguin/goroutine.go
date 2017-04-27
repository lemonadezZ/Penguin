package penguin

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

func Producer(queue chan<- int) {
	fmt.Println("我是生产者")
	for i := 0; i < 10; i++ {
		fmt.Println("生产" + strconv.Itoa(i))
		fmt.Println("加入到队列")
		queue <- i
		waitgroup.Done()

	}

}

func Consumer(queue <-chan int, name string) {

	for i := 0; i < 20; i++ {
		waitgroup.Add(1)
		v := <-queue
		fmt.Println(name, "receive:", v)

	}

}

func main() {

	queue := make(chan int, 1)
	go Producer(queue)
	go Consumer(queue, "消费者1")
	go Consumer(queue, "消费者2")
	fmt.Println(time.Now().Format("2006/01/02"))
	waitgroup.Wait()

}
