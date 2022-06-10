FROM docker:20.10.14-dind

ADD ./drone-lark /bin/

ENTRYPOINT ["/usr/local/bin/dockerd-entrypoint.sh", "/bin/drone-lark"]