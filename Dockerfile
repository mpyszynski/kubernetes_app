# https://github.com/docker-library/repo-info/blob/master/repos/alpine/remote/3.13.0.md
FROM alpine@sha256:d0710affa17fad5f466a70159cc458227bd25d4afb39514ef662ead3e6c99515

ARG binary_name=server

ADD ${binary_name} /usr/local/bin/service

ENTRYPOINT ["/usr/local/bin/service"] 