# testcs
Test for CS


## Build
I included a bash script that builds and runs the application.  You shoud run this first.  The script is named "tserver.sh".


## Server
The server is started via the "tserver.sh" script.  It runs in port 8282.
To run, go into the testcs directory and run the script:  ./tserver.sh


## Client
The client runs from the "tclient.sh" script. It delivers the files to the server via curl commands. 
Run the client after running the server.  To run, go into the testcs directory and run the script:   ./tclient.sh
The client script looks for data under ./testcs/data.  Can be changed to look from a different location.


## Result
After running the client, you should see a result similar to this:

╚══[$]> ./tserver.sh
>>  create bin directory
>>  build the test server


>>  execute the server

CS SERVICE START
Starting injest dispatcher


CSV REPORT
Average age: 34
Average person: Mikayla SERRANO

Median age: 31
Median person: Tyler BLACKWELL

Elapsed time: 0.018474389 seconds
Number of routines: 6


# Considerations
I considered using external libraries, but as the instructions indicated, I used the basic libaries.
I usually will use gorilla/mux and toml as a configuration.
I added some test cases as examples.  But I didn't included test cases for all the code.


# Assumptions
That as soon as the files are read, the Report prints.  I used an interface{} channel to deal with that.
That the files will always have a header row.


# 10M records
Break each file into multiparts that are queued.  Assuming that data is processed together.


# 20K URLs
Provide multiple ports that handle a configurable amount of URLs connections

