package model

// {
//  "coord": {
//    "lon": 145.77,
//    "lat": -16.92
//  },
//  "weather": [
//    {
//      "id": 803,
//      "main": "Clouds",
//      "description": "broken clouds",
//      "icon": "04n"
//    }
//  ],
//  "base": "cmc stations",
//  "main": {
//    "temp": 293.25,
//    "pressure": 1019,
//    "humidity": 83,
//    "temp_min": 289.82,
//    "temp_max": 295.37
//  },
//  "wind": {
//    "speed": 5.1,
//    "deg": 150
//  },
//  "clouds": {
//    "all": 75
//  },
//  "rain": {
//    "3h": 3
//  },
//  "dt": 1435658272,
//  "sys": {
//    "type": 1,
//    "id": 8166,
//    "message": 0.0166,
//    "country": "AU",
//    "sunrise": 1435610796,
//    "sunset": 1435650870
//  },
//  "id": 2172797,
//  "name": "Cairns",
//  "cod": 200
// }
type WeatherData struct {
  Name string `json:"name"`

  Coord struct {
    Longitude float64 `json:"lon"`
    Latitude  float64 `json:"lat"`
  } `json:"coord"`

  Weather []struct {
    Id          int64    `json:"id"`
    Main        string `json:"main"`
    Description string `json:"description"`
    Icon        string `json:"icon"`
  } `json:"weather"`

  Base string `json:"base"`

  Main struct {
    Kelvin    float64 `json:"temp"`
    HPa       float64 `json:"pressure"`
    Humidity  float64 `json:"humidity"`
    MinKelvin float64 `json:"temp_min"`
    MaxKelvin float64 `json:"temp_max"`
  } `json:"main"`

  Wind struct {
    Speed float64 `json:"speed"`
    Direction float64 `json:"deg"`
  } `json:"wind"`

  Clouds struct {
    All float64 `json:"all"`
  } `json:"clouds"`

  Rain struct {
    ThreeHour float64 `json:"3h"`
  } `json:"rain"`

  Date int64 `json:"dt"`

  Sys struct {
    Type int64 `json:"type"`
    Id int64 `json:"id"`
    Message float64 `json:"message"`
    Country string `json:"country"`
    Sunrise int64 `json:"sunrise"`
    Sunset int64 `json:"sunset"`
  } `json:"sys"`

  Id int64 `json:"id"`
  Cod int64 `json:"cod"`
}
