package pager

// New creates a pager to iterate through pages of strings.
func New(data []string, pageSize int) chan []string {
	if pageSize < 1 {
		panic("page size cannot be less than zero")
	}
	rv := make(chan []string)
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
