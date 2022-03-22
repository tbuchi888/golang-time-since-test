# golang-time-since-test

## Test run

`go run time-since-test.go`


## Build
`go build time-since-test.go`


## Execute
linux & mac
`./time-since-test`

windows
`.\time-since-test.exe`

## Example of Execution Result
linux & mac

```
$ ./time-since-test 
t0: 2022-03-22 17:31:17.117956 +0900 JST m=+0.000063189
t1: 2022-03-22 17:31:17.117956 +0900 JST m=+0.000063338
time.Since(t0): 340ns
2022-03-22 17:31:17.118413 +0900 JST m=+0.000520403
2022-03-22 17:31:17.118416 +0900 JST m=+0.000523500
2022-03-22 17:31:17.118419 +0900 JST m=+0.000526186
2022-03-22 17:31:17.118422 +0900 JST m=+0.000528644
2022-03-22 17:31:17.118424 +0900 JST m=+0.000531216
2022-03-22 17:31:17.118427 +0900 JST m=+0.000533754
2022-03-22 17:31:17.11843 +0900 JST m=+0.000536580
2022-03-22 17:31:17.118432 +0900 JST m=+0.000539352
2022-03-22 17:31:17.118435 +0900 JST m=+0.000541871
2022-03-22 17:31:17.118437 +0900 JST m=+0.000544331
$
```

windows
```
PS C:\Users\test001\Desktop> .\time-since-test.exe
t0: 2022-03-22 08:13:39.7611383 +0000 GMT m=+0.002171201
t1: 2022-03-22 08:13:39.7611383 +0000 GMT m=+0.002171201
time.Since(t0): 0s
Execute time.Now() and fmt.Println 10 times with for statement
2022-03-22 08:13:39.7979066 +0000 GMT m=+0.038939001
2022-03-22 08:13:39.7979066 +0000 GMT m=+0.038939001
2022-03-22 08:13:39.7979066 +0000 GMT m=+0.038939001
2022-03-22 08:13:39.7979066 +0000 GMT m=+0.038939001
2022-03-22 08:13:39.7984478 +0000 GMT m=+0.039480201
2022-03-22 08:13:39.7989824 +0000 GMT m=+0.040014801
2022-03-22 08:13:39.7996092 +0000 GMT m=+0.040641601
2022-03-22 08:13:39.8003185 +0000 GMT m=+0.041350901
2022-03-22 08:13:39.8009501 +0000 GMT m=+0.041982501
2022-03-22 08:13:39.8016367 +0000 GMT m=+0.042669101
PS C:\Users\test001\Desktop>
```