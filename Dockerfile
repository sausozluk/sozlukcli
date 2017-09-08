FROM centurylink/ca-certs

ADD bin/sozlukcli-linux /
ADD bin/config.json /

ENTRYPOINT ["/sozlukcli-linux"]