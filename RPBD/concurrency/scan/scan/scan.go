package scan

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

// TODO: реализация в main не возвращает слайс открытых портов
// необходимо реализовать функцию Scan по аналогии с кодом в main.go
// только её нужно дополнить, чтобы вернуть слайс открытых портов
// отсортированных по возрастанию

func worker(address string, ports chan int, wg *sync.WaitGroup, opened *[]int) {
	mu := sync.Mutex{}
	for p := range ports {
		wg.Done()
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, p))
		if err != nil {
			continue
		}
		conn.Close()

		mu.Lock()
		*opened = append(*opened, p)
		mu.Unlock()
	}
}

func Scan(address string) []int {
	ports := make(chan int, 200)
	wg := sync.WaitGroup{}
	var open []int

	for i := 0; i < cap(ports); i++ {
		go worker(address, ports, &wg, &open)
	}

	for i := 1; i < 10000; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)

	sort.Ints(open)

	return open
}
