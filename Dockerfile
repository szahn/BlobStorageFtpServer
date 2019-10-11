FROM golang

WORKDIR /go/src/blob_storage_ftp

COPY . .

RUN go get -d -v ./
RUN go install -v ./

# Allow sharing aws credential config
ENV AWS_SDK_LOAD_CONFIG true

CMD ["blob_storage_ftp"]