
FROM golang as builder

ARG APP_NAME
ARG GITHUB_TOKEN

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

FROM scratch

ARG APP_NAME

COPY --from=builder /app/${APP_NAME} /app

ENV MODE=production
ENV INSTANCE=1
ENV NAME=${APP_NAME}
ENV DATABASE_PORT=5433
ENV DATABASE_TIMEZONE=America/Sao_Paulo

EXPOSE 80

CMD [ "/app" ]