
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/OverNaive/Html2X)](https://github.com/OverNaive/Html2X/releases)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/OverNaive/Html2X?filename=src%2Fgo.mod)](https://github.com/golang/go)
[![wkhtmltopdf](https://img.shields.io/badge/wkhtmltopdf-0.12.6-blue)](https://github.com/wkhtmltopdf/wkhtmltopdf)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/overnaive/html2x)](https://hub.docker.com/r/overnaive/html2x/builds)
[![GitHub](https://img.shields.io/github/license/OverNaive/Html2X)](https://github.com/OverNaive/Html2X/blob/main/LICENSE)

Html2X 是一个基于 [wkhtmltox](https://github.com/wkhtmltopdf/wkhtmltopdf) 实现的开箱即用的 http 服务，帮助服务端快速生成 pdf/image。

## 免费服务

如果你对 Docker 尚未了解，但又想快速使用；抑或你仅仅只想体验一下。

这里提供一个已经部署好的免费服务，请至 [免费的 Html2X 服务](https://overnaive.github.io/Html2X/docs/free) 查看。

## 项目目的

1. 以 http 服务代替第三方包，与业务系统解耦，可独立更新；
2. 将 [wkhtmltox](https://github.com/wkhtmltopdf/wkhtmltopdf) 的安装封装于 Docker 内，可快速更新版本；
3. 直接拉取镜像即可快速完成部署，真正的开箱即用。

## 如何使用

请先自行安装好 [Docker](https://www.docker.com/)

### 1. 获取镜像

- 本地构建镜像：`docker build -t overnaive/html2x`
- 远程拉取镜像：`docker pull overnaive/html2x`

### 2. 运行镜像

使用命令：`docker run --name html2x -p 127.0.0.1:8080:8888 -it -d overnaive/html2x`，即可运行一个容器。

此时，一个 http 服务已运行。

## API 文档

请至 [API for Html2X](https://overnaive.github.io/Html2X/docs/api) 查看。

## 相关文档

- Docker 学习资料：[Docker —— 从入门到实践](https://yeasy.gitbook.io/docker_practice/)
- [wkhtmltopdf 说明文档](https://overnaive.github.io/Html2X/docs/wkhtmltopdf.txt) 
- [wkhtmltoimage 说明文档](https://overnaive.github.io/Html2X/docs/wkhtmltopdf.txt)

## 更新计划

计划使用 [gin](https://github.com/gin-gonic/gin) 来实现参数验证、鉴权、限流等复杂逻辑。

## License

MIT
