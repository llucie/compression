FROM alpine:3.9

COPY compression /home/compression
COPY resources/image.jpg /home/resources/image.jpg

WORKDIR /home

ENTRYPOINT [ "./compression" ]

CMD [ "resources/image.jpg", "huff" ]