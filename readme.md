# Simple Snap Midtrans - Golang

## 💫 About
Simple package to connect your golang app with midtrans.

##  🚀 Import
```shell
go get -u github.com/findryankp/snapMidtransGo@latest
```

## 👨🏽‍💻 Example
```go
import (
	"fmt"

	"github.com/Findryankp/snapMidtransGo"
)

func main() {
	var postData = snapMidtransGo.DataPostMidtrans{
		OrderId:   "order-001",
		Nominal:   1000,
		FirstName: "findryan",
		LastName:  "kurnia",
		Email:     "findryankp@gmail.com",
		Phone:     "081*******",
		ServerKey: "key server from midtrans",
	}

	//send your transaction
	paymenLink, err1 := snapMidtransGo.SandboxRequestSnapMidtrans(postData)
	if err1 != nil {
		panic(err1.Error())
	}

	fmt.Println(paymenLink)

	//get data from your transaction
	dataPayment, err2 := snapMidtransGo.SandboxGetTransaction(postData)
	if err2 != nil {
		panic(err2.Error())
	}

	fmt.Println(dataPayment)
}
```

## 🎯 Post your data
all data must be required

```go
var postData = snapMidtransGo.DataPostMidtrans{
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
snapMidtransGo.RequestSnapMidtrans(postData)

//get status transaction
snapMidtransGo.GetTransaction(postData)
```

- Sandbox
```go
//request transaction
snapMidtransGo.SandboxRequestSnapMidtrans(postData)

//get status transaction
snapMidtransGo.SandboxGetTransaction(postData)
```


## 🚀 Return data
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
## 👨🏽‍💻 Get your server key from Midtrans
1. Login
2. Setting
3. Access Key
4. Server Key
   - if you use sandbox, choose sandbox server key
   - likewise if you choose production choose server key production
   
## 😎 Development by
[![Findryankp](https://img.shields.io/badge/Findryankp-grey?style=for-the-badge&logo=github&logoColor=white)](https://github.com/findryankp)
[![findryankp](https://img.shields.io/badge/findryankp-blue?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/Findryankp/)

