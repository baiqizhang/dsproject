# 18842Team9
Distributed Uber Service

in case of error:
```
#!bash
# SuperNode
cd dsproject/supernode
go build . && ./supernode
# Util
cd dsproject/util
go build . 
# CarNode
cd dsproject/carnode
go run carnode.go 2.44 1.12 CAR1
# Server
go run server.go
```