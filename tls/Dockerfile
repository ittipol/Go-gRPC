FROM  busybox

WORKDIR /app

COPY server.crt server.pem ca.crt ./tls/

CMD [ "tail", "-F", "anyfile" ]