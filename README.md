# Kira Env
For reading env variables from a yaml file.

# Installing...
1. Download and install it:

```sh
$ go get -u github.com/Lafriakh/env
```

2. Import it in your code:

```go
import "github.com/Lafriakh/env"
```

# Customize

1. Default file path
by default the config file path is: `config.yaml` to change it change this variable value.
```go
env.DefaultPath = "custom/file.yaml"
```

# Simple config file
```yaml
env: development
development:
    SITE_URL: env.dev
    SERVER_PORT: 8080
production: 
    SITE_URL: env.com
    SERVER_PORT: 80
test:
    SITE_URL: env.test
    SERVER_PORT: 8080
```
# Start using it
1. default usage
```go
package main
import "github.com/Lafriakh/env"

func main() {
    // init env with default configs
    env.New()
    // get value
    env.Get("SITE_URL").(string) // env.dev
    // get string value
    env.GetString("SITE_URL") // env.dev
    // get integer value
    env.GetInt("SERVER_PORT") // 80
}
```
2. customize usage
```go
package main
import "github.com/Lafriakh/env"

func main() {
    // init from custome file path
    env.SetPAL("custom/file/path.yaml")
    // get string value
    env.GetString("SITE_URL") // env.dev
    //...

    
    // or this
    env.DefaultPath = "custom/file/path.yaml"
    // init env with default configs
    env.New()
    // get string value
    env.GetString("SITE_URL") // env.dev
}
```