package concurrrency

type websiteChecker func(string) bool

type result struct {
	string
	bool
}

func fetchurls(urls []string, wc websiteChecker) map[string]bool {
	results := make(map[string]bool)

	resultsChannel := make(chan result, 4)
	for _, url := range urls {
		go func() {
			resultsChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}
	return results
}
