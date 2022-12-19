package pager

// Channel creates a channel to iterate through pages.
// Use Func for 20x better performance and fewer allocations at the cost
// of slightly worse syntax.
func Channel[T any](data []T, pageSize int) chan []T {
	if pageSize < 1 {
		panic("page size cannot be <= 0")
	}
	rv := make(chan []T)
	go func() {
		defer close(rv)
		pages, remainder := len(data)/pageSize, len(data)%pageSize
		if remainder > 0 {
			pages++
		}
		for i := 0; i < pages; i++ {
			end := i*pageSize + pageSize
			if end > len(data) {
				end = len(data)
			}
			rv <- data[i*pageSize : end]
		}
	}()
	return rv
}

// Func uses a function parameter to iterate through pages.
// It's 20x faster than the channel variant, using fewer allocations.
func Func[T any](data []T, pageSize int, f func([]T) error) (err error) {
	if pageSize < 1 {
		panic("page size cannot be <= 0")
	}
	pages, remainder := len(data)/pageSize, len(data)%pageSize
	if remainder > 0 {
		pages++
	}
	for i := 0; i < pages; i++ {
		end := i*pageSize + pageSize
		if end > len(data) {
			end = len(data)
		}
		err = f(data[i*pageSize : end])
		if err != nil {
			return err
		}
	}
	return nil
}
