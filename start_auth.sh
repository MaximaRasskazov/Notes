#!/bin/bash

cd auth && \
 PORT=8101 \
 HOST=localhost \
 SERVER_TIMEOUT=10 \
 DB_TIMEOUT=5 \
 JWT_SECRET_KEY=secret_key \
 JWT_ACCESS_TOKEN_EXPIRATION=24 \
 JWT_REFRESH_TOKEN_EXPIRATION=168 \
 go run main.go


# cd auth && set PORT=9100 && set HOST=localhost set SERVER_TIMEOUT=10 set DB_TIMEOUT=5 set JWT_SECRET_KEY=secret_key  set JWT_ACCESS_TOKEN_EXPIRATION=24 set JWT_REFRESH_TOKEN_EXPIRATION=168   && go run main.go

#cd auth && set PORT=8101&& set HOST=localhost&& set SERVER_TIMEOUT=10&& set DB_TIMEOUT=5&& set JWT_SECRET_KEY=secret_key&& set JWT_ACCESS_TOKEN_EXPIRATION=24&& set JWT_REFRESH_TOKEN_EXPIRATION=168&& go run main.go