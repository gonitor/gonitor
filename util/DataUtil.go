package util

import (
	"strings"

	xj "github.com/basgys/goxml2json"
)

// ConvertXMLToJSON converts XML string into JSON string.
func ConvertXMLToJSON(xml string) string {
	xmlReader := strings.NewReader(xml)
	json, _ := xj.Convert(xmlReader)
	return json.String()
}
