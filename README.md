config.yaml contains configuration data.

For each http request inserting an entry into the slice with the actual Unix Timestamp.

TimeStamp slice is updated such that the all the values are just 60 sec older than latest http request.

On server exit, saving the state of the slice to a json file.


Added RATE LIMIT 15 Requests per 20 Seconds,
The count per last minute is kept and available in file.




