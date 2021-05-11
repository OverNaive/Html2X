FROM golang:1.16-buster as build

WORKDIR /home/app

COPY src .

RUN go mod download \
&& go build html2x.go


FROM debian:buster-slim as prod

WORKDIR /home/app

COPY --from=build /home/app/html2x /home/app/html2x
COPY supervisor/html2x.conf /etc/supervisor/conf.d/html2x.conf
COPY ttf/times/ /usr/share/fonts/truetype/times/
COPY ttf/chinese/ /usr/share/fonts/truetype/chinese/

RUN chmod +x html2x \
&& apt-get update \
# install tools
&& apt-get install -y --no-install-recommends wget supervisor \
# install dependencies for wkhtmltox
&& apt-get install -y --no-install-recommends fontconfig libfreetype6 libjpeg62-turbo libpng16-16 libx11-6 libxcb1 libxext6 libxrender1 xfonts-75dpi xfonts-base ca-certificates \
&& wget -O wkhtmltox.deb --no-check-certificate https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb \
&& dpkg -i wkhtmltox.deb \
&& rm wkhtmltox.deb \
&& apt-get purge -y wget \
&& apt-get autoclean \
&& rm -rf /var/lib/apt/lists/

EXPOSE 8888

CMD [ "supervisord", "-c", "/etc/supervisor/supervisord.conf" ]