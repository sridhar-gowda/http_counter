**PROBLEM**
Using only the standard library, create a Go HTTP server that on each request responds with a counter of the total
number of requests that it has received during the previous 60 seconds (moving window). The server should continue to
the return the correct numbers after restarting it, by persisting data to a file. Implement IP Based Rate Limiter for
the same.

**SOLUTION**

Used sliding window technique. At any particular time, the last window timestamps data are being 
stored. Data older than window are being deleted.

Time Complexity is O(n) number of hits in last minute(window size). 
On server exit, saving the data to a json file.

config.yaml contains configurable data like Window size, Rate of Limit, Host address, Json file location to store on exit data.

**POINTS**
* Just clone the repository. Make sure Go has been installed.
* Run command - go run main.go
* Single(Sample) Unit Test is written for util function.







