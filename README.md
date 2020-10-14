# SpamWatch API Wrapper For Golang

#### Install from go get

```
$ go get -v github.com/soekarnohatta/go-spamwatch
```

#### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/soekarnohatta/go-spamwatch/spamwatch"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
)

func main() {
	// Initialize Log Configuration
	newZapConfig := zap.NewProductionEncoderConfig()
	newZap := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(newZapConfig), os.Stdout, zap.InfoLevel))
	newSugarZap := newZap.Sugar()

	// Initialize token
	token := "PUT YOUR TOKEN HERE"

	// Initialize New Client
	client, _ := spamwatch.NewClient(newSugarZap, "", token)

	// Start requesting to the API
	bannedUser, _ := client.GetBan(123456789)
	if bannedUser != nil {
		fmt.Println("UserID:" + strconv.Itoa(bannedUser.Id))
		fmt.Println("Reason:" + bannedUser.Reason)
		fmt.Println("Message:" + bannedUser.Message)
	}
}
```