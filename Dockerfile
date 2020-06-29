FROM --platform=${BUILDPLATFORM} golang:1.14.3-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/track-price-go .

FROM scratch AS bin-unix
COPY --from=build /out/track-price-go /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/track-price-go /track-price-go.exe

FROM bin-${TARGETOS} AS bin
