# Amnesia - Health Checks
[![](https://godoc.org/github.com/dbschrooten/amnesia?status.svg)](http://godoc.org/github.com/dbschrooten/amnesia) [![Build Status](https://travis-ci.org/dbschrooten/amnesia.svg?branch=master)](https://travis-ci.org/dbschrooten/amnesia)

Amnesia is a modular tool to perform health checks. It has support for a variety of protocols and services. Users can define events per check with differing requirements, and setup alerts for E-mail, Slack and Pagerduty. Amnesia also contains a REST api with a concise overview of all checks, and an endpoint to query a embedded database. Exporters can be defined to export event data to Prometheus, InfluxDB.

In a later stage there will also a web-ui that can be activated, built with webcomponents. That can be used to perform basic queries and visualize the data through graphs. The program can be easily extended in functionality through the use of plugins.

The reason I created amnesia is because I couldn't find a proper all-in-one tool for health checking, only overpriced third party services, that charge way too much for something this simple.

Amnesia is currently still a ***work in progress***, binaries will be released when functionality of first draft has been completed and has 100% test coverage. Following that I will first proceed with documentation. First draft functionality should already provide most of the functionality to get started.


### First Draft

#### Supported Checks

##### Core
- tcp
- udp
- http
- https
- graphql
- telnet

##### Official Plugins
- rancher
- kubernetes

##### Exporters
- prometheus
- influxdb

##### Alert
- email
- pagerduty
- slack

### Later Stage

##### Official Plugins
- rabbitmq
- nats
- elasticsearch
- couchdb
- mongodb
- mysql
- postgres
  
##### Exporters
- nats
- rabbitmq

##### Other
- dashboard
