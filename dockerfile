# build stage

# get golang
FROM golang:1.18-alpine AS builder
# working directory
WORKDIR /app
#copy all files to work directory
COPY . .
# build executables files
RUN go build -o main server.go



# run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .


EXPOSE 8081
CMD ["/app/main"]