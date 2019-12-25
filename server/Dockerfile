##
# Filename: c:\CPSC\go\src\github.com\teamgenerator\teamgenerator\server\Dockerfile
# Path: c:\CPSC\go\src\github.com\teamgenerator\teamgenerator\server
# Created Date: Sunday, August 5th 2018, 7:14:34 pm
# Author: Nico Alimin
# 
# Copyright (c) 2018 Nico Alimin
##

FROM golang:1.14-rc-alpine3.10

# Environment Variables
ENV SRC_DIR=/go/src/github.com/teamgenerator/teamgenerator/server

# Metainformation
LABEL Author="Nico Alimin <nicoalimin@gmail.com>"

# Add the source code:
ADD . $SRC_DIR

# Change the directory to the source code
WORKDIR $SRC_DIR

# Install dependencies
RUN go get

# Install the binary
RUN go install

# Build it:
RUN cd $SRC_DIR; go build -o server

# Export ports
EXPOSE 3030

ENTRYPOINT [ "./server" ]
