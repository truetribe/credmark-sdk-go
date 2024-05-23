TAG ?=  $(shell git -C "$(PROJECT_PATH)" rev-parse HEAD)
ifdef VERSION
override TAG = $(VERSION)
endif
PROJECT_PATH := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

.PHONY: test
test: test-unit test-integ

.PHONY: test-unit
test-unit: ## Runs the unit tests
	go test -short ./... -coverprofile $(PROJECT_PATH)/_out/unit-cover.out

.PHONY: test-unit-coverage
test-unit-coverage: create-output-dir test-unit
test-unit-coverage: ## Get the project test coverage
	go tool cover -func=$(PROJECT_PATH)/_out/unit-cover.out

.PHONY: create-output-dir
create-output-dir: ## Creates the output directory if it does not exist
	mkdir $(PROJECT_PATH)/_out || true

#integration test
.PHONY: test-integ
test-integ: ## Runs the unit tests
	go test -v ./... -run ^TestIntegration* -coverprofile $(PROJECT_PATH)/_out/integ-cover.out

.PHONY: test-integ-coverage
test-integ-coverage: create-output-dir test-integ
test-integ-coverage: ## Get the project test coverage
	go tool cover -func=$(PROJECT_PATH)/_out/integ-cover.out

help:
	@IFS=$$'\n' ; \
    help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
    for help_line in $${help_lines[@]}; do \
        IFS=$$'#' ; \
        help_split=($$help_line) ; \
        help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
        help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
        printf "%-30s %s\n" $$help_command $$help_info ; \
    done