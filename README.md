# go-fips

Collect FIPS mode information about your runtime.

Whether the kernel was booted with FIPS mode parameters. And whether the
openssl linked can properly enable FIPS mode.


# Docs

https://godoc.org/github.com/vbatts/go-fips

# Setup

See http://www.openssl.org/docs/fips/UserGuide-2.0.pdf
to set up an environment where fips mode can be enabled

# Building

```
go build .
```

and

```
go build -tags fips .
```

# Testing

```
go test .
```

and

```
go test -tags fips .
```

