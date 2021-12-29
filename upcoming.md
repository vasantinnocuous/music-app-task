# Notes for Upcoming tasks
1. The code base is quite abstract, the initial models are intutive to the given problem. Since the code that parses the records is separated with code that implements filteration of data, In case of any new data source, We just require to implement a custom XML **Unmarshler** interface on **Records** type just as we did for *Data* type to convert a String date in Golang's **time.Time** and would need to add some code to communicate, authenticate over that online serverice

2. To implement this feature, it just require to add a new filter
    e.g. There would be a new filter called **minCDCounts**, to implement this following changes would be required
    * *initCmdFlags* needs to add an integer flag for parsing command line argument
    * *getFilters* needs to define the filter respective to the value of newly added command line argument **minCDCounts**