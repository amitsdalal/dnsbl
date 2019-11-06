FROM alpine:3.10
LABEL Maintainer="Amit S Dalal <amit@amitdalal.me>" \
      Description="Lightweight container with Nginx 1.16 & PHP-FPM 7.3 based on Alpine Linux."

# Install packages
RUN apk --no-cache add dig curl bash wget

# Configure DNSBL

RUN mkdir -p /app
RUN cd /app
RUN wget https://gist.githubusercontent.com/amitsdalal/c32329ae3075634e0941ba46b5b9c4f0/raw/8562a4d18aa5e63ad88182b3d45e996d70e6647a/dnsbl.sh

# Make sure files/folders needed by the processes are accessable when they run under the nobody user
RUN chmod dnsbl.sh

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]

