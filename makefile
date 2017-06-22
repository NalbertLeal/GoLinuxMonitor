all:
	@echo "building ..."
	go build src/main.go
	mv main GoCpu
	mv GoCpu bin