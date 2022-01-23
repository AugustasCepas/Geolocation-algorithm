# Geolocation task
 
# Admin

## To install & run:
``` 
Place database.csv file into the /src folder

Commands for the terminal:
1) cd geolocation_task/src
2) go run main.go (this creates database.parquet file)
3) go build main.go
4) py geolocation_test.py --executable main.exe --database database.parquet

```


# Performance figures

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

# Things to mention

Praquet file takes less space and should take less time to load (2-3x less) https://medium.com/@u.praneel.nihar/improving-read-write-store-performance-by-changing-file-formats-serialization-protocols-bfdb13114004

However, it needs a header and to read it you need a structure, which needs a lot of memory.

I have tried reading a line and appending start, end, code and city splices, this way it takes less memory, but takes 60seconds to read a db.