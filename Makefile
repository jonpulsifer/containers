PROJECT ?= trusted-builds
MANIFEST ?= cloudbuild.yaml

.PHONY: all cloud-build
all: help

cloud-build: ci ## submits the build to $(PROJECT)
	gcloud builds submit \
	  --config=$(MANIFEST) \
	  --project=$(PROJECT)

ci: ## go run main.go
	@go run main.go || echo "Something is wrong with file $(MANIFEST) or go is not installed"

help: ## ty jessfraz
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
