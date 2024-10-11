package loc

// Counter is a counter that will call f when it reaches maxTime.
func Counter(maxTime int, f func()) (add func()) {
	var count int
	add = func() {
		count++
		if count >= maxTime {
			f()
			count = 0
		}
	}
	return
}
