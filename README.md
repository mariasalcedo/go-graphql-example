## Go GraphQL example

### Description
The following repository contains a GraphQL API implementation example based on wind farms (i.e. a set of turbines) as entities.
Included is the Rest client use case of getting insights on the weather and geography around the wind farms.

## GraphQL Advantages over REST APIs

1. 𝘿𝙖𝙩𝙖 𝙁𝙚𝙩𝙘𝙝𝙞𝙣𝙜: GraphQL allows clients to request only the data they need, reducing over-fetching and under-fetching.

2. 𝙀𝙣𝙙𝙥𝙤𝙞𝙣𝙩 𝙎𝙩𝙧𝙪𝙘𝙩𝙪𝙧𝙚: GraphQL has a single endpoint (/graphql) for all queries, simplifying the API surface.

3. 𝙍𝙚𝙨𝙥𝙤𝙣𝙨𝙚 𝙁𝙤𝙧𝙢𝙖𝙩: Clients define the structure of the response they want, enabling more precise control over data.

4. 𝙑𝙚𝙧𝙨𝙞𝙤𝙣𝙞𝙣𝙜: GraphQL eliminates the need for versioning by allowing clients to evolve queries independently.

5. 𝙏𝙤𝙤𝙡𝙞𝙣𝙜 𝘼𝙣𝙙 𝙀𝙘𝙤𝙨𝙮𝙨𝙩𝙚𝙢: GraphQL offers robust tooling and libraries for development and testing.

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