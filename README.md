# Amnesia - Health Checks
[![](https://godoc.org/github.com/dbschrooten/amnesia?status.svg)](http://godoc.org/github.com/dbschrooten/amnesia) [![Build Status](https://travis-ci.org/dbschrooten/amnesia.svg?branch=master)](https://travis-ci.org/dbschrooten/amnesia)

Amnesia is a little program that provides health checks and notifications/alerts. It is to be written in such a modular manner that it is easy to add additional services.

This project is still a work in progress. I do not have much spare time so development goes at a slower pace.

### Supported Checks (first draft)

#### Services
- elasticsearch
- rabbitmq
- mysql
- postgresql

#### Kubernetes
- pod
- deployment
- statefulset
- cronjob

#### Other
- tcp
- udp
- http
- https
- graphql

### Notifications (first draft)
- slack
- email
- pagerduty