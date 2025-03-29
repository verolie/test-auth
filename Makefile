APP_NAME=test-auth
GO_FILES=$(shell find . -name '*.go')

.PHONY: build run clean tidy fmt lint test

## 🏗️ Build aplikasi
build:
	go build -o bin/$(APP_NAME) ./cmd/main.go

## 🚀 Jalankan aplikasi
run:
	go run cmd/main.go

## 🧹 Bersihkan file yang tidak perlu
clean:
	rm -rf bin/$(APP_NAME)

## 🔄 Update dependensi
tidy:
	go mod tidy

## 🏗️ Build & jalankan aplikasi
start: build
	./bin/$(APP_NAME)
