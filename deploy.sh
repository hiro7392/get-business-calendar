PROJECT_ID=project-monkey-chair
REGION=asia-northeast1

IMAGE_NAME=gcr.io/${PROJECT_ID}/get-business-calendar
SERVICE_NAME=get-business-calendar

# use gcloud builds to build container image for linux/amd64
echo "Building Docker image..."
gcloud builds submit --config cloudbuild.yaml .

echo "Deploying to Cloud Run..."
gcloud run deploy ${SERVICE_NAME} \
  --image ${IMAGE_NAME} \
  --platform managed \
  --region ${REGION} \
  --allow-unauthenticated
echo "Deployment complete."