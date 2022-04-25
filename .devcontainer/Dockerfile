FROM golang:1.18

RUN apt-get update && apt-get install -y zsh

RUN go install github.com/ramya-rao-a/go-outline@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

COPY --from=migrate/migrate:v4.15.1 /usr/local/bin/migrate /usr/local/bin/migrate
RUN sh -c "$(wget https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)"
