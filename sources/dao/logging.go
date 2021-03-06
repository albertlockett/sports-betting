package dao

import (
  "bytes"
  "io/ioutil"
  "log"
  "net/http"
  "strings"
)

type LoggingTransport struct{}

// RoundTrip executes a request, returning a response, and prints information about the flow.
//
func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
  // Print information about the request
  log.Println(strings.Repeat("=", 80))
  log.Printf("> %s %s\n", req.Method, req.URL.String())
  log.Println(strings.Repeat("-", 80))

  res, err := http.DefaultTransport.RoundTrip(req)
  if err == nil {
    // Print information about the response
    log.Printf("< [%s] %s", res.Status, res.Header.Get("Content-Type"))
    body, err := ioutil.ReadAll(res.Body)
    if err == nil {
      // Print the body
      defer func() { res.Body = ioutil.NopCloser(bytes.NewReader(body)) }()
      defer func() { res.Body.Close() }()
      for _, line := range strings.Split(string(body), "\n") {
        if line != "" {
          log.Printf("< %s\n", line)
        }
      }
    }
    log.Println(strings.Repeat("=", 80))
  } else {
    // Print the error
    log.Println("ERROR:", err)
  }

  return res, err
}
