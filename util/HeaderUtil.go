package util

// GetJSONHeader gets HTTP Header with JSON content type.
func GetJSONHeader() (string, string) {
	return "Content-Type", "application/json; charset=utf-8"
}
