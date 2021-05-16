# crawler

## 소개
Pinetrest 수준의 이미지 검색 시스템을 위한 이미지 데이터 수집 crawler이며 아래 2가지 방식으로 실행 가능하도록 구현이 될 예정이다.

    1. 각 컴포넌트가 함께 하나의 서버로 구동
    2. 각 컴포넌트가 각각 서버로 구동되어 확장 가능한 형태로 서비스

## components
### common
공통으로 사용되는 Data Type 또는 Util이 정의가 된다.

### frontier
URL Set에서 URL을 가져와 Fetcher에게 넘겨준다.    
요청 URL 및 도메인의 이전 응답 여부 및 크롤링 주기 체크가 여기에서 이루어진다.    

### fetcher
요청 URL에 대한 요청을 가져오고 응답에 따른 처리를 진행한다.    
응답 성공인 경우에는 html 페이지를 반환하고, 에러인 경우에는 에러 코드에 따라 별도 처리 후 에러코드를 반환한다.    

### parser
fetcher가 넘겨준 URL의 페이지에서 필요한 데이터와 새롭게 탐색할 URL을 parsing 한다.

### content checker
순환참조가 되어 있는 경우 그리고 동일한 컨텐츠가 복사 붙여넣기가 되어 있는 경우에 대비하기 위해 parsing된 데이터의 Finger Print를 사용해 이전에 탐색된 페이지인지 확인을 한다.

### url filter
해당 도메인의 robots.txt 파일을 참조해서 어떤 URL들을 URL set에 넣을 것인지 filter 한다.

## data
### url set
중복을 제거한 url들을 가지고 있는다.

### Document's Finger Print
문서들의 Finger Print를 가지고 있는다. 

### error log
잘못된 응답을 받은 페이지, 과도한 요청으로 금지당한 목록을 저장한다.