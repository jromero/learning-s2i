BUILDER_IMAGE=s2i-golang
RUNTIME_IMAGE=s2i-golang-runtime
APP_IMAGE=my-go-app

.PHONY: builder
builder:
	docker build -t $(BUILDER_IMAGE) ./builder

.PHONY: runtime
runtime:
	docker build -t $(RUNTIME_IMAGE) ./runtime

.PHONY: build
build: builder runtime

.PHONY: test-builder
test-builder: builder
	s2i build --copy test-app/ $(BUILDER_IMAGE) $(APP_IMAGE) --incremental
	docker run --rm -it -p 8080:8080 $(APP_IMAGE)

.PHONY: test-with-runtime
test-with-runtime: builder runtime
	s2i build --copy test-app/ $(BUILDER_IMAGE) $(APP_IMAGE) --runtime-image $(RUNTIME_IMAGE) --runtime-artifact /opt/app-root/bin/app
	docker run --rm -it -p 8080:8080 $(APP_IMAGE)