FROM golang
RUN mkdir /stackcoder
WORKDIR  /stackcoder
COPY . .
RUN cp /stackcoder/protoc/protoc /go 
RUN export PATH=$PATH:/stackcoder/bin
RUN export PATH=$PATH:/stackcoder/protoc
RUN export PATH=$PATH:/go
RUN wget -q  https://raw.githubusercontent.com/micro/micro/master/scripts/install.sh -O - | /bin/bash
RUN go install github.com/micro/micro/v3@latest
EXPOSE 8080 
ENTRYPOINT ["micro", "run" , "." ]

