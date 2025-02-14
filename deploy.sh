PROJECT_ID=project-monkey-chair
REGION=asia-northeast1
# Docker イメージ名とサービス名を設定
IMAGE_NAME=gcr.io/${PROJECT_ID}/get-business-calendar
SERVICE_NAME=get-business-calendar

# Docker イメージのビルド
echo "Building Docker image..."
docker build -t ${IMAGE_NAME} .

Docker イメージを Google Container Registry に push
echo "Pushing Docker image..."
docker push ${IMAGE_NAME}

# Cloud Run にデプロイ
echo "Deploying to Cloud Run..."
gcloud run deploy ${SERVICE_NAME} \
  --image ${IMAGE_NAME} \
  --platform managed \
  --region ${REGION} \
  --allow-unauthenticated \
  --use-http2 \
  --timeout=300s

echo "Deployment complete."