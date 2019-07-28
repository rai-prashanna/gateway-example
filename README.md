# go-vendor-example
An example that using "govendor" into your project.

## How to use
  Step 1: Clone repo and setup GOPATH.
```
git clone https://github.com/kilfu0701/go-vendor-example.git
cd go-vendor-example
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH
```

  Setp 2: Install `govender`
```
go get -u github.com/kardianos/govendor
rm -rf src/github.com src/golang.org     # after install govender, remove "src/github.com" & "src/golang.org"
```

  Step 3: Download vendor packages
```
cd src/app
govendor sync
```

  Step 4: Run application, and check `localhost:8081`
```
cd $GOPATH
go run src/main.go
```

## Support version
  - Go 1.5 (run with `GO15VENDOREXPERIMENT=1 go run src/main.go`)
  - Go 1.6 above

## Contributors
kilfu0701 (kilfu0701@gmail.com)

## License
MIT




	// change the request host to match the target
//	r.Host = "www.example.org:80"
	r.Host = "www.example.org:80"
	u, _ := url.Parse("http://www.example.org")
	proxy := httputil.NewSingleHostReverseProxy(u)
	// You can optionally capture/wrap the transport if that's necessary (for
	// instance, if the transport has been replaced by middleware). Example:
	// proxy.Transport = &myTransport{proxy.Transport}
	proxy.Transport = &myTransport{}

	proxy.ServeHTTP(w, r)