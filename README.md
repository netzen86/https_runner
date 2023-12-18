# https_runner
# run app after https request

## compile app
```go build -o http_run```

## compile app for windows in linux
```env GOOS=windows GOARCH=amd64 go build -o http_run_w64.exe```

## Generate private key
```openssl ecparam -genkey -name secp384r1 -out server.key```

## Generation of self-signed public key bsed on the private
```openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650```

## run app, arg contains application for running
```http_run gedit```

## test app
```curl -k https://localhost:8888/?param=test.txt```

After https runner receive http request, https runner run 'gedit' with arg 'test.txt'.
