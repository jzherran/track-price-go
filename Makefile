.PHONY:
run:
	@docker-compose up -d --build

.PHONY:
lint: gopkgcache
	@echo ">> running golangci-lint"
	@golangci-lint run ./...

.PHONY:
gopkgcache:
# 'go list' needs to be executed before staticcheck to prepopulate the modules cache.
# Otherwise staticcheck might fail randomly for some reason not yet explained.
	@go list -e -compiled -test=true -export=false -deps=true -find=false -tags= -- ./... > /dev/null