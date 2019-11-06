# Docker RBL aka DNSBL checker on Alpine Linux
Example Setup of custom shell script for Docker, build on [Alpine Linux](http://www.alpinelinux.org/).
The image is only +/- 10MB large.

Repository: https://github.com/amitsdalal/dnsbl


* Built on the lightweight and secure Alpine Linux distribution
* Very small Docker image size (+/-10MB)
* Uses dig and curl with bash env
* Follows the KISS principle (Keep It Simple, Stupid) to make it easy to understand and adjust the image
* You will get the output aginst the DNSBL name if your IP is blacklisted there.

[![Docker Pulls](https://img.shields.io/docker/pulls/amitdalal/dnsbl.svg)](https://hub.docker.com/r/amitdalal/dnsbl/)
![License Apache](https://img.shields.io/badge/license-apache-blue.svg)


## Usage

Start the Docker container:

    docker run -it --entrypoint=bash amitdalal/dnsbl dnsbl <IP address to be checked against the DNSBL>
