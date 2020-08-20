# ie11_tls1_2
Registry Edit changed for TLS 1.2 

**Summary:**

This program was created to add registry keys to enable TLS 1.2 security features in Window 7-10. This was written in Go, because it can be compiled into a portable executable binary. This ensures consistent execution and ability to sign if an organization has a signing mechanism for .exe.

**Instructions:**

In order to build .exe from source code.

1. Install [Go](https://golang.org/dl/) 
2. Pull source code. 
3. Get dependency: `go get golang.org/x/sys/windows/registry`
4. Go build: `go build .\TLS12Update_Server.go`
5. Run .exe as Administrator.
