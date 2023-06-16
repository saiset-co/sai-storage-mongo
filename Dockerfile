# Build stage
FROM golang as BUILD

WORKDIR /src/

COPY ./ /src/

RUN go build -o service-bin -buildvcs=false

FROM ubuntu

WORKDIR /srv

# Copy binary from build stage
COPY --from=BUILD /src/service-bin /srv/service-bin

RUN chmod +x /srv/service-bin

# Set command to run your binary
CMD /srv/service-bin start