
# Docker RBL aka DNSBL checker on Alpine Linux


* Built on the lightweight and secure Alpine Linux distribution
* Very small Docker image size (+/-25MB)
* Uses dig and curl with bash env
* Follows the KISS principle (Keep It Simple, Stupid) to make it easy to understand and adjust the image
* You will get the output aginst the DNSBL name if your IP is blacklisted there.

[![Docker Pulls](https://img.shields.io/docker/pulls/amitdalal/dnsbl.svg)](https://hub.docker.com/r/amitdalal/dnsbl/)
![License Apache](https://img.shields.io/badge/license-apache-blue.svg)


## Usage

You can simply check your IP using docker run, by :

    docker run -it --entrypoint=dnsbl amitdalal/dnsbl 1.1.1.1

To run the app :

    docker run -p 5000:5000 amitdalal/dnsbl

    Now try curl localhost:5000/1.1.1.1

Or even simple, just do :

    curl https://dnsbl.webscoot.io/1.1.1.1

#### Replace 1.1.1.1 with the <IP address to be checked against the DNSBL>