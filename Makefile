test:
	go test -cover ./src/...

plugins:
	go build -buildmode=plugin -o plugins/elasticsearch.so src/plugin/plugins/elasticsearch/elasticsearch.go
