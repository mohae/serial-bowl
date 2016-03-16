# serial-bowl

Some Go serialization benchmarks using various, simple data structures.

* [Cap'n Proto2](https://github.com/zombiezen/go-capnproto2)
* encoding/json
* [Flatbuffers](https://github.com/google/flatbuffers/go)
* [Protocol Buffers v3](https://github.com/golang/protobuf/proto)

Encode and decode benchmarks are performed with each listed protocol using structs with randomly generated data.  Each protocol uses the same data for a given run.  The data is randomly generated prior to running any of the benchmarks.

## Installing and running

Assuming Go is installed and $GOPATH is set:

    $ go get github.com/mohae/serial-bowl
    $ serial-bowl

### Supported flags

Flag|Default|Desc  
:--|:--|:--  
output|stdout|if not stdout: destination file for output  
o|stdout|if not stdout: destination file for output (short)  
format|txt|format of output: txt, csv, md  
f|txt|format of output: txt, csv, md (short)

## Benchmarks
The benchmarks test a few different situations, none of which are particularly complex data structures.  Each data structure has a slice of 1000 elements randomly generated.  These are used for each serialization protocol so that each are working with the exact same data for a given benchmark run.

These results were obtained with a VirtualBox guest running Debian Jessie with 1 CPU and 4GB of RAM.  The host is a Windows 7 professional system  with  an Intel i5-3570K @ 3.40 GHz CPU and 32GB of DDR3 1336MHz RAM using 9-9-9-24-2 timing.

### Data structures
#### BasicMemInfo
BasicMemInfo is a datastructure consisting of a subset of the keys in `/proc/meminfo`.

Protocol|Operation|Data Type|Operations|Ns/Op|Bytes/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|Marshal|BasicMemInfo|2000000|823|260|6  
CapnProto2|Unmarshal|BasicMemInfo|3000000|527|224|5  
Flatbuffers|Serialize|BasicMemInfo|3000000|400|0|0  
Flatbuffers|Deserialize|BasicMemInfo|20000000|61|32|1  
JSON|Marshal|BasicMemInfo|500000|2604|712|5  
JSON|Unmarshal|BasicMemInfo|200000|7091|503|10  
ProtobufV3|Marshal|BasicMemInfo|2000000|995|456|6  
ProtobufV3|Unmarshal|BasicMemInfo|2000000|626|208|1  

#### MemInfo
MemInfo is a datastructure consisting of a number of ints, which are defined as int64 for Flatbuffers and int for Go data structures.  MemInfo's fields are all the keys in `/proc/meminfo`.

Protocol|Operation|Data Type|Operations|Ns/Op|Bytes/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|Marshal|MemInfo|1000000|1673|532|6  
CapnProto2|Unmarshal|MemInfo|3000000|548|224|5  
Flatbuffers|Serialize|MemInfo|1000000|1679|0|0  
Flatbuffers|Deserialize|MemInfo|20000000|61|32|1  
JSON|Marshal|MemInfo|100000|11003|5065|8  
JSON|Unmarshal|MemInfo|30000|46653|1527|42  
ProtobufV3|Marshal|MemInfo|500000|2903|1224|8  
ProtobufV3|Unmarshal|MemInfo|1000000|2069|208|1  

#### Message
Message is a basic datastructure that has a `Data` field that can be a varying number of bytes.  This structure is tested using data sizes of, in bytes: 16, 256, 1024, 4096.  The 4096 byte data test is not done with Cap'n Proto as it results in an error: capnp: NewMessage called on arena with data.  I will enable this benchmark when I get a better understanding of Cap'n Proto.  

Protocol|Operation|Data Type|Operations|Ns/Op|Bytes/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|Marshal|Message 16B|1000000|1493|404|12  
CapnProto2|Unmarshal|Message 16B|3000000|571|224|5  
Flatbuffers|Serialize|Message 16B|5000000|260|0|0  
Flatbuffers|Deserialize|Message 16B|20000000|66|32|1  
JSON|Marshal|Message 16B|1000000|1852|440|6  
JSON|Unmarshal|Message 16B|300000|4044|319|7  
ProtobufV3|Marshal|Message 16B|2000000|775|328|5  
ProtobufV3|Unmarshal|Message 16B|3000000|591|232|3  
CapnProto2|Marshal|Message 256B|1000000|1984|660|12  
CapnProto2|Unmarshal|Message 256B|3000000|600|224|5  
Flatbuffers|Serialize|Message 256B|5000000|274|0|0  
Flatbuffers|Deserialize|Message 256B|20000000|69|32|1  
JSON|Marshal|Message 256B|500000|3398|1400|7  
JSON|Unmarshal|Message 256B|100000|10650|575|7  
ProtobufV3|Marshal|Message 256B|2000000|976|552|5  
ProtobufV3|Unmarshal|Message 256B|2000000|735|472|3  
CapnProto2|Marshal|Message 1024B|500000|2963|1492|12  
CapnProto2|Unmarshal|Message 1024B|3000000|564|224|5  
Flatbuffers|Serialize|Message 1024B|5000000|292|0|0  
Flatbuffers|Deserialize|Message 1024B|20000000|73|32|1  
JSON|Marshal|Message 1024B|300000|5929|6040|8  
JSON|Unmarshal|Message 1024B|50000|30181|1439|7  
ProtobufV3|Marshal|Message 1024B|1000000|1374|1416|5  
ProtobufV3|Unmarshal|Message 1024B|2000000|1089|1240|3  
Flatbuffers|Serialize|Message 4096B|3000000|445|0|0  
Flatbuffers|Deserialize|Message 4096B|20000000|60|32|1  
JSON|Marshal|Message 4096B|100000|13329|15512|9  
JSON|Unmarshal|Message 4096B|10000|108039|4895|7  
ProtobufV3|Marshal|Message 4096B|500000|2542|4872|5  
ProtobufV3|Unmarshal|Message 4096B|1000000|2228|4312|3  

#### RedditAccount
RedditAccount is the data structure for a Reddit Account.  It consistes of a `thing` whose object is `account` per https://github.com/reddit/reddit/wiki/JSON.

Protocol|Operation|Data Type|Operations|Ns/Op|Bytes/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|Marshal|RedditAccount|500000|3571|1045|27  
CapnProto2|Unmarshal|RedditAccount|3000000|515|224|5  
Flatbuffers|Serialize|RedditAccount|2000000|896|0|0  
Flatbuffers|Deserialize|RedditAccount|20000000|61|32|1  
JSON|Marshal|RedditAccount|300000|4267|1358|6  
JSON|Unmarshal|RedditAccount|100000|11398|565|11  
ProtobufV3|Marshal|RedditAccount|1000000|1442|696|6  
ProtobufV3|Unmarshal|RedditAccount|1000000|1341|507|7  

## Notes
There is probably room for improvement, esp. with Cap'n Proto.

### Flatbuffers
The Flatbuffers benches reuses the `builder` object; doing a `Reset()` instead of creating a new `builder`.  Creating a new `builder` via `flatbuffers.NewBuilder(0)` results in a significant degradation of performance and should be avoided when serializing data with Flatbuffers.
