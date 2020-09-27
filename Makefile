csv-split: cmd/split/main.go pkg/split/split.go
	go build -o split cmd/split/main.go


test: csv-split
	go test -v ./...


.PHONY: clean
clean:
	rm *.csv csv-split
