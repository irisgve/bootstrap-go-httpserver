PORT := 9000

run: build
	docker run -p $(PORT):$(PORT) bootstrap-go-httpserver:latest

build: config.yaml
	docker build -t bootstrap-go-httpserver .

config.yaml: config.yaml.in
	rm -rf config.yaml
	sed -e 's#%PORT%#$(PORT)#' < $< > $@