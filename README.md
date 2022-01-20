# Geolocation task
 
# Admin

## To install & run:
``` 
1) cd geolocation_task/src
2) go run main.go
```

## Things to mention:
```
1) Before launching, database.csv file must be placed to the /src folder

2) Currently, there must be space after the input, for example "LOAD ", "LOOKUP 1.0.0.0 ".

3) Creating an executable and running it with python script, might now work atm.

```

## Other useful commands:

# Python:
py geolocation_test.py --executable sample_app.exe --database database.csv
py geolocation_test.py --executable main.exe --database database.csv

# CPP
g++ sample_app.cpp -o sample_app
./sample_app database.csv

# Golang
go build main.go