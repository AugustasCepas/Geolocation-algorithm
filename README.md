# Geolocation task
 
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

# Conclusion

Praquet on papre takes less space and should take less time to load (2-3x less) https://medium.com/@u.praneel.nihar/improving-read-write-store-performance-by-changing-file-formats-serialization-protocols-bfdb13114004

However, it needs a header and to read it you need a structure, which needs a lot of memory.

I have tried reading a line and appending start, end, code and city splices, this way it requires less memory, but takes 60seconds to read a db file.

Therefore, I got back to CSV file, but removed unecessary fields and achieved the least points as I could.

Lookup time varies, however, shorter CSV data file gives best result.

# Performance figures

Shorter CSV:
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

Original CSV:
Database loaded Memory usage: 653.86mb Load time: 1s
OK    1.0.0.0 US Los Angeles Memory usage: 653.86mb Lookup time: 0ns
OK    71.6.28.0 US San Jose Memory usage: 653.87mb Lookup time: 0ns
OK    71.6.28.255 US San Jose Memory usage: 653.87mb Lookup time: 0ns
OK    71.6.29.0 US Concord Memory usage: 653.9mb Lookup time: 0ns
OK    53.103.144.0 DE Stuttgart Memory usage: 653.9mb Lookup time: 0ns
OK    53.255.255.255 DE Stuttgart Memory usage: 653.9mb Lookup time: 0ns
OK    54.0.0.0 US Rahway Memory usage: 653.9mb Lookup time: 0ns
OK    223.255.255.255 AU Brisbane Memory usage: 653.9mb Lookup time: 0ns
OK    5.44.16.0 GB Hastings Memory usage: 653.9mb Lookup time: 0ns
OK    8.24.99.0 US Hastings Memory usage: 653.9mb Lookup time: 0ns

Load time points: 1573 pts
Memory usage points: 6539 pts
Lookup time points: 0 pts
Total points: 8112 pts

Parquet:
Database loaded Memory usage: 1.62gb Load time: 1s
OK    1.0.0.0 US Los Angeles Memory usage: 1.62gb Lookup time: 760μs
OK    71.6.28.0 US San Jose Memory usage: 1.62gb Lookup time: 0ns
OK    71.6.28.255 US San Jose Memory usage: 1.62gb Lookup time: 166μs
OK    71.6.29.0 US Concord Memory usage: 1.62gb Lookup time: 0ns
OK    53.103.144.0 DE Stuttgart Memory usage: 1.62gb Lookup time: 0ns
OK    53.255.255.255 DE Stuttgart Memory usage: 1.62gb Lookup time: 0ns
OK    54.0.0.0 US Rahway Memory usage: 1.62gb Lookup time: 0ns
OK    223.255.255.255 AU Brisbane Memory usage: 1.62gb Lookup time: 508μs
OK    5.44.16.0 GB Hastings Memory usage: 1.62gb Lookup time: 0ns
OK    8.24.99.0 US Hastings Memory usage: 1.62gb Lookup time: 0ns

Load time points: 1161 pts
Memory usage points: 16609 pts
Lookup time points: 1436 pts
Total points: 19206 pts