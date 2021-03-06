# whatismyip
![GitHub](https://img.shields.io/github/license/beckxie/whatismyip?style=plastic)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fbeckxie%2Fwhatismyip.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fbeckxie%2Fwhatismyip?ref=badge_shield)

> This ia a web server that can show your real public IP address.

# Features

- Show your real public IP address.
- Show HTTP headers of your request.

# Demo 

![Website](https://img.shields.io/website?label=status&style=plastic&url=https%3A%2F%2Fip.beckxie.com)

[whatismyip](https://ip.beckxie.com) 

# Screenshots

![demo](screenshots/demo.png)

# Principles

Get first value of HTTP header `X-Forwarded-For`(XFF).

#### HTTP request (Example):

> X-Forwarded-For: `client`, proxy1, proxy2

> note: `HTTP header is modifiable`.

# Prerequisites

- [Go (at least Go 1.11)](https://golang.org/dl/)
- Proxy Server

  According to [Principles], `whatismyip` need the proxy server.
  (See [proxy server config example].)

[principles]: #Principles
[proxy server config example]: example/proxy_config

# Build

1. `git clone https://github.com/beckxie/whatismyip.git`
2. `make all`

# Deployment

1. Get the executable file by [build] or [download].
2. Run executable file.

### Usage

```
  -p int
        http server port (default 9999)
  -tmpl string
        tmpl file (default "../web/template/whatismyip.tmpl")
  -v    version
```

[download]: https://github.com/beckxie/whatismyip/releases
[build]: #Build

# References

- [HTTP headers | X-Forwarded-For](https://www.geeksforgeeks.org/http-headers-x-forwarded-for/)
- [HTTP 请求头中的 X-Forwarded-For](https://imququ.com/post/x-forwarded-for-header-in-http.html)

# License

This project is licensed under the MIT License - see the [LICENSE] file for details

[license]: ./LICENSE


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fbeckxie%2Fwhatismyip.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fbeckxie%2Fwhatismyip?ref=badge_large)

# Acknowledgments

[GitHub Corners - tholman/github-corners](https://github.com/tholman/github-corners)

[shields.io - badges/shields ](https://shields.io/)