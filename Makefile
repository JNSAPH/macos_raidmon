EXECUTABLE =moss_raidmon
BUILD_DIR = bin
RELEASE_DIR = release

build:
	go build -o bin/$(EXECUTABLE) 

run: build
	./bin/$(EXECUTABLE)

dev: build
	./bin/$(EXECUTABLE) --dev

release: build
	rm -rf $(RELEASE_DIR)
	mkdir -p $(RELEASE_DIR)
	cp $(BUILD_DIR)/$(EXECUTABLE) $(RELEASE_DIR)/$(EXECUTABLE)
	cp -r templates $(RELEASE_DIR)/templates
	cp -r assets $(RELEASE_DIR)/assets
	cp -r config $(RELEASE_DIR)/config
	cp -r for_release/* $(RELEASE_DIR)/

clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(RELEASE_DIR)