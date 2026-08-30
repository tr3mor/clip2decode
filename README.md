## Clip2Decode

### Description:
This is a simple tool to decode base64 from your clipboard, print it to terminal and also put it back to clipboard.  
Pretty useful when working with [Vault Transit Secret engine](https://www.vaultproject.io/docs/secrets/transit).  
**Works only for MacOS**

### Installation

#### Using go install
```
go install github.com/tr3mor/clip2decode/cmd/clip2decode@latest
```
This installs the `clip2decode` binary to `$(go env GOPATH)/bin` (make sure that's in your `$PATH`).

#### From source
```
git clone https://github.com/tr3mor/clip2decode.git
cd clip2decode
make
mv bin/c2d <dir in your $PATH>/clip2decode
```