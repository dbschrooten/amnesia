test:
	go test -cover ./src/...

plugins:
	go build -buildmode=plugin -o plugins/service-elasticsearch.so src/extension/plugins/service-elasticsearch/elasticsearch.go
