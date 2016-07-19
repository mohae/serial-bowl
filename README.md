# serial-bowl

Some Go serialization benchmarks using various data structures.  I wrote this to see how the performance of various serialization formats changed based on data structure complexity and size.  Also, for formats with multiple packages implementing them, I wanted to see the relative performance between the implementations.

* [Cap'n Proto2](https://capnproto.org/)
* [Flatbuffers](https://google.github.io/flatbuffers/)
* [Gencode](https://github.com/andyleap/gencode)
* [Gob](https://golang.org/pkg/encoding/gob/)
* [JSON](https://www.json.org)
* [MessagePack](http://msgpack.org/index.html)
* [Protocol Buffers v3](https://developers.google.com/protocol-buffers/)

## Installing and running

Assuming Go is installed, $GOPATH is set, and $GOPATH/bin is part of your path:

    $ go install github.com/mohae/serial-bowl
    $ serial-bowl

### Supported flags

Flag|Default|Desc  
:--|:--|:--  
output|stdout|if not stdout: destination file for output  
o|stdout|if not stdout: destination file for output (short)  
format|txt|format of output: txt, csv, md  
f|txt|format of output: txt, csv, md (short)  
namesections|false|use group as section name: some restrictions apply  
n|false|use group as section name: some restrictions apply  
sections|false|don't separate groups of tests into sections  
s|false|don't separate groups of tests into sections  
sectionheader|false|if there are sections, add a section header row  
h|false|if there are sections, add a section header row  
sysinfo|false|add the system information to the output  
i|false|add the system information to the output  
datasetsize|1000|the size of the test dataset for each benchmark  
d|1000|the size of the test dataset for each benchmark  
benchformat||format(s) to benchmark  
b||format(s) to benchmark  

## Benchmarks
The benchmarks test a few different situations, none of which are particularly complex data structures.  Each data structure has a slice of 1000 elements randomly generated.  These are used for each serialization protocol so that each are working with the exact same data for a given benchmark run.

The observed results can vary between runs, with some operations varying by over 100%.  The decision on which serialization format to use should involve a number of factors that are specific to your use-case and requirements, of which, speed should be one.  The end decision shouldn't be made on benchmark results alone.

Information on the system that the benchmarks were run on:

```
Processors:  4
Model:       Intel(R) Core(TM) i5-3570K CPU @ 3.40GHz
Cache:       6144 KB
Memory:      32 GB
OS:          Ubuntu Ubuntu 16.04 LTS
Kernel:      4.4.0-21-generic
```
### BasicMemInfo
BasicMemInfo is a datastructure consisting of a subset of the keys in `/proc/meminfo`.

####BasicMemInfo  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|2000000|1161|228|3  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|1122|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|5000000|375|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|51|32|1  
GenCode|andyleap/gencode|Marshal|50000000|33|0|0  
GenCode|andyleap/gencode|Unmarshal|50000000|31|0|0  
Gob|encoding/gob|Marshal|2000000|1215|64|1  
Gob|encoding/gob|Unmarshal|2000000|879|23|1  
JSON|encoding/json|Marshal|500000|2796|687|5  
JSON|encoding/json|Unmarshal|200000|6639|503|10  
JSON|pquerna/ffjson|Marshal|500000|2543|1258|12  
JSON|pquerna/ffjson|Unmarshal|500000|3885|336|5  
JSON|pquerna/ffjson (buffer)|Marshal|1000000|1549|256|8  
JSON|pquerna/ffjson (buffer)|Unmarshal|500000|3535|339|5  
MessagePack|vmihailenco/msgpack|Marshal|1000000|2084|672|6  
MessagePack|vmihailenco/msgpack|Unmarshal|1000000|2562|240|11  
MessagePack|tinylib/msgp|Marshal|5000000|245|499|0  
MessagePack|tinylib/msgp|Unmarshal|3000000|365|0|0  
ProtobufV3|golang/protobuf/proto|Marshal|2000000|990|456|6  
ProtobufV3|golang/protobuf/proto|Unmarshal|2000000|645|208|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|5000000|267|80|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|5000000|234|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|5000000|291|80|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|10000000|229|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|5000000|276|80|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|5000000|242|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|5000000|268|80|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|5000000|229|0|0  

### MemInfo
MemInfo is a datastructure consisting of a number of ints, which are defined as int64 for Flatbuffers and int for Go data structures.  MemInfo's fields are all the keys in `/proc/meminfo`.

####MemInfo  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|1000000|1775|500|3  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|877|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|1000000|1544|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|50|32|1  
GenCode|andyleap/gencode|Marshal|10000000|180|0|0  
GenCode|andyleap/gencode|Unmarshal|10000000|172|0|0  
Gob|encoding/gob|Marshal|500000|2648|320|1  
Gob|encoding/gob|Unmarshal|20000000|65|0|0  
JSON|encoding/json|Marshal|200000|9526|5078|8  
JSON|encoding/json|Unmarshal|30000|37774|1527|42  
JSON|pquerna/ffjson|Marshal|200000|9302|5568|46  
JSON|pquerna/ffjson|Unmarshal|100000|15074|336|5  
JSON|pquerna/ffjson (buffer)|Marshal|200000|6879|1280|40  
JSON|pquerna/ffjson (buffer)|Unmarshal|100000|14422|341|5  
MessagePack|vmihailenco/msgpack|Marshal|300000|5635|2656|8  
MessagePack|vmihailenco/msgpack|Unmarshal|200000|9346|688|43  
MessagePack|tinylib/msgp|Marshal|2000000|1225|3325|0  
MessagePack|tinylib/msgp|Unmarshal|1000000|1639|0|0  
ProtobufV3|golang/protobuf/proto|Marshal|1000000|2341|1224|8  
ProtobufV3|golang/protobuf/proto|Unmarshal|1000000|1902|208|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|2000000|957|448|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|1000000|1268|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|2000000|954|448|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|1000000|1278|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|2000000|979|448|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|1000000|1267|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|2000000|978|448|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|1000000|1274|0|0  

### CPU
CPUInfo is a datastructure that represents the output of `/proc/cpuinfo` and consists mostly of strings and ints, with a couple of float32 fields.  Instead of querying the running system's `cpuinfo`, the benchmark uses a mixture of predefined values and random ints.  There are multiple versions of this benchmark to test the impact of varying slice (array) sizes on performance.  

####CPUInfo: 1  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|100000|17784|6580|37  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|1247|304|5   
Flatbuffers|google/flatbuffers/go|Marshal|1000000|2574|215|17  
Flatbuffers|google/flatbuffers/go|Unmarshal|20000000|76|32|1  
GenCode|andyleap/gencode|Marshal|10000000|174|0|0  
GenCode|andyleap/gencode|Unmarshal|2000000|990|400|14  
Gob|encoding/gob|Marshal|1000000|2192|179|1  
Gob|encoding/gob|Unmarshal|2000000|941|24|2  
JSON|encoding/json|Marshal|200000|10196|2632|7  
JSON|encoding/json|Unmarshal|100000|22583|720|27  
JSON|pquerna/ffjson|Marshal|300000|5303|1712|23  
JSON|pquerna/ffjson|Unmarshal|200000|8730|1536|21  
JSON|pquerna/ffjson (buffer)|Marshal|200000|5951|1712|23  
JSON|pquerna/ffjson (buffer)|Unmarshal|200000|8027|1536|21  
MessagePack|vmihailenco/msgpack|Marshal|200000|6331|1536|7  
MessagePack|vmihailenco/msgpack|Unmarshal|100000|11469|2096|51  
MessagePack|tinylib/msgp|Marshal|2000000|825|1482|0  
MessagePack|tinylib/msgp|Unmarshal|1000000|1800|400|14  
ProtobufV3|golang/protobuf/proto|Marshal|500000|3488|1496|7  
ProtobufV3|golang/protobuf/proto|Unmarshal|500000|2480|936|17  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|2000000|863|416|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|1000000|1930|728|16  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|2000000|787|416|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|1000000|1920|728|16  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|2000000|796|416|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|1000000|1749|728|16  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|2000000|893|416|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|1000000|1891|728|16  

####CPUInfo: 4  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|50000|38023|13204|127  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|999|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|200000|6267|414|17  
Flatbuffers|google/flatbuffers/go|Unmarshal|20000000|71|32|1  
GenCode|andyleap/gencode|Marshal|2000000|684|0|0  
GenCode|andyleap/gencode|Unmarshal|500000|3630|1600|56  
Gob|encoding/gob|Marshal|300000|5411|161|1  
Gob|encoding/gob|Unmarshal|2000000|1027|24|2  
JSON|encoding/json|Marshal|100000|32081|11592|9  
JSON|encoding/json|Unmarshal|20000|81164|1968|96  
JSON|pquerna/ffjson|Marshal|100000|15893|6401|78  
JSON|pquerna/ffjson|Unmarshal|50000|27076|4656|65  
JSON|pquerna/ffjson (buffer)|Marshal|100000|16144|6401|78  
JSON|pquerna/ffjson (buffer)|Unmarshal|50000|25446|4656|65  
MessagePack|vmihailenco/msgpack|Marshal|100000|16318|7296|9  
MessagePack|vmihailenco/msgpack|Unmarshal|50000|29894|6032|168  
MessagePack|tinylib/msgp|Marshal|300000|3388|7345|0  
MessagePack|tinylib/msgp|Unmarshal|200000|6058|1600|56  
ProtobufV3|golang/protobuf/proto|Marshal|200000|6805|4952|9  
ProtobufV3|golang/protobuf/proto|Unmarshal|200000|7711|3144|64  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|1000000|2194|1664|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|300000|5723|2936|63  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|1000000|2267|1664|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|300000|5400|2936|63  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|1000000|2274|1664|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|200000|5350|2936|63  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|1000000|2176|1664|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|300000|5462|2936|63  

####CPUInfo: 8  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|20000|72218|30388|250  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|901|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|100000|10333|677|17  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|57|32|1  
GenCode|andyleap/gencode|Marshal|1000000|1389|0|0  
GenCode|andyleap/gencode|Unmarshal|300000|5647|3200|112  
Gob|encoding/gob|Marshal|200000|9866|174|1  
Gob|encoding/gob|Unmarshal|1000000|1156|24|2  
JSON|encoding/json|Marshal|30000|47568|23880|10  
JSON|encoding/json|Unmarshal|10000|158297|3632|188  
JSON|pquerna/ffjson|Marshal|50000|29855|20835|151  
JSON|pquerna/ffjson|Unmarshal|30000|46545|8816|122  
JSON|pquerna/ffjson (buffer)|Marshal|50000|28763|20834|151  
JSON|pquerna/ffjson (buffer)|Unmarshal|30000|46577|8816|122  
MessagePack|vmihailenco/msgpack|Marshal|50000|26360|15488|10  
MessagePack|vmihailenco/msgpack|Unmarshal|20000|51932|11280|324  
MessagePack|tinylib/msgp|Marshal|200000|7947|21954|0  
ProtobufV3|golang/protobuf/proto|Marshal|100000|10019|10840|11  
ProtobufV3|golang/protobuf/proto|Unmarshal|100000|13828|6088|125  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|500000|3361|3328|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|200000|8752|5880|124  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|500000|3284|3328|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|200000|9252|5880|124  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|500000|3300|3328|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|200000|8783|5880|124  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|500000|3372|3328|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|200000|8840|5880|124  
MessagePack|tinylib/msgp|Unmarshal|100000|12111|3200|112  

####CPUInfo: 16  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|10000|136991|60340|491  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|939|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|100000|18966|1221|17  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|48|32|1  
GenCode|andyleap/gencode|Marshal|500000|2830|0|0  
GenCode|andyleap/gencode|Unmarshal|100000|10438|6400|224  
Gob|encoding/gob|Marshal|100000|19072|165|1  
Gob|encoding/gob|Unmarshal|1000000|1349|24|2  
JSON|encoding/json|Marshal|20000|90002|48456|11  
JSON|encoding/json|Unmarshal|5000|317165|7011|372  
JSON|pquerna/ffjson|Marshal|30000|58273|26883|352  
JSON|pquerna/ffjson|Unmarshal|20000|91851|17392|235  
JSON|pquerna/ffjson (buffer)|Marshal|30000|59089|26883|352  
JSON|pquerna/ffjson (buffer)|Unmarshal|20000|94249|17392|235  
MessagePack|vmihailenco/msgpack|Marshal|30000|49690|31872|11  
MessagePack|vmihailenco/msgpack|Unmarshal|20000|100851|22288|636  
MessagePack|tinylib/msgp|Marshal|100000|13653|29327|0  
MessagePack|tinylib/msgp|Unmarshal|50000|24060|6400|224  
ProtobufV3|golang/protobuf/proto|Marshal|100000|21575|29784|14  
ProtobufV3|golang/protobuf/proto|Unmarshal|50000|26267|11976|246  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|200000|6057|6656|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|100000|17103|11768|245  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|200000|6142|6656|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|100000|17248|11768|245  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|200000|5946|6656|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|100000|17136|11768|245  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|200000|6079|6656|1  

ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|100000|17159|11768|245  
### Message
Message is a basic datastructure that has a `Data` field that can be a varying number of bytes.  This structure is tested using data sizes of, in bytes: 16, 256, 1024, 4096.  The 4096 byte data test is not done with Cap'n Proto as it results in an error: capnp: NewMessage called on arena with data.  I will enable this benchmark when I get a better understanding of Cap'n Proto.  

####Message: 16  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|1000000|1605|404|7  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|871|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|5000000|244|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|48|32|1  
GenCode|andyleap/gencode|Marshal|50000000|25|0|0  
GenCode|andyleap/gencode|Unmarshal|100000000|22|0|0  
Gob|encoding/gob|Marshal|2000000|947|64|1  
Gob|encoding/gob|Unmarshal|3000000|539|12|1  
JSON|encoding/json|Marshal|1000000|1484|440|6  
JSON|encoding/json|Unmarshal|500000|3519|319|7  
JSON|pquerna/ffjson|Marshal|1000000|2504|2845|10  
JSON|pquerna/ffjson|Unmarshal|1000000|2035|384|7  
JSON|pquerna/ffjson (buffer)|Marshal|1000000|1882|2461|6  
JSON|pquerna/ffjson (buffer)|Unmarshal|1000000|2262|406|7  
MessagePack|vmihailenco/msgpack|Marshal|2000000|959|240|4  
MessagePack|vmihailenco/msgpack|Unmarshal|1000000|1516|181|8  
MessagePack|tinylib/msgp|Marshal|10000000|146|232|0  
ProtobufV3|golang/protobuf/proto|Marshal|3000000|640|328|5  
ProtobufV3|golang/protobuf/proto|Unmarshal|3000000|526|232|3  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|10000000|141|48|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|10000000|214|24|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|10000000|147|48|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|10000000|213|24|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|10000000|141|48|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|10000000|210|24|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|10000000|140|48|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|10000000|212|24|2  
MessagePack|tinylib/msgp|Unmarshal|5000000|239|0|0  

####Message: 64  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|1000000|1643|452|7  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|916|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|5000000|245|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|50|32|1  
GenCode|andyleap/gencode|Marshal|50000000|29|0|0  
GenCode|andyleap/gencode|Unmarshal|50000000|27|0|0  
Gob|encoding/gob|Marshal|2000000|951|64|1  
Gob|encoding/gob|Unmarshal|3000000|525|11|0  
JSON|encoding/json|Marshal|1000000|1987|859|7  
JSON|encoding/json|Unmarshal|300000|4699|367|7  
JSON|pquerna/ffjson|Marshal|500000|2544|2844|10  
JSON|pquerna/ffjson|Unmarshal|500000|2781|656|8  
JSON|pquerna/ffjson (buffer)|Marshal|1000000|2253|2461|6  
JSON|pquerna/ffjson (buffer)|Unmarshal|500000|3083|668|8  
MessagePack|vmihailenco/msgpack|Marshal|1000000|1382|432|5  
MessagePack|vmihailenco/msgpack|Unmarshal|1000000|1508|181|8  
MessagePack|tinylib/msgp|Marshal|10000000|184|388|0  
MessagePack|tinylib/msgp|Unmarshal|5000000|276|0|0  
ProtobufV3|golang/protobuf/proto|Marshal|3000000|655|360|5  
ProtobufV3|golang/protobuf/proto|Unmarshal|3000000|542|280|3  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|10000000|155|96|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|10000000|232|72|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|10000000|157|96|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|10000000|229|72|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|10000000|163|96|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|10000000|229|72|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|10000000|163|96|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|10000000|231|72|2  

####Message: 256  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|1000000|1932|660|7  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|913|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|5000000|254|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|50|32|1  
GenCode|andyleap/gencode|Marshal|30000000|40|0|0  
GenCode|andyleap/gencode|Unmarshal|50000000|27|0|0  
Gob|encoding/gob|Marshal|2000000|984|64|1  
Gob|encoding/gob|Unmarshal|2000000|632|12|1  
JSON|encoding/json|Marshal|500000|2539|1400|7  
JSON|encoding/json|Unmarshal|200000|9225|575|7  
JSON|pquerna/ffjson|Marshal|500000|3478|3645|11  
JSON|pquerna/ffjson|Unmarshal|300000|4719|1120|8  
JSON|pquerna/ffjson (buffer)|Marshal|500000|2318|2461|6  
JSON|pquerna/ffjson (buffer)|Unmarshal|300000|4919|929|12  
MessagePack|vmihailenco/msgpack|Marshal|1000000|1428|624|5  
MessagePack|vmihailenco/msgpack|Unmarshal|1000000|1555|181|8  
MessagePack|tinylib/msgp|Marshal|5000000|322|1005|0  
MessagePack|tinylib/msgp|Unmarshal|5000000|304|0|0  
ProtobufV3|golang/protobuf/proto|Marshal|2000000|727|552|5  
ProtobufV3|golang/protobuf/proto|Unmarshal|3000000|635|472|3  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|10000000|234|288|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|5000000|292|264|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|10000000|232|288|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|5000000|304|264|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|10000000|237|288|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|5000000|308|264|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|10000000|233|288|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|5000000|294|264|2  

####Message: 1024  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|500000|2986|1492|7  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|895|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|5000000|278|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|50|32|1  
GenCode|andyleap/gencode|Marshal|20000000|86|0|0  
GenCode|andyleap/gencode|Unmarshal|50000000|29|0|0  
Gob|encoding/gob|Marshal|2000000|1014|64|1  
Gob|encoding/gob|Unmarshal|2000000|657|11|0  
JSON|encoding/json|Marshal|300000|6135|6040|8  
JSON|encoding/json|Unmarshal|50000|27929|1439|7  
JSON|pquerna/ffjson|Marshal|300000|4861|4669|11  
JSON|pquerna/ffjson|Unmarshal|100000|13014|3520|8  
JSON|pquerna/ffjson (buffer)|Marshal|500000|3705|2461|6  
JSON|pquerna/ffjson (buffer)|Unmarshal|100000|12809|3511|8  
MessagePack|vmihailenco/msgpack|Marshal|1000000|1854|1392|5  
MessagePack|vmihailenco/msgpack|Unmarshal|1000000|1574|181|8  
MessagePack|tinylib/msgp|Marshal|2000000|1014|4288|0  
MessagePack|tinylib/msgp|Unmarshal|5000000|385|0|0  
ProtobufV3|golang/protobuf/proto|Marshal|1000000|1058|1416|5  
ProtobufV3|golang/protobuf/proto|Unmarshal|2000000|910|1240|3  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|5000000|616|1152|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|3000000|641|1032|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|5000000|615|1152|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|3000000|664|1032|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|5000000|636|1152|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|3000000|638|1032|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|5000000|590|1152|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|3000000|655|1032|2  

####Message: 2048  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|300000|4629|2644|7  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|932|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|5000000|312|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|52|32|1  
GenCode|andyleap/gencode|Marshal|10000000|172|0|0  
GenCode|andyleap/gencode|Unmarshal|50000000|31|0|0  
Gob|encoding/gob|Marshal|2000000|1166|64|1  
Gob|encoding/gob|Unmarshal|2000000|722|11|0  
JSON|encoding/json|Marshal|200000|8327|7064|8  
JSON|encoding/json|Unmarshal|30000|52073|2591|7  
JSON|pquerna/ffjson|Marshal|100000|10838|10845|12  
JSON|pquerna/ffjson|Unmarshal|50000|22905|6720|8  
JSON|pquerna/ffjson (buffer)|Marshal|300000|5107|2461|6  
JSON|pquerna/ffjson (buffer)|Unmarshal|100000|20191|4293|12  
MessagePack|vmihailenco/msgpack|Marshal|1000000|2534|2544|5  
MessagePack|vmihailenco/msgpack|Unmarshal|1000000|1660|181|8  
MessagePack|tinylib/msgp|Marshal|1000000|1973|8085|0  
MessagePack|tinylib/msgp|Unmarshal|3000000|411|0|0  
ProtobufV3|golang/protobuf/proto|Marshal|1000000|1746|2568|5  
ProtobufV3|golang/protobuf/proto|Unmarshal|1000000|1479|2264|3  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|2000000|1134|2304|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|2000000|1027|2056|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|2000000|1084|2304|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|1000000|1253|2056|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|2000000|1093|2304|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|2000000|1127|2056|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|2000000|1080|2304|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|2000000|1135|2056|2  

####Message: 4096  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
Flatbuffers|google/flatbuffers/go|Marshal|3000000|433|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|49|32|1  
GenCode|andyleap/gencode|Marshal|3000000|463|0|0  
GenCode|andyleap/gencode|Unmarshal|50000000|32|0|0  
Gob|encoding/gob|Marshal|1000000|1508|64|1  
Gob|encoding/gob|Unmarshal|2000000|856|12|1  
JSON|encoding/json|Marshal|100000|15373|15512|9  
JSON|encoding/json|Unmarshal|10000|101432|4895|7  
JSON|pquerna/ffjson|Marshal|100000|13716|10845|12  
JSON|pquerna/ffjson|Unmarshal|30000|43264|13120|8  
JSON|pquerna/ffjson (buffer)|Marshal|200000|8539|2461|6  
JSON|pquerna/ffjson (buffer)|Unmarshal|50000|31151|7641|12  
MessagePack|vmihailenco/msgpack|Marshal|500000|3334|4848|5  
MessagePack|vmihailenco/msgpack|Unmarshal|1000000|1831|181|8  
MessagePack|tinylib/msgp|Marshal|500000|3296|14893|0  
MessagePack|tinylib/msgp|Unmarshal|3000000|537|0|0  
ProtobufV3|golang/protobuf/proto|Marshal|500000|3017|4872|5  
ProtobufV3|golang/protobuf/proto|Unmarshal|1000000|2357|4312|3  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|1000000|2415|4608|1  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|1000000|2161|4104|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|500000|2226|4608|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|1000000|2143|4104|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|1000000|2293|4608|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|1000000|2211|4104|2  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|1000000|2379|4608|1  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|1000000|2053|4104|2  

#### RedditAccount
RedditAccount is the data structure for a Reddit Account.  It consists of a `thing` whose object is `account` per https://github.com/reddit/reddit/wiki/JSON.

####RedditAccount  
Format|Package|Operation|Ops|ns/Op|B/Op|Allocs/Op  
:--|:--|:--|--:|--:|--:|--:  
CapnProto2|zombiezen.com/go/capnproto2|Marshal|500000|3997|1126|17  
CapnProto2|zombiezen.com/go/capnproto2|Unmarshal|2000000|967|304|5  
Flatbuffers|google/flatbuffers/go|Marshal|2000000|845|0|0  
Flatbuffers|google/flatbuffers/go|Unmarshal|30000000|52|32|1  
GenCode|andyleap/gencode|Marshal|10000000|108|0|0  
GenCode|andyleap/gencode|Unmarshal|5000000|348|203|5  
Gob|encoding/gob|Marshal|1000000|1533|144|1  
Gob|encoding/gob|Unmarshal|5000000|322|3|0  
JSON|encoding/json|Marshal|500000|3872|1364|6  
JSON|encoding/json|Unmarshal|100000|10176|566|11  
JSON|pquerna/ffjson|Marshal|500000|2572|1488|14  
JSON|pquerna/ffjson|Unmarshal|300000|4262|764|11  
JSON|pquerna/ffjson (buffer)|Marshal|1000000|1521|305|9  
JSON|pquerna/ffjson (buffer)|Unmarshal|300000|4635|784|13  
MessagePack|vmihailenco/msgpack|Marshal|500000|3288|1424|6  
MessagePack|vmihailenco/msgpack|Unmarshal|500000|4065|648|25  
MessagePack|tinylib/msgp|Marshal|3000000|459|921|0  
MessagePack|tinylib/msgp|Unmarshal|2000000|912|204|5  
ProtobufV3|golang/protobuf/proto|Marshal|2000000|1238|696|6  
ProtobufV3|golang/protobuf/proto|Unmarshal|2000000|1264|508|7  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Marshal|20000000|64|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gofast|Unmarshal|50000000|28|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Marshal|20000000|64|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogofast|Unmarshal|50000000|29|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Marshal|20000000|64|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogofaster|Unmarshal|50000000|30|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Marshal|20000000|64|0|0  
ProtobufV3|gogo/protobuf/protoc-gen-gogoslick|Unmarshal|50000000|29|0|0  

## Notes
There is probably room for improvement, esp. with Cap'n Proto.

Running all of the benchmarks can take 10 minutes or more, depending on your system speed.

### Flatbuffers
The Flatbuffers benches reuses the `builder` object; doing a `Reset()` instead of creating a new `builder`.  Creating a new `builder` via `flatbuffers.NewBuilder(0)` results in a significant degradation of performance and should be avoided when serializing data with Flatbuffers.
