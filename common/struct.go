package common

type Fetched struct {
	URL     string
	Content string
}

type Error struct {
	URL        string
	StatusCode int
	Error      string
}
