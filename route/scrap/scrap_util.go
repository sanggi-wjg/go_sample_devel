package scrap

import (
	"bytes"
	"golang.org/x/net/html/charset"
	"io/ioutil"
)

func ConvertToUTF8(str string, origEncoding string) string {
	strByte := []byte(str)
	byteReader := bytes.NewReader(strByte)
	reader, _ := charset.NewReaderLabel(origEncoding, byteReader)
	strByte, _ = ioutil.ReadAll(reader)
	return string(strByte)
}
