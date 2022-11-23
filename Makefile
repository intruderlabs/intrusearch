# all these commands below are for local environment
.PHONY: clean
clean:
	rm -rf build coverage.txt

.PHONY: updates
updates:
	go list -u -m all > updates.txt

.PHONY: build
build:
	go get github.com/google/wire/internal/wire
	go get github.com/google/wire/cmd/wire
	cd main; go run github.com/google/wire/cmd/wire
	go mod tidy

.PHONY: start
start:
	make build

.PHONY: test
test:
	go test -v -tags test,unit,integration \
		-coverpkg ./main/... \
		-covermode=count ./main/... \
		-coverprofile=coverage.txt
	go tool cover -func coverage.txt
