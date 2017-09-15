FROM centurylink/ca-certs

ADD bin/sozluk /
ADD bin/sozluk.ini /

ENTRYPOINT ["/sozluk"]
