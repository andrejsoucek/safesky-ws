.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make <target>\n"} /^[A-Za-z-]+:.*?##/ { printf "  %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: build run


build: ## Build container
	docker build -t safesky-ws .

run: ## Run app
	docker run -p 8000:8000 --name safesky-ws -d --restart unless-stopped safesky-ws
