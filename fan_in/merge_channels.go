package fan_in

import "sync"

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	resChan := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				resChan <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	return resChan
}
