# .ENV
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

## PIP VENV

VENV_DIR := .venv
PY_CMD := ./$(VENV_DIR)/bin/python3

install_venv: requirements.txt
	python3 -m venv $(VENV_DIR)
	./$(VENV_DIR)/bin/pip install -r requirements.txt

# Script

get_test_data_ETHUSDT_1m:
	$(PY_CMD) script/test_data.py -s "ETHUSDT" -i "1m" -l 80 -e

# Test

test_csv_reader:
	go test -v ./internal/csv/reader_test.go

test_ichimoku:
	go test -v ./pkg/ichimoku/ichimoku_test.go