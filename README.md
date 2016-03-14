# serial-bowl

Some Go serialization benchmarks using various, simple datastructures.

* flatbuffers
* encoding/json

Encode and decode benchmarks are performed with each listed protocol using structs with randomly generated data.  Each protocol uses the same data for a given run.  The data is randomly generated prior to running any of the benchmarks.

## Benchmarks
The benchmarks test a few different situations, none of which are particularly complex data structures.  Each data structure has a slice of 1000 elements randomly generated.  These are used for each serialization protocol so that each are working with the exact same data for a given benchmark run.

These results were obtained with a VirtualBox guest running Debian Jessie with 1 CPU and 4GB of RAM.  The host is a Windows 7 professional system  with  an Intel i5-3570K @ 3.40 GHz CPU and 32GB of DDR3 1336MHz RAM using 9-9-9-24-2 timing.

### RedditAccount
RedditAccount is the data structure for a Reddit Account.  It consistes of a `thing` whose object is `account` per https://github.com/reddit/reddit/wiki/JSON.

##### Serialize/Marshal  
    RedditAccountSerializeFB:           2000000 ops      878 ns/Op        0 bytes/Op     0 allocs/Op  
    RedditAccountJSONMarshal:            300000 ops     4199 ns/Op     1363 bytes/Op     6 allocs/Op  


##### Deserialize/Unmarshal  
    RedditAccountDeserializeFB:        30000000 ops       59 ns/Op       32 bytes/Op     1 allocs/Op  
    RedditAccountJSONUnmarshal:          100000 ops    11685 ns/Op      565 bytes/Op    11 allocs/Op  

### BasicMemInfo
BasicMemInfo is a datastructure consisting of a subset of the keys in `/proc/meminfo`.

##### Serialize/Marshal  
    BasicMemInfoSerializeFB:            5000000 ops      383 ns/Op        0 bytes/Op     0 allocs/Op  
    BasicMemInfoJSONMarshal:            1000000 ops     2438 ns/Op      712 bytes/Op     5 allocs/Op  

##### Deserialize/Unmarshal  
    BasicMemInfoDeserializeFB:         30000000 ops       59 ns/Op       32 bytes/Op     1 allocs/Op  
    BasicMemInfoJSONUnmarshal:           200000 ops     6783 ns/Op      503 bytes/Op    10 allocs/Op  

### MemInfo
MemInfo is a datastructure consisting of a number of ints, which are defined as int64 for Flatbuffers and int for Go data structures.  MemInfo's fields are all the keys in `/proc/meminfo`.

##### Serialize/Marshal  
    MemInfoSerializeFB:                 1000000 ops     1616 ns/Op        0 bytes/Op     0 allocs/Op  
    MemInfoJSONMarshal:                  100000 ops    10022 ns/Op     5448 bytes/Op     8 allocs/Op  

##### Deserialize/Unmarshal  
    MemInfoDeserializeFB:              30000000 ops       57 ns/Op       32 bytes/Op     1 allocs/Op  
    MemInfoJSONUnmarshal:                 30000 ops    40766 ns/Op     1527 bytes/Op    42 allocs/Op  

### Message
Message is a basic datastructure that has a `Data` field that can be a varying number of bytes.  This structure is tested using data sizes of, in bytes: 256, 1024, 4096.

##### Serialize/Marshal: Data: 256 bytes  
    MessageSerializeFB:                 5000000 ops      262 ns/Op        0 bytes/Op     0 allocs/Op  
    MessageJSONMarshal:                  500000 ops     2884 ns/Op     1400 bytes/Op     7 allocs/Op  

##### Deserialize/Unmarshal Data: 256 bytes  
    MessageDeserializeFB:              20000000 ops       61 ns/Op       32 bytes/Op     1 allocs/Op  
    MessageJSONUnmarshal:                100000 ops    10664 ns/Op      575 bytes/Op     7 allocs/Op  

##### Serialize/Marshal: Data: 1024 bytes  
    MessageSerializeFB:                 5000000 ops      286 ns/Op        0 bytes/Op     0 allocs/Op  
    MessageJSONMarshal:                  300000 ops     5844 ns/Op     6040 bytes/Op     8 allocs/Op  

##### Deserialize/Unmarshal Data: 1024 bytes  
    MessageDeserializeFB:              30000000 ops       59 ns/Op       32 bytes/Op     1 allocs/Op  
    MessageJSONUnmarshal:                 50000 ops    31464 ns/Op     1439 bytes/Op     7 allocs/Op  

##### Serialize/Marshal: Data: 4096 bytes  
    MessageSerializeFB:                 3000000 ops      446 ns/Op        0 bytes/Op     0 allocs/Op  
    MessageJSONMarshal:                  100000 ops    13456 ns/Op    15512 bytes/Op     9 allocs/Op  


##### Deserialize/Unmarshal Data: 4096 bytes  
    MessageDeserializeFB:              30000000 ops       53 ns/Op       32 bytes/Op     1 allocs/Op  
    MessageJSONUnmarshal:                 10000 ops   113275 ns/Op     4895 bytes/Op     7 allocs/Op  

## Notes
Currently, the data generator assumes that the benchmark is being run on a 64bit system.  That is, the max value of an int is defined to be `1<<63-1`.

### Flatbuffers
The Flatbuffers benches reuses the `builder` object; doing a `Reset()` instead of creating a new `builder`.  Creating a new `builder` via `flatbuffers.NewBuilder(0)` results in a significant degradation of performance and should be avoided when serializing data with Flatbuffers.

#### RedditAccount w `builder := flatbuffers.NewBuilder(0)`
