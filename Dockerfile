FROM scratch

ADD ./view /view/
COPY main.out /

EXPOSE 8220
ENTRYPOINT ["/main.out"]