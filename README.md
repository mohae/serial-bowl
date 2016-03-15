# serial-bowl

Some Go serialization benchmarks using various, simple datastructures.

* flatbuffers
* encoding/json

Encode and decode benchmarks are performed with each listed protocol using structs with randomly generated data.  Each protocol uses the same data for a given run.  The data is randomly generated prior to running any of the benchmarks.

## Benchmarks
The benchmarks test a few different situations, none of which are particularly complex data structures.  Each data structure has a slice of 1000 elements randomly generated.  These are used for each serialization protocol so that each are working with the exact same data for a given benchmark run.

These results were obtained with a VirtualBox guest running Debian Jessie with 1 CPU and 4GB of RAM.  The host is a Windows 7 professional system  with  an Intel i5-3570K @ 3.40 GHz CPU and 32GB of DDR3 1336MHz RAM using 9-9-9-24-2 timing.

### Data structures
#### BasicMemInfo
BasicMemInfo is a datastructure consisting of a subset of the keys in `/proc/meminfo`.

#### MemInfo
MemInfo is a datastructure consisting of a number of ints, which are defined as int64 for Flatbuffers and int for Go data structures.  MemInfo's fields are all the keys in `/proc/meminfo`.

#### Message
Message is a basic datastructure that has a `Data` field that can be a varying number of bytes.  This structure is tested using data sizes of, in bytes: 16, 256, 1024, 4096.

#### RedditAccount
RedditAccount is the data structure for a Reddit Account.  It consistes of a `thing` whose object is `account` per https://github.com/reddit/reddit/wiki/JSON.

### Results
Protocol|Operation|Data Type|Operations|Ns/Op|Bytes/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
Flatbuffers|Serialize|BasicMemInfo|3000000|399|0|0  
Flatbuffers|Deserialize|BasicMemInfo|20000000|61|32|1  
JSON|Marshal|BasicMemInfo|500000|2599|712|5  
JSON|Unmarshal|BasicMemInfo|200000|6959|503|10  
ProtobufV3|Marshal|BasicMemInfo|2000000|1000|456|6  
ProtobufV3|Unmarshal|BasicMemInfo|3000000|614|208|1  
Flatbuffers|Serialize|MemInfo|1000000|1635|0|0  
Flatbuffers|Deserialize|MemInfo|20000000|61|32|1  
JSON|Marshal|MemInfo|100000|11161|5071|8  
JSON|Unmarshal|MemInfo|30000|44418|1527|42  
ProtobufV3|Marshal|MemInfo|500000|2953|1224|8  
ProtobufV3|Unmarshal|MemInfo|1000000|1952|208|1  
Flatbuffers|Serialize|Message 16B|5000000|261|0|0  
Flatbuffers|Deserialize|Message 16B|20000000|64|32|1  
JSON|Marshal|Message 16B|1000000|1773|440|6  
JSON|Unmarshal|Message 16B|300000|4041|319|7  
ProtobufV3|Marshal|Message 16B|2000000|768|328|5  
ProtobufV3|Unmarshal|Message 16B|3000000|596|232|3  
Flatbuffers|Serialize|Message 256B|5000000|275|0|0  
Flatbuffers|Deserialize|Message 256B|20000000|65|32|1  
JSON|Marshal|Message 256B|500000|3168|1400|7  
JSON|Unmarshal|Message 256B|100000|10947|575|7  
ProtobufV3|Marshal|Message 256B|2000000|926|552|5  
ProtobufV3|Unmarshal|Message 256B|2000000|716|472|3  
Flatbuffers|Serialize|Message 1024B|5000000|291|0|0  
Flatbuffers|Deserialize|Message 1024B|20000000|72|32|1  
JSON|Marshal|Message 1024B|200000|6404|6040|8  
JSON|Unmarshal|Message 1024B|50000|32652|1439|7  
ProtobufV3|Marshal|Message 1024B|1000000|1392|1416|5  
ProtobufV3|Unmarshal|Message 1024B|2000000|1088|1240|3  
Flatbuffers|Serialize|Message 4096B|3000000|445|0|0  
Flatbuffers|Deserialize|Message 4096B|20000000|60|32|1  
JSON|Marshal|Message 4096B|100000|13522|15512|9  
JSON|Unmarshal|Message 4096B|10000|116834|4895|7  
ProtobufV3|Marshal|Message 4096B|500000|2519|4872|5  
ProtobufV3|Unmarshal|Message 4096B|1000000|2254|4312|3  
Flatbuffers|Serialize|RedditAccount|2000000|902|0|0  
Flatbuffers|Deserialize|RedditAccount|20000000|60|32|1  
JSON|Marshal|RedditAccount|300000|4220|1359|6  
JSON|Unmarshal|RedditAccount|100000|11374|566|11  
ProtobufV3|Marshal|RedditAccount|1000000|1481|696|6  
ProtobufV3|Unmarshal|RedditAccount|1000000|1297|508|7  

## Notes
Currently, the data generator assumes that the benchmark is being run on a 64bit system.  That is, the max value of an int is defined to be `1<<63-1`.

### Flatbuffers
The Flatbuffers benches reuses the `builder` object; doing a `Reset()` instead of creating a new `builder`.  Creating a new `builder` via `flatbuffers.NewBuilder(0)` results in a significant degradation of performance and should be avoided when serializing data with Flatbuffers.

#### RedditAccount w `builder := flatbuffers.NewBuilder(0)`
