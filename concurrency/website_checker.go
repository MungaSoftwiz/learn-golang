package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// help control the communication between different processes, allowing us to avoid a race condition bug
	resultChannel := make(chan result)


	// We normally wait for a func to return/process to finish(It blocking)
	// An op that doesn't block will run in a separate process called a "goroutine"
	// Only way to start a goroutine is to put go im front of a function call
	// Below uses an anonymous function

	for _, url := range urls {
		go func(u string) {
	// We send a result struct for each call to wc to resultChannel with a "send statement"
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// We give each anony func a param cause it was taking a reference to the "url" var.
	// u is a copy of the value of url & fixed as value of url for iteration
	//this "()"makes goroutine execute the same time they are declared

	for i := 0; i < len(urls); i++ {
	// We use a receive expression (<-) which we assign a value received from the channel to a var
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

/* go test -race */ //Helps us debug problems with concurrent code
