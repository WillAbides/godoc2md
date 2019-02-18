

# goreadme
`import "github.com/WillAbides/godoc2md/goreadme"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func ReadmeMD(pkgName string, writer io.Writer) error](#ReadmeMD)
* [func VerifyReadme(pkgName, readmePath string) (bool, error)](#VerifyReadme)
* [func WriteReadme(pkgName, readmePath string) (err error)](#WriteReadme)


#### <a name="pkg-files">Package files</a>
[goreadme.go](./goreadme.go) 





## <a name="ReadmeMD">func</a> [ReadmeMD](./goreadme.go?s=1026:1079#L50)
``` go
func ReadmeMD(pkgName string, writer io.Writer) error
```
ReadmeMD writes readme content for the given package to writer



## <a name="VerifyReadme">func</a> [VerifyReadme](./goreadme.go?s=606:665#L31)
``` go
func VerifyReadme(pkgName, readmePath string) (bool, error)
```
VerifyReadme checks that the file at readmePath has the correct content for pkgName



## <a name="WriteReadme">func</a> [WriteReadme](./goreadme.go?s=262:318#L18)
``` go
func WriteReadme(pkgName, readmePath string) (err error)
```
WriteReadme writes a README.md for pkgname to the given path








