It provides a rest api to consume [world happiness](https://worldhappiness.report/) data

## Getting started

1. Discover the [api](https://app.swaggerhub.com/apis-docs/beyond-gdp/wold-happiness_api/v1).
2. [Signup](https://world-happiness.auth.eu-west-1.amazoncognito.com/login) to get prod tokens. 
3. Use the api via `https://api.worldhappiness.beyondgdp.co.uk/v1`.

## Examples

**get countries**
```
GET https://api.worldhappiness.beyondgdp.co.uk/v1/countries
Authorization: {{token}}
```
**get countries by name**
```
GET https://api.worldhappiness.beyondgdp.co.uk/v1/countries/spain
Authorization: {{token}}
```

### Support or Contact

Say hello at [eneko@beyondgdp.co.uk](mailto:eneko@beyondgdp.co.uk)
