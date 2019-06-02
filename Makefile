PROJECT_ID := ojichat-api-251807
SERVICE_NAME := ojichat-api
DOCKER_IMAGE := gcr.io/$(PROJECT_ID)/$(SERVICE_NAME)
REGION := us-central

gcloud/build:
	gcloud builds submit --project ${PROJECT_ID} --tag $(DOCKER_IMAGE)

gcloud/deploy:
	gcloud beta run deploy $(SERVICE_NAME) --project $(PROJECT_ID) --image $(DOCKER_IMAGE) --platform managed -q
