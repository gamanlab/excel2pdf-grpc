# EXCEL2PDF-GRPC

This is running like tools to convert ms excel file to pdf over gRPC API.

Tested on windows 10 Pro 64bit with ms. office pro plus 2019

> Please don't run this application as service

Default gRPC port is `8345`.

## How to use

Use custom port:

```bash
 ./excel2pdf-grpc -port CUSTOM_PORT_NUMBER
```

Compile proto file

```bash
protoc -I  . *.proto --go_out=plugins=grpc:pb
```
