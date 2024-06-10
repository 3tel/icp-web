FROM alpine:latest

ADD icpweb /home/app/
COPY conf /home/app/conf

WORKDIR /home/app
ENTRYPOINT ["./icpweb", "web"]

EXPOSE 9315