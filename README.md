# Geolocation task
 
# Admin

## To install & run:
``` 
1) cd geolocation_task/src
2) go build main.go
3) py geolocation_test.py --executable main.exe --database database.csv

```

## Things to mention:
```
1) Before launching, database.csv file must be placed to the /src folder

```

## Other useful commands:

# Python:
py geolocation_test.py --executable sample_app.exe --database database.csv
py geolocation_test.py --executable main.exe --database database.csv

# CPP
g++ sample_app.cpp -o sample_app
./sample_app database.csv

# Performance figures

Database loaded Memory usage: 722.71mb Load time: 1s
OK    1.0.0.0 US Los Angeles Memory usage: 722.72mb Lookup time: 0ns
OK    71.6.28.0 US San Jose Memory usage: 722.72mb Lookup time: 3ms
OK    71.6.28.255 US San Jose Memory usage: 722.73mb Lookup time: 2ms
OK    71.6.29.0 US Concord Memory usage: 722.73mb Lookup time: 2ms
OK    53.103.144.0 DE Stuttgart Memory usage: 722.77mb Lookup time: 2ms
OK    53.255.255.255 DE Stuttgart Memory usage: 722.77mb Lookup time: 1ms
OK    54.0.0.0 US Rahway Memory usage: 722.79mb Lookup time: 999μs
OK    223.255.255.255 AU Brisbane Memory usage: 722.8mb Lookup time: 10ms
OK    5.44.16.0 GB Hastings Memory usage: 722.8mb Lookup time: 1ms
OK    8.24.99.0 US Hastings Memory usage: 722.8mb Lookup time: 0ns