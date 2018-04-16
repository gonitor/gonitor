package util

import (
	"strings"

	xj "github.com/basgys/goxml2json"
)

// ConvertXMLToJSON .
func ConvertXMLToJSON(xml string) string {
	xmlReader := strings.NewReader(xml)
	json, _ := xj.Convert(xmlReader)
	return json.String()
}
