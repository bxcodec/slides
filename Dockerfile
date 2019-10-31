FROM golang
RUN mkdir /slides
WORKDIR /slides
RUN go get golang.org/x/tools/cmd/present
COPY . /slides/content/
RUN cp -r $GOPATH/src/golang.org/x/tools/cmd/present/* /slides
ENV GAE_ENV "standard"
CMD ["present", "-notes"]
