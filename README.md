

 
# IP Tracker CLI 

![cobra logo](https://cloud.githubusercontent.com/assets/173412/10886352/ad566232-814f-11e5-9cd0-aa101788c117.png)

Cobra is a library for creating powerful modern CLI applications. This project is a CLI application built using Cobra.

[![](https://img.shields.io/github/workflow/status/spf13/cobra/Test?longCache=tru&label=Test&logo=github%20actions&logoColor=fff)](https://github.com/spf13/cobra/actions?query=workflow%3ATest)
[![Go Reference](https://pkg.go.dev/badge/github.com/spf13/cobra.svg)](https://pkg.go.dev/github.com/spf13/cobra)
[![Go Report Card](https://goreportcard.com/badge/github.com/spf13/cobra)](https://goreportcard.com/report/github.com/spf13/cobra)

# Installing

Clone this project
```
https://github.com/mbakin/iptracker.git
```

Open project
```
cd iptracker
```

Using Cobra is easy. First, use `go get` to install the latest version
of the library.     

```
go get -u github.com/spf13/cobra@latest
```

Next, include Cobra in your application:

```go
import "github.com/spf13/cobra"
```



 

## Directory Structure
```
.iptracker
├── cmd
│   ├── root.go
│   ├── trace.go
│   ├── version.go   
├── go.mod
├── go.sum
├── main.go

```

  
## IP Information

IP Info Model:

```bash 
type Ip struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
}
```



## Commands

The following commands are executed using the `iptracker` command

![cobra logo](https://user-images.githubusercontent.com/68995469/163720217-37fbd31b-7a45-46fd-bacb-7a52713fac7b.png)


Run this command to get the address information of the IP address `iptracker trace X.X.X.X`

![cobra logo](https://user-images.githubusercontent.com/68995469/163720557-afcf0cbb-79dc-4d94-bf88-e4a2be9228bd.png)
## Licenses

[MIT](https://choosealicense.com/licenses/mit/)

  
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/build/)
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)

  
