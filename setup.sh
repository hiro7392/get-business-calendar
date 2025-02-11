#!/bin/bash
set -e

# プロジェクトIDとリージョン（東京）を設定
PROJECT_ID=project-monkey-chair      # ご自身のプロジェクトIDに置き換えてください
REGION=asia-northeast1           # 東京リージョン

# gcloud のプロジェクトと Cloud Run のリージョンを設定
gcloud config set project "${PROJECT_ID}"
gcloud config set run/region "${REGION}"

# 設定内容を確認
echo "Using project: $(gcloud config get-value project)"
echo "Using region: $(gcloud config get-value run/region)"
