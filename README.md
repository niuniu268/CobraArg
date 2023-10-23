# Test Golang Libraries
## os.arg
## flag
## cobra

## Section 1 
``` go build -o "arg_demo"  ```
``` ./arg_demo a b c d ```

- Results:
 
``` 
    arg[0] is ./arg_demo 
    arg[1] is a 
    arg[2] is b 
    arg[3] is c 
    arg[4] is d 

```

## Section 2

``` go build -o "arg_demo" ```
``` ./arg_demo -help ```

- Results: 

``` 
Usage of ./arg_demo:
  -age int
        年龄 (default 18)
  -d duration
        延迟的时间间隔
  -married
        婚否
  -name string
        姓名 (default "张三")

```

``` ./arg_demo -name 沙河娜扎 --age 28 -ma-marrried=false -d=1h30m ```

- Results:
``` 
沙河娜扎 28 false 1h30m0s
[]
0
4

```

## Section 3

``` go build -o "arg_demo" ```
``` ./arg_demo -ip ip -address 127.0.0.1  ```

- Results: 
``` 
{ip: 127.0.0.1, loopback: true}

```