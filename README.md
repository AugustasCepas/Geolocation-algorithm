# Geolocation algorithm
 
# Admin

## To install & run:
``` 
Place database.csv file into the /src folder

Commands for the terminal:
1) cd geolocation_task/src
2) go run main.go (this creates shorterdb.csv file)
3) go build main.go
4) Run project with shorter (contains only necessary fields) or regular db file
    a) py geolocation_test.py --executable main.exe --database shorterdb.csv
    b) py geolocation_test.py --executable main.exe --database database.csv

```

# Performance figures
```
Database loaded Memory usage: 474.06mb Load time: 1s
OK    1.0.0.0 US Los Angeles Memory usage: 474.06mb Lookup time: 0ns
OK    71.6.28.0 US San Jose Memory usage: 474.07mb Lookup time: 0ns
OK    71.6.28.255 US San Jose Memory usage: 474.07mb Lookup time: 0ns
OK    71.6.29.0 US Concord Memory usage: 474.09mb Lookup time: 0ns
OK    53.103.144.0 DE Stuttgart Memory usage: 474.09mb Lookup time: 0ns
OK    53.255.255.255 DE Stuttgart Memory usage: 474.09mb Lookup time: 0ns
OK    54.0.0.0 US Rahway Memory usage: 474.09mb Lookup time: 0ns
OK    223.255.255.255 AU Brisbane Memory usage: 474.09mb Lookup time: 0ns
OK    5.44.16.0 GB Hastings Memory usage: 474.09mb Lookup time: 0ns
OK    8.24.99.0 US Hastings Memory usage: 474.09mb Lookup time: 0ns

Load time points: 1015 pts
Memory usage points: 4741 pts
Lookup time points: 0 pts
Total points: 5756 pts
```
