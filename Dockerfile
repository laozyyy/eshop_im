FROM golang:latest

WORKDIR /app

# 复制源代码
COPY . .

# 构建应用
RUN go build eshop_im

EXPOSE 9000
ENTRYPOINT ["./eshop_im"]