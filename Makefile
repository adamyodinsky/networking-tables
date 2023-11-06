setup-python:
	cd nt-python && poetry install

setup-go:
	cd nt-go && go install

setup-all: setup-python setup-go

run-python:
	cd nt-python && poetry run python networking_tabels/main.py

run-go:
	cd nt-go && go run main.go

test-go:
	cd nt-go && go test
