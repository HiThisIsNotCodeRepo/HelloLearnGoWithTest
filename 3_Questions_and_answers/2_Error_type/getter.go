package error_type

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type BadStatusError struct {
	URL    string
	Status int
}

func (b BadStatusError) Error() string {
	return fmt.Sprintf("did not get 200 from %s, got %d", b.URL, b.Status)
}
func DumbGetter(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("problem fetching from %s,%v", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return "", BadStatusError{url, res.StatusCode}
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body), nil
}
