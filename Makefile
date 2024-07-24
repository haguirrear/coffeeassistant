IMAGE=coffe-assistant

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: docker-image
docker-image:
	docker build -t $(IMAGE) --target prod .

.PHONY: docker-run
docker-run:
	docker run --rm --name $(IMAGE) -p 8000:8000 $(IMAGE)

.PHONY: dev
dev:
	bun run dev & air
