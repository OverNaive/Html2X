FROM golang:1.16-buster as build

WORKDIR /home/app

COPY src .

# build
RUN go mod download \
&& go build html2x.go


FROM debian:buster-slim as prod

WORKDIR /home/app

# binary file
COPY --from=build /home/app/html2x /home/app/html2x

# supervisor setting
COPY supervisor/html2x.conf /etc/supervisor/conf.d/html2x.conf

# fonts
COPY fonts/ /usr/share/fonts/truetype/

RUN chmod +x html2x \
&& apt-get update \
# install tools
&& apt-get install -y --no-install-recommends wget supervisor \
# install dependencies for wkhtmltox
&& apt-get install -y --no-install-recommends fontconfig libfreetype6 libjpeg62-turbo libpng16-16 libx11-6 libxcb1 libxext6 libxrender1 xfonts-75dpi xfonts-base ca-certificates \
# get and install wkhtmltox
&& wget -O wkhtmltox.deb --no-check-certificate https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb \
&& dpkg -i wkhtmltox.deb \
&& rm wkhtmltox.deb \
# clean up for smaller size
&& apt-get purge -y wget \
&& apt-get autoclean \
&& rm -rf /var/lib/apt/lists/

EXPOSE 8888

CMD [ "supervisord", "-c", "/etc/supervisor/supervisord.conf" ]