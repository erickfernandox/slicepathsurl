[![made-with-Go](https://img.shields.io/badge/made%20with-Go-brightgreen.svg)](http://golang.org)
<h1 align="center">SlicePathURL</h1> <br>

<p align="center">
  <a href="#--usage--explanation">Usage</a> •
  <a href="#--installation--requirements">Installation</a>
</p>

## - Installation & Requirements:
```bash
> go install github.com/erickfernandox/slicepathurl@latest
```
OR
```bash
> git clone https://github.com/erickfernandox/slicepathurl.git
> cd slicepathurl
> go build slicepathurl.go
> chmod +x slicepathurl
> ./slicepathurl -h
```
<br>

## - Why use SlicePathURL?

Examples:

Sometimes, Nuclei may fail to identify a vulnerability in the root domain, for example, in https://example.com/. However, it is possible that vulnerabilities may exist in paths beyond the root domain, such as in https://example.com/path_one/. 

Below is a real example that was found:

```bash
echo "https://subdomain.example.com/"|nuclei -tags rce

[INF] No results found. Better luck next time!
```

```bash
echo "https://subdomain.example.com/extranet/"|nuclei -tags rce

[2023-01-01 23:54:42] [CVE-2017-5638] [http] [critical] https://subdomain.example.com/extranet/
```


## - How does SlicePathURL work?


