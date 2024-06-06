package weather

type WeatherData struct {
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Timezone  string        `json:"timezone"`
	Offset    float64       `json:"offset"`
	Elevation int64         `json:"elevation"`
	Currently Currently     `json:"currently"`
	Minutely  Minutely      `json:"minutely"`
	Hourly    Hourly        `json:"hourly"`
	Daily     Daily         `json:"daily"`
	Alerts    []interface{} `json:"alerts"`
	Flags     Flags         `json:"flags"`
}

type Currently struct {
	Time                 int64      `json:"time"`
	Summary              Summary    `json:"summary"`
	Icon                 Icon       `json:"icon"`
	NearestStormDistance *float64   `json:"nearestStormDistance,omitempty"`
	NearestStormBearing  *float64   `json:"nearestStormBearing,omitempty"`
	PrecipIntensity      float64    `json:"precipIntensity"`
	PrecipProbability    float64    `json:"precipProbability"`
	PrecipIntensityError float64    `json:"precipIntensityError"`
	PrecipType           PrecipType `json:"precipType"`
	Temperature          float64    `json:"temperature"`
	ApparentTemperature  float64    `json:"apparentTemperature"`
	DewPoint             float64    `json:"dewPoint"`
	Humidity             float64    `json:"humidity"`
	Pressure             float64    `json:"pressure"`
	WindSpeed            float64    `json:"windSpeed"`
	WindGust             float64    `json:"windGust"`
	WindBearing          float64    `json:"windBearing"`
	CloudCover           float64    `json:"cloudCover"`
	UvIndex              float64    `json:"uvIndex"`
	Visibility           float64    `json:"visibility"`
	Ozone                float64    `json:"ozone"`
	PrecipAccumulation   *float64   `json:"precipAccumulation,omitempty"`
}

type Daily struct {
	Summary Summary      `json:"summary"`
	Icon    Icon         `json:"icon"`
	Data    []DailyDatum `json:"data"`
}

type DailyDatum struct {
	Time                        int64      `json:"time"`
	Icon                        Icon       `json:"icon"`
	Summary                     Summary    `json:"summary"`
	SunriseTime                 int64      `json:"sunriseTime"`
	SunsetTime                  int64      `json:"sunsetTime"`
	MoonPhase                   float64    `json:"moonPhase"`
	PrecipIntensity             float64    `json:"precipIntensity"`
	PrecipIntensityMax          float64    `json:"precipIntensityMax"`
	PrecipIntensityMaxTime      int64      `json:"precipIntensityMaxTime"`
	PrecipProbability           float64    `json:"precipProbability"`
	PrecipAccumulation          float64    `json:"precipAccumulation"`
	PrecipType                  PrecipType `json:"precipType"`
	TemperatureHigh             float64    `json:"temperatureHigh"`
	TemperatureHighTime         int64      `json:"temperatureHighTime"`
	TemperatureLow              float64    `json:"temperatureLow"`
	TemperatureLowTime          int64      `json:"temperatureLowTime"`
	ApparentTemperatureHigh     float64    `json:"apparentTemperatureHigh"`
	ApparentTemperatureHighTime int64      `json:"apparentTemperatureHighTime"`
	ApparentTemperatureLow      float64    `json:"apparentTemperatureLow"`
	ApparentTemperatureLowTime  int64      `json:"apparentTemperatureLowTime"`
	DewPoint                    float64    `json:"dewPoint"`
	Humidity                    float64    `json:"humidity"`
	Pressure                    float64    `json:"pressure"`
	WindSpeed                   float64    `json:"windSpeed"`
	WindGust                    float64    `json:"windGust"`
	WindGustTime                int64      `json:"windGustTime"`
	WindBearing                 float64    `json:"windBearing"`
	CloudCover                  float64    `json:"cloudCover"`
	UvIndex                     float64    `json:"uvIndex"`
	UvIndexTime                 int64      `json:"uvIndexTime"`
	Visibility                  float64    `json:"visibility"`
	TemperatureMin              float64    `json:"temperatureMin"`
	TemperatureMinTime          int64      `json:"temperatureMinTime"`
	TemperatureMax              float64    `json:"temperatureMax"`
	TemperatureMaxTime          int64      `json:"temperatureMaxTime"`
	ApparentTemperatureMin      float64    `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime  int64      `json:"apparentTemperatureMinTime"`
	ApparentTemperatureMax      float64    `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime  int64      `json:"apparentTemperatureMaxTime"`
}

type Flags struct {
	Sources        []string    `json:"sources"`
	SourceTimes    SourceTimes `json:"sourceTimes"`
	NearestStation int64       `json:"nearest-station"`
	Units          string      `json:"units"`
	Version        string      `json:"version"`
}

type SourceTimes struct {
	HrrrSubh string `json:"hrrr_subh"`
	Hrrr018  string `json:"hrrr_0-18"`
	Nbm      string `json:"nbm"`
	NbmFire  string `json:"nbm_fire"`
	Hrrr1848 string `json:"hrrr_18-48"`
	Gfs      string `json:"gfs"`
	Gefs     string `json:"gefs"`
}

type Hourly struct {
	Summary Summary     `json:"summary"`
	Icon    Icon        `json:"icon"`
	Data    []Currently `json:"data"`
}

type Minutely struct {
	Summary Summary         `json:"summary"`
	Icon    string          `json:"icon"`
	Data    []MinutelyDatum `json:"data"`
}

type MinutelyDatum struct {
	Time                 int64      `json:"time"`
	PrecipIntensity      float64    `json:"precipIntensity"`
	PrecipProbability    float64    `json:"precipProbability"`
	PrecipIntensityError float64    `json:"precipIntensityError"`
	PrecipType           PrecipType `json:"precipType"`
}

type Icon string

const (
	ClearDay          Icon = "clear-day"
	ClearNight        Icon = "clear-night"
	Cloudy            Icon = "cloudy"
	PartlyCloudyDay   Icon = "partly-cloudy-day"
	PartlyCloudyNight Icon = "partly-cloudy-night"
)

type PrecipType string

const (
	None PrecipType = "none"
	Rain PrecipType = "rain"
)

type Summary string

const (
	Clear         Summary = "Clear"
	PartlyCloudy  Summary = "Partly Cloudy"
	SummaryCloudy Summary = "Cloudy"
)
