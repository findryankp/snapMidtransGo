# Simple Snap Go - Midtrans
Simple package to connect your golang app with midtrans.

**Development by:** 
- Findryankp

## Import
```shell
go get github.com/Findryankp/snapGoMidtrans
```

## Post your data
all data must be required

```go
var postData = snapGoMidtrans.DataPostMidtrans{
    OrderId:   "order-001",
    Nominal:   1000,
    FirstName: "findryan",
    LastName:  "kurnia",
    Email:     "findryankp@gmail.com",
    Phone:     "081*******",
    ServerKey: "key server from midtrans",
}
```

to call function request transaction
- Production
```go
//request transaction
snapGoMidtrans.RequestSnapMidtrans(postData)

//get status transaction
snapGoMidtrans.GetTransaction(postData)
```

- Sandbox
```go
//request transaction
snapGoMidtrans.SanboxRequestSnapMidtrans(postData)

//get status transaction
snapGoMidtrans.SanboxGetTransaction(postData)
```


## Return data
success
```json
{
    "response_data"
}
```
bad request
```json
{
    "error message"
}
```
## Get your server key from Midtrans
1. Login
2. Setting
3. Access Key
4. Server Key
   - if you use sandbox, choose sandbox server key
   - likewise if you choose production choose server key production

