# image 
This is used for rapid development of digital image processing in Go.

## Quick Start
#### Download and install

```
go get github.com/ljq0002/image
```

#### Create file `hello.go`
```go
package main

import "github.com/ljq0002/image/imageProcess"

func main(){
	image, err := imageProcess.CreateImageFromFile("resource/lena.jpeg")
	if err != nil {
		panic(err)
	}
	image.GaussianBlur(5, 1.5).SaveToJpegFile("blur.jpeg")    
}
```

#### Build and run
```
go build hello.go
./hello
```

## Features
* Blur
  - Gaussian
  - Median
* Coming SOON

## License
ljq0002/image source code is licensed under the MIT License.
