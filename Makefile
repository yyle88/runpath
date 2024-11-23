COVERAGE_DIR ?= .coverage

# cp from: https://github.com/golang-migrate/migrate/blob/c378583d782e026f472dff657bfd088bf2510038/Makefile#L36
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
