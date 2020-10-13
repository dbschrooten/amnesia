FROM rancher/cli2:v2.3.2

RUN apk add curl && \
    curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.15.9/bin/linux/amd64/kubectl && \
    chmod +x kubectl && \
    mv ./kubectl /usr/local/bin/kubectl