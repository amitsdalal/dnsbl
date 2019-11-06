FROM alpine:3.10
LABEL Maintainer="Amit S Dalal <amit@amitdalal.me>" \
      Description="Lightweight container with dig and bash."

# Install packages
RUN apk --no-cache add curl bash bind-tools

# Configure DNSBL

RUN curl -s https://raw.githubusercontent.com/amitsdalal/dnsbl/master/dnsbl.sh > /bin/dnsbl

# Make sure it is excuteable.
RUN chmod +x /bin/dnsbl

CMD ["/bin/dnsbl"]

