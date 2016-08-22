package main

import (
    "net/http"
    "encoding/json"
    "strings"
    "time"
    "net/url"
    //"fmt"
    "log"
    "io/ioutil"
)

func main() {
  println("[Server Start]")
  http.HandleFunc("/hello",hello)

  http.HandleFunc("/story/", func(writer http.ResponseWriter, r *http.Request) {
      begin := time.Now()
      storyid := strings.SplitN(r.URL.Path, "/", 3)[2]

      fullURL, err := url.Parse("https://www.fanfiction.net/s/" + storyid)
      if err != nil {
      	log.Fatal(err)
      }
      resp, err := http.Get(fullURL.String())
      if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
      }
      defer resp.Body.Close()
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
          http.Error(writer, err.Error(), http.StatusInternalServerError)
          return
      }

      writer.Header().Set("Content-Type", "application/json; charset=utf-8")
      json.NewEncoder(writer).Encode(map[string]interface{}{
          "storyid": storyid,
          "body": body,
          "took": time.Since(begin).String(),
      })
  })

  http.ListenAndServe(":8080", nil)
}


func hello(writer http.ResponseWriter, request *http.Request) {
  writer.Write([]byte("ydara says hello\n"))
}
