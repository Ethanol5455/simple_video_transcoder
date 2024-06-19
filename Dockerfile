FROM golang:bookworm

RUN apt update && \
    apt install -y ffmpeg

COPY . /src

WORKDIR /src

# RUN rm -rf /src

CMD ["go", "run", "."]
