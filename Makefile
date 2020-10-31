build: build_go build_wc

test: test_go

watch:
	make -j 2 watch_go watch_wc

build_go:
	go build -o ./bin/amnesia ./src/main.go

test_go:
	go test -cover ./src/...

watch_go:
	air -c .air.toml

plugins_go:
	go build -buildmode=plugin -o plugins/service-elasticsearch.so src/extension/plugins/service-elasticsearch/elasticsearch.go

build_wc:
	npm run build

test_wc:
	npm run test

watch_wc:
	npm run watch
