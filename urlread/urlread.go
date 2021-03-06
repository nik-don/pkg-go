/* Package urlread

github.com/nik-don/pkg-go/urlread
*/
package urlread

import (
	"bufio"
	"io"
	"net/http"
)

// LinesFromReader accepts argument of io.Reader, which is an interface
// and return a slice of string of the lines.
func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// GetFromURL returns a slice of string of in the form of lines from a URL
func GetFromURL(URL string) ([]string, error) {

	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return LinesFromReader(response.Body)

}
