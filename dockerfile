# 기본 이미지를 Go 언어를 실행할 수 있는 Alpine Linux로 설정합니다.
FROM golang:alpine

# 작업 디렉토리를 생성하고, WORKDIR 명령어를 통해 작업 디렉토리로 이동합니다.
WORKDIR /app

# 소스 코드를 작업 디렉토리에 복사합니다.
COPY . .

# 필요한 Go 패키지를 다운로드합니다.
RUN go mod download

# 애플리케이션을 빌드합니다.
RUN go build -o main .

# 컨테이너 실행 시 애플리케이션을 실행합니다.
CMD ["/app/main"]
