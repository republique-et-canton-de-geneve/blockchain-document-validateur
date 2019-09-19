FROM nginx:alpine AS builder

# nginx:alpine contains NGINX_VERSION environment variable, like so:
# ENV NGINX_VERSION 1.15.0

# Our NCHAN version
ENV MORE_HEADERS_VERSION 0.33

# Download sources
RUN wget "http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz" -O nginx.tar.gz && \
  wget "https://github.com/openresty/headers-more-nginx-module/archive/v${MORE_HEADERS_VERSION}.zip" -O headers-more-nginx-module.zip

# For latest build deps, see https://github.com/nginxinc/docker-nginx/blob/master/mainline/alpine/Dockerfile
RUN apk add --no-cache --virtual .build-deps \
  gcc \
  libc-dev \
  make \
  openssl-dev \
  pcre-dev \
  zlib-dev \
  linux-headers \
  curl \
  gnupg \
  libxslt-dev \
  gd-dev \
  geoip-dev \
  zip

# Reuse same cli arguments as the nginx:alpine image used to build
RUN CONFARGS=$(nginx -V 2>&1 | sed -n -e 's/^.*arguments: //p') \
    mkdir /usr/src && \
	tar -zxC /usr/src -f nginx.tar.gz && \
  unzip "headers-more-nginx-module.zip" && \
  MOREHEADERSDIR="$(pwd)/headers-more-nginx-module-${MORE_HEADERS_VERSION}" && \
  cd /usr/src/nginx-$NGINX_VERSION && \
  ./configure --with-compat $CONFARGS --add-dynamic-module=$MOREHEADERSDIR && \
  make && make install

FROM jwilder/nginx-proxy:alpine
# Extract the dynamic module from the builder image
COPY --from=builder /usr/local/nginx/modules/ngx_http_headers_more_filter_module.so /usr/local/nginx/modules/ngx_http_headers_more_filter_module.so
COPY nginx-prod.conf /etc/nginx/nginx.conf

VOLUME ["/etc/nginx/certs", "/etc/nginx/dhparam"]
