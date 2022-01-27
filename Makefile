# Code generated by go.einride.tech/sage. DO NOT EDIT.
# To learn more, see .sage/sagefile.go and https://github.com/einride/sage.

.DEFAULT_GOAL := all

sagefile := .sage/tools/bin/sagefile

$(sagefile): .sage/go.mod $(shell find .sage/.. -type f -name '*.go')
	@cd .sage && go run ../cmd/build

.PHONY: clean-sage
clean-sage:
	@git clean -fdx .sage/tools

.PHONY: all
all: $(sagefile)
	@$(sagefile) All

.PHONY: convco-check
convco-check: $(sagefile)
	@$(sagefile) ConvcoCheck

.PHONY: format-markdown
format-markdown: $(sagefile)
	@$(sagefile) FormatMarkdown

.PHONY: format-yaml
format-yaml: $(sagefile)
	@$(sagefile) FormatYAML

.PHONY: git-verify-no-diff
git-verify-no-diff: $(sagefile)
	@$(sagefile) GitVerifyNoDiff

.PHONY: go-mod-tidy
go-mod-tidy: $(sagefile)
	@$(sagefile) GoModTidy

.PHONY: go-review
go-review: $(sagefile)
	@$(sagefile) GoReview

.PHONY: go-test
go-test: $(sagefile)
	@$(sagefile) GoTest

.PHONY: golangci-lint
golangci-lint: $(sagefile)
	@$(sagefile) GolangciLint
