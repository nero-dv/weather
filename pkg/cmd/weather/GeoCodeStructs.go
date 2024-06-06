package weather

type Geocode struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
	Query    Query     `json:"query"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
	Bbox       []float64  `json:"bbox"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	Datasource    Datasource `json:"datasource"`
	Country       string     `json:"country"`
	CountryCode   string     `json:"country_code"`
	State         string     `json:"state"`
	County        string     `json:"county"`
	City          string     `json:"city"`
	Postcode      string     `json:"postcode"`
	Suburb        string     `json:"suburb"`
	Street        string     `json:"street"`
	Housenumber   string     `json:"housenumber"`
	Lon           float64    `json:"lon"`
	Lat           float64    `json:"lat"`
	StateCode     string     `json:"state_code"`
	ResultType    string     `json:"result_type"`
	Formatted     string     `json:"formatted"`
	AddressLine1  string     `json:"address_line1"`
	AddressLine2  string     `json:"address_line2"`
	Category      string     `json:"category"`
	Timezone      Timezone   `json:"timezone"`
	PlusCode      string     `json:"plus_code"`
	PlusCodeShort string     `json:"plus_code_short"`
	Rank          Rank       `json:"rank"`
	PlaceID       string     `json:"place_id"`
}

type Datasource struct {
	Sourcename  string `json:"sourcename"`
	Attribution string `json:"attribution"`
	License     string `json:"license"`
	URL         string `json:"url"`
}

type Rank struct {
	Importance            float64 `json:"importance"`
	Popularity            float64 `json:"popularity"`
	Confidence            float64 `json:"confidence"`
	ConfidenceCityLevel   int64   `json:"confidence_city_level"`
	ConfidenceStreetLevel int64   `json:"confidence_street_level"`
	MatchType             string  `json:"match_type"`
}

type Timezone struct {
	Name             string `json:"name"`
	OffsetSTD        string `json:"offset_STD"`
	OffsetSTDSeconds int64  `json:"offset_STD_seconds"`
	OffsetDST        string `json:"offset_DST"`
	OffsetDSTSeconds int64  `json:"offset_DST_seconds"`
	AbbreviationSTD  string `json:"abbreviation_STD"`
	AbbreviationDST  string `json:"abbreviation_DST"`
}

type Query struct {
	Text   string `json:"text"`
	Parsed Parsed `json:"parsed"`
}

type Parsed struct {
	Housenumber  string `json:"housenumber"`
	Street       string `json:"street"`
	Postcode     string `json:"postcode"`
	District     string `json:"district"`
	Country      string `json:"country"`
	ExpectedType string `json:"expected_type"`
}
