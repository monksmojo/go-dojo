package design03

import (
	"fmt"
	"strings"
)

func BuilderDesign() {
	var url string = getNewUrlBuilder().setProtocol("https://").setHostName("localhost").setPort(8080).setPathParam("/").build()
	fmt.Println(url)
}

type urlBuilder struct {
	protocol   string
	hostname   string
	port       int
	pathParam  string
	queryParam string
}

func getNewUrlBuilder() *urlBuilder {
	return &urlBuilder{}
}

func (url *urlBuilder) setProtocol(protocol string) *urlBuilder {
	url.protocol = protocol
	return url
}

func (url *urlBuilder) setHostName(hostName string) *urlBuilder {
	url.hostname = hostName
	return url
}

func (url *urlBuilder) setPort(port int) *urlBuilder {
	url.port = port
	return url
}

func (url *urlBuilder) setPathParam(pathParam string) *urlBuilder {
	url.pathParam = pathParam
	return url
}

func (url *urlBuilder) setQueryParam(queryParam string) *urlBuilder {
	url.queryParam = queryParam
	return url
}

func (url *urlBuilder) build() string {
	var stringBuilder strings.Builder
	if url.protocol != "" {
		stringBuilder.WriteString(url.protocol)
	}
	if url.hostname != "" {
		stringBuilder.WriteString(url.hostname)
	}
	if url.pathParam != "" {
		stringBuilder.WriteString(url.pathParam)
	}
	if url.queryParam != "" {
		stringBuilder.WriteString(url.hostname)
	}
	return stringBuilder.String()
}
