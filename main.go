package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

// 
func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/weather/", weather)
    http.ListenAndServe(":8080", nil)
}

// 
func hello(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("Hello!"))
}

func weather(writer http.ResponseWriter, request *http.Request) {
    city := strings.SplitN(request.URL.Path, "/", 3)[2]
    data, err := query(city)

    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }

    writer.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(writer).Encode(data)
}

// 
type weatherData struct {
    Name string `json:"name"`

    Main struct {
	Kelvin float64 `json:"temp"`
    } `json:"main"`
}

//
func query(city string) (weatherData, error) {
    resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=d9bf08ae25948d305fc7ee2db0cde094&q=" + city)

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

