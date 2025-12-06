module backend

go 1.24.0

toolchain go1.24.9

require github.com/mattn/go-sqlite3 v1.14.32

require github.com/google/uuid v1.6.0

require (
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/joho/godotenv v1.5.1
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

require (
	golang.org/x/crypto v0.43.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
)
