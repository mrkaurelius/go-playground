FROM ethereum/client-go:release-1.10

RUN apk add --no-cache bash

RUN mkdir -p /root/.ethereum/bkys
COPY sgeth /root/

ENTRYPOINT [ "/bin/bash", "/root/rungeth" ]