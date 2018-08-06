##
# Filename: c:\CPSC\go\src\github.com\teamgenerator\teamgenerator\server\Dockerfile
# Path: c:\CPSC\go\src\github.com\teamgenerator\teamgenerator\server
# Created Date: Sunday, August 5th 2018, 7:14:34 pm
# Author: Nico Alimin
# 
# Copyright (c) 2018 Nico Alimin
##

FROM iron/go:dev

###
# Build arguments
###

# Environment Variables
ENV SRC_DIR=/go/src/github.com/teamgenerator/teamgenerator/server

###
# Metainformation
###

LABEL Author="Nico Alimin <nicoalimin@gmail.com>"

# Add the source code:
ADD . $SRC_DIR

# Change the directory to the source code
WORKDIR $SRC_DIR

# Install dependencies
# RUN dep ensure

# Install the binary (into /go/bin/)
RUN go install

# Build it:
RUN cd $SRC_DIR; go build -o myapp

# Export ports
EXPOSE 8080

ENTRYPOINT ["./myapp"]
