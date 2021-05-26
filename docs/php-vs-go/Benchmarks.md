# Benchmarking PHP vs Go!

Inside this folder there are various projects that contain code where you can compare raw performance of webservers using Go or PHP.


- test-apache
    - Contains simple Hello World example for an apache server running php code.
- test-go-simple
    - Contains simple Hello World example for a golang based server using fasthttp.


## Requirements
To execute these tests you need:
- Docker
- go >= 1.16
- wrk

## How to run

First start by running the Apache example. Inside the `test-apache` folder execute.

#### Build & run image

```
$ docker build -t test-apache .
$ docker run -p 8091:80 test-apache
```

You should have your apache server running on port 8091. Go to http://localhost:8091 and you should see "Hello PHP!".

#### Perform Test

```
$ wrk -t12 -c400 -d30s http://localhost:8091
```

**Result:**
```
Running 30s test @ http://localhost:8091/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   296.92ms  447.34ms   2.00s    80.79%
    Req/Sec     2.00k     2.37k   26.90k    94.48%
  591798 requests in 30.09s, 103.39MB read
  Socket errors: connect 0, read 0, write 0, timeout 43
Requests/sec:  19665.04
Transfer/sec:      3.44MB
```

#### Go test

Make sure to shutdown the apache server first. And then run the benchmark for the go server.

```
$ docker build -t test-go-simple .
$ docker run -p 8899:8899 test-go-simple
```

```
$ wrk -t12 -c400 -d30s http://localhost:8899
```

**Result:**
```
Running 30s test @ http://localhost:8899/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.19ms    7.68ms 223.94ms   97.74%
    Req/Sec     6.96k     1.06k   15.12k    74.35%
  2487425 requests in 30.08s, 339.22MB read
Requests/sec:  82680.83
Transfer/sec:     11.28MB
```


## Results

For these simple examples you can see that go beats PHP for more than 4 times (counting rq/s). Actually 4.2 times more handled requests using go.
