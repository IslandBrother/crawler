package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	. "github.com/island-brother/crawler/common"
)

func GetContent(html string) {
	fmt.Println("Parser Start...")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		log.Fatal(err)
	}

	//title := getTitleInDocument(doc)
	//description := getDescriptionInDocument(doc)
	imgInfos := getImgDatasInDocument(doc)
	links := getURLsInDocument(doc, pageURL)
	writeLinkInFile(links)
	writeImgInFile(imgInfos)
	//fmt.Println(title, description, links)

}

func getTitleInDocument(doc *goquery.Document) string {
	docTitle := doc.Find("title")
	return docTitle.Text()
}

func getDescriptionInDocument(doc *goquery.Document) string {
	docMetas := doc.Find("meta")
	description := ""

	docMetas.Each(func(i int, meta *goquery.Selection) {
		if name, _ := meta.Attr("name"); name == "description" {
			description, _ = meta.Attr("content")
		}
	})

	return description
}

// img 태그와 picture 태그를 각자 다룰지 아니면 한 번에 다룰지 각 태그의 속성을 확인을 해봐야겠다.
func getImgDatasInDocument(doc *goquery.Document) []ImgInfo {
	var imgInfos []ImgInfo
	docImgs := doc.Find("img")

	docImgs.Each(func(i int, img *goquery.Selection) {
		urlOfImg, _ := img.Attr("src")
		altOfImg, _ := img.Attr("alt")
		imgInfos = append(imgInfos, ImgInfo{urlOfImg, altOfImg})
	})

	return imgInfos
}

func getPicDatasInDocument(doc *goquery.Document) []ImgInfo {
	var picInfos []ImgInfo
	docPicture := doc.Find("picture")

	docPicture.Each(func(i int, pic *goquery.Selection) {
		urlOfImg, _ := pic.Attr("src")
		altOfImg, _ := pic.Attr("alt")
		picInfos = append(picInfos, ImgInfo{urlOfImg, altOfImg})
	})

	return picInfos
}

// rel = "nofollow" or rel="ugc" 속성 확인 필요
func getURLsInDocument(doc *goquery.Document, baseURL string) []string {
	var URLs []string
	docATags := doc.Find("a")

	// URL 중복 제거 기능 필요하겠는데
	docATags.Each(func(i int, tag *goquery.Selection) {
		tempLink, _ := tag.Attr("href")

		if isValidURL(tempLink) {
			link := GetCleansingURL(baseURL, tempLink)
			fmt.Println(link)
			URLs = append(URLs, link)
		}
	})

	return URLs
}

func isValidURL(url string) bool {
	//조건은 검색해보면서 더 추가해야겠네.
	urlSplitWords := strings.Split(url, "/")
	urlSplitWordCnt := len(urlSplitWords)

	return urlSplitWordCnt > 1
}

/*
*	작성 필요
*	Case 1 : URL 형식이 아닌 문자열 삭제 (예 : void(0))
*	Case 2 : 기본 URL이 없는 문자열 정제 (예 : /img/example.png) => (base_url)http://demo.samplePage + (extract_url)/img/example.png
*	Case 3 : 정상적인 URL 문자열 (예 : http://demo.samplePage/img/example.png)
 */
func GetCleansingURL(baseURL string, url string) string {
	cleansingURL := url

	// 먼저 URL이 유효한지부터 확인을 해야겠구나.
	if !strings.Contains(url, "http") {
		cleansingURL = baseURL + url
	}

	return cleansingURL
}

func writeLinkInFile(links []string) {
	file, err := os.Create("links.csv")

	if err != nil {
		log.Fatalln(err)
	}
	//에러 체크 함수 필요

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, link := range links {
		w.Write([]string{link})
		//에러 체크 함수 필요
	}

}

func writeImgInFile(imgInfos []ImgInfo) {
	file, err := os.Create("imgInfos.csv")

	if err != nil {
		log.Fatalln(err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, imgInfo := range imgInfos {
		w.Write([]string{imgInfo.url, imgInfo.alt})
	}
}

/*
*	함수 작성 필요
*	논의 사항 - 데이터 입력 때 DB커넥션이 가장 적은 방안 논의 / insert, batch, loader
 */
func insertDataInDB() {}

/*
func getJsonInDocument(doc *goquery.Documnet) string{
	doc.Find("script").Each(func(i int, script *goquery.Selection) {
		scriptType, _ := script.Attr("type")
		if scriptType == "application/ld+json" {
			fmt.Println(script.Text())
		}
	})
}
*/
