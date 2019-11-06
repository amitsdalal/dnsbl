FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o runtime .


FROM alpine:latest
LABEL Maintainer="Amit S Dalal <amit@amitdalal.me>" \
      Description="Lightweight container with dig and bash and a runtime"

# Install packages
RUN apk --no-cache add curl bash bind-tools coreutils &&\
curl -sSL -o /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub &&\
curl -sSL -O https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.28-r0/glibc-2.28-r0.apk &&\
apk add glibc-2.28-r0.apk &&\
rm glibc-2.28-r0.apk

# Configure DNSBL

COPY dnsbl.sh  /bin/dnsbl

# Make sure it is excuteable.
RUN chmod +x /bin/dnsbl

RUN groupadd -g 999 appuser && \
    useradd -r -u 999 -g appuser appuser
USER appuser


COPY --from=builder /app/runtime /bin

EXPOSE 5000

CMD ["/bin/runtime"]

