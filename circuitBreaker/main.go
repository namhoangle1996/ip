package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP TEST"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	st.IsSuccessful = func(err error) bool {
		fmt.Println("st err", err)
		return err != nil
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {

		fmt.Println("cb.Counts().Requests prev", cb.Counts().Requests)
		fmt.Println("cb.Counts().State prev", cb.State().String())

		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		fmt.Println("getErr,", err)
		return nil, err
	}

	fmt.Println("cb.Counts().TotalFailures", cb.Counts().TotalFailures)
	fmt.Println("cb.Counts().Requests", cb.Counts().Requests)

	return body.([]byte), nil
}

func main() {

	for i := 1; i < 10; i++ {
		body, err := Get("https://www.google1.com/robots.txt1")
		body, err = Get("https://www.google.com/robots.txt1")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}

}

func findLongestSubStr(input string) (res string) {
	for i, v := range input {
		tmpRes := fmt.Sprintf("%c", v)
		fmt.Println("get tmpRes", tmpRes)
		tmpHmap := make(map[string]bool, len(input)-i)
		tmpHmap[tmpRes] = true

		for i1 := i + 1; i1 < len(input); i1++ {
			x1 := fmt.Sprintf("%c", input[i1])
			if _, ok := tmpHmap[x1]; ok {
				break
			} else {
				tmpHmap[x1] = true
				tmpRes = tmpRes + x1
				if len(tmpRes) > len(res) {
					res = tmpRes
				}
			}
		}

	}

	return res
}
