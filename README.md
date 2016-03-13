# serial-bowl

Some Go serialization benchmarks.

* flatbuffers
* encoding/json

Encode and decode benchmarks are performed with each listed protocol using structs with randomly generated data.  Each protocol uses the same data for a given run.  The data is randomly generated prior to running any of the benchmarks.

## Notes
Currently, the data generator assumes that the benchmark is being run on a 64bit system.  That is, the max value of an int is defined to be `1<<63-1`.
