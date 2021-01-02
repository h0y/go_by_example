package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?&%*^_(_"

	stdEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(stdEnc)

	stdDec, _ := base64.StdEncoding.DecodeString(stdEnc)
	fmt.Println(string(stdDec))
	fmt.Println()

	urlEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlEnc)
	urlDec, _ := base64.URLEncoding.DecodeString(urlEnc)
	fmt.Println(string(urlDec))
}