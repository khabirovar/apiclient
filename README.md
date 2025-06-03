Api Client
================================
Simple api client implementation.

See it in action:

First of all get token from [SignUp Page](https://reqres.in/signup).
Add token to environment

```bash
export TOKEN=<your-token>
```

Run some code
```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/khabirovar/apiclient"
)

func main() {
	token := os.Getenv("TOKEN")

	apiclient, err := apiclient.NewApiClient(time.Second*20, token)
	if err != nil {
		log.Fatal(err)
	}

	id, err := apiclient.AddUser("Johnathan", "SysOps")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New user id: %d\n", id)

	user, err := apiclient.GetUser(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", user)
}
```