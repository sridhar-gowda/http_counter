An IP Based Rate Limiter using sliding window approach.

config.yaml contains configuration data.

Time Complexity is O(n) number of hits in last minute(or window size). At any particular time the last windw data is available.

On server exit, saving the state of the slice to a json file.






