# SENTENCEPIECEGO
This is a golang interface of [sentencepiece](https://github.com/google/sentencepiece) for serving. For more information, please follow google sentencepiece.

# Usages
1. Download shared library
Please check sentencepiece library first with pip and spm-train. 

If you use amd64, you can download v0.1.98-x86-64 in releases, which libsentencepiecego.so is built based on google/sentencepiece of 0.1.98.
```bash
sudo wget https://github.com/activezhao/sentencepiecego/releases/download/v0.1.98-x86-64/libsentencepiecego.so -P /usr/local/lib/
ldconfig
```
If you use arm64, you can download v0.1.98-arm-64 in releases, which libsentencepiecego.so is built based on google/sentencepiece of 0.1.98.
```bash
sudo wget https://github.com/activezhao/sentencepiecego/releases/download/v0.1.98-arm-64/libsentencepiecego.so -P /usr/local/lib/
ldconfig
```
2. Export CGO variable
```
export CGO_CXXFLAGS=-std=c++11
```
3. Go get
```
go get github.com/evan176/sentencepiecego
```

```go
package main

import (
        "fmt"

        "github.com/evan176/sentencepiecego"
)

func main() {
        // Load pre-trained spm model
        m, err := sentencepiecego.Load("spm.model")
        if err != nil {
                panic(err)
        }
        // Encode text to ids([]int) with spm model
        ids, err := m.Encode("test")
        if err != nil {
                panic(err)
        }
        fmt.Printf("%+v\n", ids)
        // Release model before exit
        m.Free()
}
~
```

# References
https://github.com/google/sentencepiece
