package pager

// New creates a pager to iterate through pages.
func New[T any](data []T, pageSize int) chan []T {
	if pageSize < 1 {
		panic("page size cannot be less than zero")
	}
	rv := make(chan []T)
	var index int
	go func() {
		for {
			if index >= len(data) {
				close(rv)
				return
			}
			end := index + pageSize
			if end > len(data) {
				end = len(data)
			}
			rv <- data[index:end]
			index += pageSize
		}
	}()
	return rv
}
