# Geodata-API

Returns geojson for a given country's borders.

The data is sourced from [natural-earth-vector](https://github.com/nvkelso/natural-earth-vector). The specific geojson files can be found here:
- 10m: https://raw.githubusercontent.com/nvkelso/natural-earth-vector/master/geojson/ne_10m_admin_0_countries.geojson
- 50m: https://raw.githubusercontent.com/nvkelso/natural-earth-vector/master/geojson/ne_50m_admin_0_countries.geojson
- 110m: https://raw.githubusercontent.com/nvkelso/natural-earth-vector/master/geojson/ne_110m_admin_0_countries.geojson

I had no part in the creation of this data.

## Usage

- View all countries: https://geojson-api.usingthe.computer/countries
- View geojson for specific country: https://geojson-api.usingthe.computer/countries/brazil
- Specify detail (default 110m): https://geojson-api.usingthe.computer/countries/brazil?detail=10m