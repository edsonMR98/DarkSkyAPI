# DarkSkyAPI
This repo uses the [Dark Sky API](https://darksky.net/dev) to get the weather forecast for a specific location.

The objective of this repo is to consume the data that Dark Sky API offers us. This was done in different programming languages in order to demostrate the variety of available libraries that each language uses to consume this API.


## [Python](https://darksky.net/dev/docs/libraries#python-library)
The [darksky_weather](https://pypi.org/project/darksky_weather/) package was used to consume the API. Also, the [geopy](https://pypi.org/project/geopy/) package is used to get a location from a address (lat - lon) and then get the weather forecast for this address using the Dark Sky API. This API offers us a lot of information related to the weather of a location, including a daily forecast of one week from the current date, wich is displayed.


## [Go](https://darksky.net/dev/docs/libraries#go-library)
The [darksky](https://github.com/shawntoffel/darksky) package allows us to create a Dark Sky API client in Go and consume it. Also, the [reflect](https://golang.org/pkg/reflect/) package is used to get information (name, type and value) about all properties of a struct.
