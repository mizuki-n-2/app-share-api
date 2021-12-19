# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/app-share-api
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/app-share-api
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app-share-api

RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]