# --- Steg 1: Bygg miljön ---
FROM golang:1.26-alpine AS builder

# Installera nödvändiga verktyg för att kompilera (git behövs ofta för go mod)
RUN apk add --no-cache git

# Sätt arbetsmapp
WORKDIR /app

# Kopiera beroendefiler först för att utnyttja Docker-cachen
COPY go.mod go.sum ./
RUN go mod download

# Kopiera resten av källkoden
COPY . .

# Bygg den körbara filen. 
# Byt ut './src' till '.' om main.go ligger i roten.
# Vi kompilerar för Linux (nödvändigt om du utvecklar på Windows/Mac)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src

# --- Steg 2: Produktions miljön ---
FROM alpine:latest

# Installera CA-certifikat (viktigt för https-anrop från din app)
RUN apk --no-cache add ca-certificates

# Sätt arbetsmapp
WORKDIR /root/

# Kopiera den körbara filen från bygg-steget
COPY --from=builder /app/main .

# Exponera port 8888
EXPOSE 8888

# Kör programmet
CMD ["./main"]   
