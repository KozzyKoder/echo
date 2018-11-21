# GO (TCP) Echo

A Simple go TCP echo server.

## Run application

Run the container from a terminal:
```bash
docker build -t kozzykoder/echo:0.1 . -f Dockerfile
docker run -it -e TCP_PORT=6558 -u appuser -p 6558:6558 kozzykoder/echo:0.1
```

In another terminal run:
```bash
telnet localhost 6558
```