package utils

import (
	"net/url"
	"strings"
)

var (
	encodingMap = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	decodingMap = makeDecodingMap(encodingMap)
)

func IsValidURL(urlString string) bool {
	u, err := url.Parse(urlString)

	return err == nil && u.Scheme != "" && u.Host != ""
}

func makeDecodingMap(encodingMap string) map[byte]int {
	decodingMap := make(map[byte]int)
	for i := 0; i < len(encodingMap); i++ {
		decodingMap[encodingMap[i]] = i
	}
	
	return decodingMap
}

func EncodeID(input string) string {
	var result strings.Builder
	input = input[10:15]
	for i := 0; i < len(input); i++ {
		index := int(input[i]) % len(encodingMap)
		result.WriteByte(encodingMap[index])
	}

	return result.String()
}

func DecodeID(input string) string {
	var result strings.Builder
	input = input[10:15]
	for i := 0; i < len(input); i++ {
		index := decodingMap[input[i]]
		result.WriteByte(byte(index))
	}
	return result.String()
}