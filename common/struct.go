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

type ExtractedData struct {
	Title       string
	Description string
	Tags        []string
	URLs        []string
}

type ImgInfo struct {
	URL         string
	Description string
}
