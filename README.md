# 🦠📈 covid-19-nl-influxdb

Data importer that writes Dutch historic COVID-19 case data as metrics to [InfluxDB](https://www.influxdata.com/products/influxdb-overview/), an open source times series database. The data can
be used for manual exploration, via the Data Explorer and Dashboards feature in
InfluxDB's UI, or as a data source for Grafana.

The project consists of a `docker-compose.yml` file and a `datawriter` program
written in [Go](https://golang.org/) that fetches the data and writes it as
metrics to a locally running InfluxDB service.

![Imgur](https://imgur.com/Dkq9Snr.png)

## Usage

Clone the repository and run the following from the repo directory:
Set the following environment variables. Alter these if you expose the service
to the outside world.

```
export INFLUXDB_TOKEN=covid-19-nl-token
export INFLUXDB_PASSWORD=covid-19-nl-password
```

Run the services via Docker Compose:

```sh
docker-compose up
```

Login at http://localhost:9999 (username: `covid-19-nl`, password: `covid-19-nl-password`).

The metrics currently have the following characteristics, but this is subject to change:

- `confirmed`: Number of confirmed COVID-19 cases (Field)
- `citizens` Number of citizens in related municipality (Field)
- `municipality` Name of the municipality (Tag)
- `municipality_num` Municipality number (_TODO: explain_) (Tag)

Data is fetched and stored at start up, and every 15 minutes.

## Prerequisites

- [Docker](https://www.docker.com/get-started) (17.09.0+)

## TODO

- [ ] Assert that the metrics correspond with the scraped data.
- [ ] Host publicly available InfluxDB UI.
- [ ] Write dashboards
- [ ] Fetch/scrape more data (IC, deaths)

## Acknowledgements

Thanks to the [COVID-19 Scraper for NL](https://github.com/Kapulara/COVID-19-NL).

## Other resources

- [Code for NL](https://www.codefor.nl/corona/)
- [#corona-data](https://praatmee.codefor.nl/) (Slack)

## License

[MIT](LICENSE)
