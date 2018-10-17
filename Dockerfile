FROM alpine:latest
WORKDIR /
COPY conf /conf
COPY docs /docs
COPY OnlineJudge /OnlineJudge
EXPOSE 3030
ENTRYPOINT ["./OnlineJudge"]