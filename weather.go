package main

import (
  "encoding/json"
  "github.com/bimtaye/weather/model"
  "log"
  "net/http"
  "os"
  "strings"
)

//
func main() {
  logger()

  http.HandleFunc("/hello", hello)
  http.HandleFunc("/hello/", helloWithName)
  http.HandleFunc("/weather/", weather)
  http.ListenAndServe(":8080", nil)
}

func logger() {
  // Create your file with desired read/write permissions.
  logfile, err := os.OpenFile("/var/log/weather/weather.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

  if err != nil {
    log.Fatal(err)
  }

  // Defer to close when you're done with it, not because you think it's idiomatic!
  defer logfile.Close()

  // Set output of logs to logfile.
  log.SetOutput(logfile)

  // Initial message.
  log.Println("Creating logfile.")
}

// Prints greeting without a name.
func hello(writer http.ResponseWriter, request *http.Request) {
  writer.Write([]byte("Hello!"))
}

// Prints greeting along with supplied string.
func helloWithName(writer http.ResponseWriter, request *http.Request) {
  name := strings.SplitN(request.URL.Path, "/", 3)[2]
  log.Println("name = " + name)

  writer.Write([]byte("Hello, " + name + "!"))
}

func weather(writer http.ResponseWriter, request *http.Request) {
  city := strings.SplitN(request.URL.Path, "/", 3)[2]
  log.Println("city = " + city)
  data, err := query(city)

  if err != nil {
    http.Error(writer, err.Error(), http.StatusInternalServerError)

    return
  }

  writer.Header().Set("Content-Type", "application/json; charset=utf-8")
  json.NewEncoder(writer).Encode(data)
}

func query(city string) (model.WeatherData, error) {
  response, err := http.Get(
    "http://api.openweathermap.org/data/2.5/weather?APPID=d9bf08ae25948d305fc7ee2db0cde094&q=" + city)

  if err != nil {
    return model.WeatherData{}, err
  }

  defer response.Body.Close()
  var weather model.WeatherData

  if err := json.NewDecoder(response.Body).Decode(&weather); err != nil {
    return model.WeatherData{}, err
  }

  return weather, nil
}
