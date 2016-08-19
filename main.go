package main

import (
    "net/http"
    "encoding/json"
    "strings"
)

func main() {
  println("[Server Start]")
  http.HandleFunc("/hello",hello)

  http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
        city := strings.SplitN(r.URL.Path, "/", 3)[2]

        data, err := query(city)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        json.NewEncoder(w).Encode(data)
    })

  http.ListenAndServe(":8080", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
  writer.Write([]byte("ydara says hello\n"))
}

type weatherData struct {
  Name string `json:"name"`
    Main struct {
        Kelvin float64 `json:"temp"`
    } `json:"main"`
}

func query (city string) (weatherData, error) {
  resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?&appid=7b1e1ba73cc6d7063c88208e5fc50adc&q=" + city)
  if err != nil {
    return weatherData{}, err
  }
  defer resp.Body.Close()
  var d weatherData
  if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
    return weatherData{}, err
  }
  return d, nil
}
