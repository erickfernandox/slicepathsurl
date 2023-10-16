[![made-with-Go](https://img.shields.io/badge/made%20with-Go-brightgreen.svg)](http://golang.org)
<h2 align="center">SlicePathsURL</h2> <br>

<p align="center">
  <a href="#--usage--explanation">Usage</a> •
  <a href="#--installation--requirements">Installation</a> •
  <a href="#--why-use-slicepathsurl">Why use SlicePathsURL?</a> •
  <a href="#--how-does-slicepathsurl-work">How does SlicePathsURL work?</a>
</p>

<h3 align="center">SlicePathsURL slices a URL into directory levels to complement tools like Nuclei in searching for vulnerabilities in directories beyond the root of the URL.</h3>


## - Installation & Requirements:

```bash
go install github.com/erickfernandox/slicepathsurl@latest
```
OR
```bash
git clone https://github.com/erickfernandox/slicepathsurl.git
cd slicepathsurl
go build slicepathsurl.go
chmod +x slicepathsurl
./slicepathsurl -h
```
<br>

## - Why use SlicePathsURL?

Examples:

Sometimes, Nuclei may fail to identify a vulnerability in the root domain, for example, in https://example.com/. However, it is possible that vulnerabilities may exist in paths beyond the root domain, such as in https://example.com/path_one/. 

Below is a real example that was found:

```bash
echo "https://subdomain.example.com/"|nuclei -tags rce

[INF] No results found. Better luck next time!
```

```bash
echo "https://subdomain.example.com/extranet/"|nuclei -tags rce

[2023-01-01 00:00:00] [CVE-2017-5638] [http] [critical] https://subdomain.example.com/extranet/
```

An RCE vulnerability, CVE-2017-5638, was discovered in Apache Struts in an application hosted at https://example.com/extranet/, but it was not found in the root directory of https://example.com/.

Below are additional examples where SlicePathURL was used to identify vulnerabilities that were not located in the root directory of the domain, but rather in a subdirectory:

```bash
[crlf-injection] [http] [medium] https://example.com/path_level2/%0d%0aSet-Cookie:crlfinjection=1; -> CRLF Injection
[open-redirect] [http] [low] https://subdomain.example.com/path_level2///interact.sh/%2F -> Open Redirect
[elmah-log-file] [http] [medium] https://xxx.example.com.br/perdiminhasenha/elmah.axd?AspxAutoDetectCookieSupport=1 -> Debug Information Exposed
[git-exposed] [http] [medium] https://xxx.example.com.br/path_level2/.git/config -> Git Exposed
[cache-poisoning] [http] [low] https://www.example.com/insights/?cb=poisoning [host.cache.interact.sh] - X-Forwarded-Host Cache Poisioning 
```

## - How does SlicePathsURL work?


```bash
echo "example.com"|gauplus > example_gauplus.txt

https://example.com/applications/data/user?id=123
https://example.com/applications/data/user?id=123&msg=error
https://example.com/applications/data/user/config?id=1
https://example.com/applications/data/config?test=tese
https://example.com/applications/data/config/info?data={}
https://example.com/applications/finder/search?q=123
https://example.com/applications/finder/search?q=123&order=desc

cat example_gauplus.txt|slicepathsurl -l 2

https://example.com/
https://example.com/applications

cat example_gauplus.txt|slicepathsurl -l 3

https://example.com/
https://example.com/applications
https://example.com/applications/data
https://example.com/applications/finder/

cat example_gauplus.txt|slicepathsurl -l 4

https://example.com/
https://example.com/applications
https://example.com/applications/data
https://example.com/applications/data/user
https://example.com/applications/data/config
https://example.com/applications/finder/
https://example.com/applications/finder/search

```

```bash
subfinder -d example.com | gauplus | slicepathsurl -l 2 > urls_all_paths_level2.txt
cat urs_all_paths_level2.txt | nuclei -tags crlf,rce,redirect
```

Identifying Git Exposed in 3 levels of URLs:
<br>The slicepathsurl tool takes a URL and divides it into 3 levels:</br>

```bash
https://example.com/
https://example.com/level2
https://example.com/level2/level3

```

Next, the URLs previously acquired via gauplus can be used in conjunction with httpx to extract the three-level hierarchy of the URLs and search for the .git file at every level of the URL. An example of this is shown below:

```bash
cat urls_all_paths_level2.txt | slicepathsurl -l 3 | httpx -path /.git/config -mr "refs/heads"

https://example.com/.git/config
https://example.com/level2/.git/config
https://example.com/level2/level3/.git/config

```

