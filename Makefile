all: build

.PHONY: init
init:
	git submodule update --init --recursive

.PHONY: update-submodule
update-submodule: init
	# Fetch the latest tags
	cd noVNC && git fetch --tags
	# Get the latest tag
	$(eval LATEST_TAG := $(shell cd noVNC && git describe --tags `git rev-list --tags --max-count=1`))
	@echo "Latest tag for noVNC: $(LATEST_TAG)"
	# Checkout the latest tag
	cd noVNC && git checkout $(LATEST_TAG)
	@echo "Updated submodule noVNC to latest tag: ${LATEST_TAG}"

.PHONY: clean
clean:
	rm -rf dist/*

.PHONY: build
build: clean
	cp -r noVNC/dist/* dist/