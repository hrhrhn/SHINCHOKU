(View With `command + shift + V`)

http/1.1 Server (No Certification)
```
(base) htakagi@Mac ~ % curl -I --http2 --insecure https://localhost:8080/
curl: (35) error:1408F10B:SSL routines:ssl3_get_record:wrong version number
```

Running http/2 Server
```
(base) htakagi@Mac ~ % curl -I --http2 --insecure https://localhost:8080/
HTTP/1.1 200 OK
Date: Wed, 23 Mar 2022 06:08:37 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8
```