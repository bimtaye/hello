package main

import "net/http"

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("Hello!"))
}

type weatherData struct {
    Name string `json:"name"`
    Main struct {
        Kelvin float64 `json:"temp"`
    } `json:"main"`
}

type weatherData struct {
    Name string `json:"name"`

    Main struct {
        Kelvin float64 `json:"temp"`
    } `json:"main"`
}

func query(city string) (weatherData, error) {
    resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=YOUR_API_KEY&q=" + city)

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
