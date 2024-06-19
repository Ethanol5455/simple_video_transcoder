#!/bin/bash

if docker build . -t simple_video_transcoder:latest ;
then
    docker run --rm -it -p 8080:8080 simple_video_transcoder:latest
fi
