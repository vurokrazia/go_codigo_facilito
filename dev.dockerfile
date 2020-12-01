FROM golang:1.15.2-buster
WORKDIR /app
RUN go mod init go_course_cf
COPY . .
EXPOSE 3000
CMD bash