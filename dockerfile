# Docker 이미지를 빌드할 때 사용할 base 이미지
FROM golang:alpine AS builder

# 작업 디렉토리 생성
RUN mkdir /app
WORKDIR /app

# 로컬에 있는 모든 파일을 컨테이너에 복사
COPY . .

# Go 패키지 다운로드
RUN go mod download

# 애플리케이션 빌드
RUN go build -o main .

# 포트 번호 설정
EXPOSE 8080

# 컨테이너 시작 시 실행할 명령어
CMD ["./main"]
