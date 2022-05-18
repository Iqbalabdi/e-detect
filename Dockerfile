# FROM golang:1.17

# ## We create an /app directory within our
# ## image that will hold our application source
# ## files
# RUN mkdir /app

# ## We specify that we now wish to execute
# ## any further commands inside our /app
# ## directory
# WORKDIR /app

# COPY go.mod /app
# COPY go.sum /app
# RUN go mod download

# ## We copy everything in the root directory
# ## into our /app directory
# ADD . /app

# ## we run go build to compile the binary
# ## executable of our Go program
# RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o main ./app/main.go

# EXPOSE 8080
# ## Our start command which kicks off
# ## our newly created binary executable
# CMD ["./main"]

FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app ./app/main.go


# ------------------------------------


FROM alpine:latest

WORKDIR /kemasan

COPY --from=builder /app/main.app .

CMD [ "/kemasan/main.app" ]