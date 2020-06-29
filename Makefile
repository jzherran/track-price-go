all: bin/track-price-go
.PHONY: bin/track-price-go
bin/traker-price-go:
	@docker	build	.	--target	bin	\
	--output bin/ \
	--platform local