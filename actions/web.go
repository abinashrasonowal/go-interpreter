package actions

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"regexp"
)

// ReadURL fetches the content of a URL and returns a simplified text version.
func ReadURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body: %v", err)
	}

	content := string(bodyBytes)
	
	// Basic HTML stripping
	// 1. Remove scripts and styles
	reScript := regexp.MustCompile(`(?s)<script.*?>.*?</script>`)
	content = reScript.ReplaceAllString(content, "")
	
	reStyle := regexp.MustCompile(`(?s)<style.*?>.*?</style>`)
	content = reStyle.ReplaceAllString(content, "")

	// 2. Remove tags
	reTags := regexp.MustCompile(`<[^>]*>`)
	content = reTags.ReplaceAllString(content, " ")

	// 3. Clean up whitespace
	content = strings.Join(strings.Fields(content), " ")

	return content, nil
}
