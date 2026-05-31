# Go backend server

- mux router
- file sever on "/static"

## Kör test

```bash
go test -v ./src/ ./src/passwd/
```

## Bygg docker image:

```bash
docker build -t go-backend .
```

## Kör containern med port-mappning:

```bash
docker run -p 8888:8888 go-backend
```
