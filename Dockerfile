FROM --platform=${BUILDPLATFORM} golang:1.14.3-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/track-price-go .

FROM scratch AS bin
COPY --from=build /out/track-price-go /