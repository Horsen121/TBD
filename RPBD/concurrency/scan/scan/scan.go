package scan

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

// TODO: реализация в main не возвращает слайс открытых портов
// необходимо реализовать функцию Scan по по аналогии с кодом в main.go
// только её нужно дополнить, чтобы вернуть слайс открытых портов
// отсортированных по возрастанию

func worker(address string, ports chan int, wg *sync.WaitGroup, opened chan<- int) {
	for p := range ports {
		wg.Done()
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, p))
		if err != nil {
			continue
		}
		conn.Close()

		fmt.Println("port opened:", p)
		opened <- p
	}
}

func Scan(address string) []int {
	ports := make(chan int, 200)
	wg := sync.WaitGroup{}
	var open []int
	opened := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(address, ports, &wg, opened)
	}

	for i := 1; i < 10000; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)
	close(opened)

	for el := range opened {
		open = append(open, el)
	}
	sort.Slice(open, func(i, j int) bool {
		return open[i] < open[j]
	})

	return open
}
