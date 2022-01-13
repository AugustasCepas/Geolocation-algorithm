# Task for Great Low Level Developer

## Prerequisites
You should have received these files:
- Geolocation.md
- geolocation/database.csv
- geolocation/sample_app.cpp
- geolocation/geolocation_test.py
- geolocation/CMakeLists.txt

## Specification
Using systems programming language of your choice (C/C++/rust/go/zig...) build
a console app that performs geolocation lookups for given IP addresses.

To do the lookup, use the supplied geolocation information from .csv file,
(later on referred to as "database" file).

### The Protocol
Your application is expected to implement specific line-based request-response
protocol for communication with other processes.

The protocol itself consists only of three commands:
* LOAD
* LOOKUP <IPv4 address>
* EXIT

#### LOAD
LOAD command is used for loading the database into memory. The application
should respond with the string `OK` once the database is fully loaded. Note, as
this is line-based protocol, the string `OK` must end with new line.

Example command execution:
```
> LOAD
< OK

```

Here `>` means data that your app receives via stdin, and `<` means data your
app writes into the stdout.

#### LOOKUP
LOOKUP command is used for performing geolocation lookup. The application
should respond with the location formated as `<COUNTRY CODE>,<CITY>`, there
should be no spaces before or after the comma. The response must end with
new line.

Example command execution:
```
> LOOKUP 71.8.28.3
< US,San Jose
```

Here `>` means data that your app receives via stdin, and `<` means data your
app writes into the stdout.

#### EXIT
EXIT command indicates that the application should now exit. Note that before
exiting, the application should respond with OK. The response must end with
new line.

Example command execution:
```
> EXIT
< OK
```

Here `>` means data that your app receives via stdin, and `<` means data your
app writes into the stdout.

---

You can use the provided `geolocation/geolocation_test.py` script to test the
compatability of your implementation.

### The Goal
Your main goal is to make this geolocation app as efficient as possible. Any
and all optimizations are on the table as long as it can help lower the
resource usage.

More specifically by "resource usage" we mean
* load time
* lookup time
* memory usage
in roughly these proportions:

```
points = load_time_ms + memory_usage_mb * 10 + lookup_time_ms * 1000
```

The goal is to collect as low amount of points as possible.

In order to achieve the best performance, you are encouraged to preprocess the
database file into your own, custom, on-disk format, which will be used during
database loading.

The same goes for in-memory data structures. You are encouraged to research and
implement any structures/algorithms which will help you lower the points.

Parallel processing can also be used, if you believe that it will help you
achieve better results, you can assume, that the target machine has 4 CPU
cores.

Choosing a good approach for this task is the most important choice, so do not be
afraid to use your creativity and _destroy_ the numbers!

### Requirements
- Provide README, explaining how to compile and setup the program
- Specify operating system it is designed for
- Document your approach at solving the problem
- If you choose to use preprocessed database format, provide a script or some
  program (with source code) which was used to do preprocessing. And of course
  describe your database format
- Provide expected performance figures (in case we would get lower speeds
  due to some minor mistake, we will contact you, to resolve the issue)
- You can use any 3rd party libraries/frameworks (although you will need to
  explain why it's needed, and how it works internally)

---

Good luck!
