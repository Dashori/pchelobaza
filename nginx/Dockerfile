FROM debian:12.2-slim AS builder

RUN apt-get update
RUN apt-get install wget autoconf gcc libpcre3 libpcre3-dev zlib1g zlib1g-dev cmake  -y

WORKDIR /build

RUN wget https://github.com/openresty/headers-more-nginx-module/archive/refs/tags/v0.35.tar.gz && \    
    tar xzf v0.35.tar.gz

RUN wget http://nginx.org/download/nginx-1.22.1.tar.gz && \
    rm -rf nginx-1.22.1 && \
    tar xzf nginx-1.22.1.tar.gz

WORKDIR nginx-1.22.1
RUN ./configure \
    --with-compat \
    $(nginx -V) \
    --add-dynamic-module=/build/headers-more-nginx-module-0.35 && \
    make && \
    make install

FROM nginx:1.22.1

COPY --from=builder /build/nginx-1.22.1/objs/* /etc/nginx/modules/

COPY --from=builder /usr/local/lib/* /lib/x86_64-linux-gnu/

CMD ["nginx", "-g", "daemon off;"]
