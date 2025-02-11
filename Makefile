# プロジェクト名（ビルド後のバイナリ名）
BINARY_NAME := get-business-calendar

# proto ファイルと生成先のディレクトリ
PROTO_DIR := internal/proto
PROTO_FILE := $(PROTO_DIR)/business_days.proto
GENERATED_DIR := $(PROTO_DIR)/generated

.PHONY: all proto build run clean

# デフォルトターゲット: proto を生成してからビルドする
all: proto build

# Buf を使って proto コードを生成
proto:
	buf generate $(PROTO_FILE)

# cmd ディレクトリ内の main.go をビルド
build:
	go build -o $(BINARY_NAME) ./cmd

# ビルドしたバイナリを実行
run: build
	./$(BINARY_NAME)

# 生成されたバイナリを削除
clean:
	rm -f $(BINARY_NAME)