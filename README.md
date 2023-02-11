# GO_DB_2023

[![](https://img.shields.io/badge/sqlitebroswer-3.12.2-white.svg)](https://sqlitebrowser.org/) 
[![](https://img.shields.io/badge/Go-1.19.3-blue.svg)](https://golang.org/pkg/)

Experiments with GoLang: 

1. Database + Domain
2. Database singleton client connection
3. Worker Pooling
4. Client-Server
5. SSL
6. Web server go routines

## Installation, Setup, and Use

Go, etc.

GCC (and make):

1. Install via [MSYS2](https://www.msys2.org/) on Windows. **MSYS2** still seems to be the best Windows Linux tools distribution around (2022). Will install a Linux terminal and add itself to **Environment Variables** automatically. Will come bundled with `pacman`.
2. I'd recommend the [toolchain](https://packages.msys2.org/package/mingw-w64-x86_64-gcc?repo=mingw64): `pacman -S mingw-w64-x86_64-toolchain`
3. You'll still need to add the `gcc` dir to the **Environment Variables** (`C:\msys64\mingw64\bin`) `PATH`.

> If you run into pacman PGP key trust errors per [this](https://github.com/msys2/MINGW-packages/issues/240):

```BASH
pacman-key --init
pacman-key --populate msys2
pacman-key --refresh-keys
```

Think the code repositories rotate keys every 3 months or so. Make sure to update those!

Execute the following commands to grab all the `Go` dependencies manually:

1. `go clean && go clean -modcache`
1. `go get github.com/gofrs/uuid`

For a valid self-signed SSL:

1. `openssl genrsa -out key.pem 2048`
1. `openssl req -new -sha256 -key key.pem -out csr.csr`
1. `openssl req -x509 -sha256 -days 365 -key key.pem -in csr.csr -out certificate.pem`

Navigate to [./goserver](./goserver):

1. `go run httpsServer.go`

You should see:

```BASH
=================== POLLING EVERY 5s ===================
=================== POLLING EVERY 5s ===================
=================== POLLING EVERY 5s ===================
2023/02/02 16:55:38 http: TLS handshake error from [::1]:60797: remote error: tls: unknown certificate
2023/02/02 16:55:40 http: TLS handshake error from [::1]:60802: remote error: tls: unknown certificate
=================== POLLING EVERY 5s ===================
=================== POLLING EVERY 5s ===================
=================== POLLING EVERY 5s ===================
=================== POLLING EVERY 5s ===================
```

## Views and API Endpoints

1. https://localhost/public/
2. https://localhost/public/queryworker.html
3. https://localhost/public/querystatus.html
4. https://localhost/public/querystatuses.html
5. https://localhost/public/addworker.html
6. https://localhost/public/stopworker.html
7. https://localhost/api/examples
 