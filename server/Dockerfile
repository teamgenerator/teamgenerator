##
# Filename: c:\CPSC\go\src\github.com\teamgenerator\teamgenerator\server\Dockerfile
# Path: c:\CPSC\go\src\github.com\teamgenerator\teamgenerator\server
# Created Date: Sunday, August 5th 2018, 7:14:34 pm
# Author: Nico Alimin
# 
# Copyright (c) 2018 Nico Alimin
##

FROM iron/go:dev

# Environment Variables
ENV SRC_DIR=/go/src/github.com/teamgenerator/teamgenerator/server

# Install the dependency manager

RUN curl -s https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install tools for analysis
# RUN go get -u mvdan.cc/interfacer
# RUN go get github.com/GoASTScanner/gas/cmd/gas/...
# RUN go get github.com/jgautheron/goconst/cmd/goconst
# RUN go get github.com/jgautheron/gogetimports
# RUN go get github.com/jgautheron/usedexports
# RUN go get github.com/walle/lll/...
# RUN go get github.com/mdempsky/maligned
# RUN go get -u github.com/alexkohler/prealloc
# RUN go get github.com/gordonklaus/ineffassign
# RUN go get github.com/fzipp/gocyclo
# RUN go get -u github.com/tsenart/deadcode

# Metainformation
LABEL Author="Nico Alimin <nicoalimin@gmail.com>"

# Add the source code:
ADD . $SRC_DIR

# Change the directory to the source code
WORKDIR $SRC_DIR

# Install dependencies
RUN dep ensure

# Install the binary
RUN go install

# Build it:
RUN cd $SRC_DIR; go build -o myapp

# Export ports
EXPOSE 8080

ENTRYPOINT ["./myapp"]
