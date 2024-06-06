**Weather CLI Tool**

* * * 

A simple command-line tool to retrieve the current weather for a given address.

<br>

**Usage**

* * *

    $ go run ./cmd/weather/weather.go <address or postal code>

Replace `<address>` with the address or postal code for which you want to retrieve the weather.

<br>

**Environment Variables**

* * *

- `WEATHER_API_KEY`: Your Pirate Weather API key.
- `GEO_API_KEY`: Your Geoapify API key.

<br>

**Functionality**

* * *

This tool uses the Pirate Weather API to retrieve the current temperature for a given city. It first attempts to geocode the city name to retrieve its coordinates, and then uses these coordinates to retrieve the current weather.

<br>

**Note**

* * *

This code is for demonstration purposes only and should not be used in production without proper error handling and testing.

<br>

**License**

* * *

This code is licensed under the MIT License.