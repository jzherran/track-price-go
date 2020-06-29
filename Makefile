all: bin/track-price-go
.PHONY: bin/track-price-go
bin/track-price-go:
	@docker	build	.	--target	bin	\
	--output bin/ \
	--platform ${PLATFORM}
