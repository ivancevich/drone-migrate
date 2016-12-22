# Docker image for the docker plugin
#
#     docker build --rm=true -t ivancevich/drone-migrate .

FROM gliderlabs/alpine

ADD drone-migrate /bin/
ENTRYPOINT ["/bin/drone-migrate"]
