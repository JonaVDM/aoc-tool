package aoc

import (
	"fmt"
	"io"
	"net/http"
)

func DownloadInput(year, day int, token string) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Cookie":     {fmt.Sprintf("session=%s", token)},
		"User-Agent": {"https://github.com/jonavdm/aoc-tool at jonajvdm@gmail.com"},
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
