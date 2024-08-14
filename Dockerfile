
FROM golang:latest as builder

ARG GITHUB_TOKEN

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN export GOPATH=$HOME/go && export PATH=$PATH:$GOPATH/bin

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

RUN swag init --parseDependency --parseInternal --generatedTime --quiet

FROM openjdk:8-jre-alpine

RUN apk add --no-cache wkhtmltopdf

WORKDIR /app
COPY --from=builder /app/docs/swagger.json ./docs/swagger.json
COPY --from=builder /app/api ./

ENV MODE=production
ENV APP_PORT=80
ENV APP_NAME=${APP_NAME}
ENV APP_HOST=${APP_HOST}

ENV DATABASE_PORT=5432
ENV DATABASE_NAME=${DATABASE_NAME}
ENV DATABASE_USERNAME=${DATABASE_USERNAME}
ENV DATABASE_PASSWORD=${DATABASE_PASSWORD}
ENV DATABASE_TIMEZONE=America/Sao_Paulo

ENV MESSAGE_BIRD_KEY=${MESSAGE_BIRD_KEY}
ENV SEND_GRID_KEY=${SEND_GRID_KEY}

EXPOSE 80

CMD [ "./api"]