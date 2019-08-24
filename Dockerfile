FROM scratch

WORKDIR /app

ADD ./main /app

EXPOSE 3000

ENTRYPOINT [ "./main" ]