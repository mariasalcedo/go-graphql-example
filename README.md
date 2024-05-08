## Go GraphQL example

### Description
The following repository contains a GraphQL API implementation example based on wind farms (i.e. a set of turbines) as entities.
Included is the Rest client use case of getting insights on the weather and geography around the wind farms.

### Description of the implementation

- GraphQL is implemented with `99designs/gqlgen` golang library
- The API communicates with [Open-Meteo](https://open-meteo.com/) for any weather or geography information.
- The API allows CRUD operations, including pagination
- Postgres is used to persist "Wind farm" entities
- Wind farms store the following properties:
    - Name - string
    - Latitude - float
    - Longitude - float
- When updating a wind farm info, an user is required to provide the `id` of the wind farm, but can optionally provide any of those three properties
- When fetching wind farm info, there are a few expectations:
    - Wind farm properties, `weatherForecasts` and `elevation` should be retrieved from [Weather Forecast API](https://open-meteo.com/en/docs) and [Elevation API](https://open-meteo.com/en/docs/elevation-api).
    - For the `weatherForecasts` property, we need the following Hourly Weather Variables from the [Weather API](https://open-meteo.com/en/docs)  below:
        - Temperature (2 m) in Celsius
        - Precipitation (rain + showers + snow) in Millimeter
        - Wind Speed (10 m) in Km/h
        - Wind Direction (10 m)
    - Wind farm property `hasPrecipitationToday` should be determined based on the provided `weatherForecasts`
- Everything is Dockerized for continuous deployment purposes
- Unit tests are provided