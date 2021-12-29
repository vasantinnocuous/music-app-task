# Steps to compile application
1. Switch to the directory containing source code
```
    go build -o listMaker .
```


# Steps to run compiled binary
1. Run following command to invoke feature one, following command will output data on standard console
```
    ./listMaker -filepath path_to_some_file.xml -minTrackCount 11 -releasedBefore 2001.01.01
```

2. To invoke help
```
    ./listMaker --help
```

# Filters Available

* minTrackCount int
    	Minimum number of tracks
* releasedBefore string
    	Date to before the records are asked for (default "2099.12.31")


# Working Process
1. Application generates Filters based on command line arguments provided
2. Application then applies those filters on every record
3. Application finally outputs the filtered data
