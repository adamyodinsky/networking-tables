setup-python:
	cd nt_python && poetry install

setup-go:
	cd nt_go && go install

setup-all: setup-python setup-go

run-python:
	cd nt_python && poetry run python networking_tabels/main.py

run-go:
	cd nt_go && go run main.go

test-go:
	cd nt_go && go test

test-python:
	cd nt_python && poetry run python tests/test_e2e.py
