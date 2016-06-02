using Go = import "../../../../zombiezen.com/go/capnproto2/go.capnp";
# using relative because that is the only way I could get capn compile -ogo
# to work; otherwise I get an Import failed: error.
# using Go = import "zombiezen.com/go/capnproto2/go.capnp";
# from what the docs say, the above should work, but it doesn't
@0xc29b0335fa0bc73a;
$Go.package("capnp");
$Go.import("github.com/mohae/serial-bowl/capnp");

struct BasicMemInfo {
    memTotal @0 :Int64;
	memFree @1 :Int64;
	memAvailable @2 :Int64;
	buffers @3 :Int64;
	cached  @4 :Int64;
	swapCached @5 :Int64;
	swapTotal  @6 :Int64;
	swapFree   @7 :Int64;
}

struct MemInfo {
    memTotal @0 :Int64;
	memFree @1 :Int64;
	memAvailable @2 :Int64;
	buffers @3 :Int64;
	cached  @4 :Int64;
	swapCached @5 :Int64;
	active @6 :Int64;
	inactive @7 :Int64;
	activeAnon @8 :Int64;
	inactiveAnon @9 :Int64;
	activeFile @10 :Int64;
	inactiveFile @11 :Int64;
	unevictable @12 :Int64;
	mlocked @13 :Int64;
	swapTotal @14 :Int64;
	swapFree @15 :Int64;
	dirty @16 :Int64;
	writeback @17 :Int64;
	anonPages @18 :Int64;
	mapped @19 :Int64;
	shmem @20 :Int64;
	slab @21 :Int64;
	sReclaimable @22 :Int64;
	sUnreclaim @23 :Int64;
	kernelStack @24 :Int64;
	nFSUnstable @25 :Int64;
	bounce @26 :Int64;
	writebackTmp @27 :Int64;
	commitLimit @28 :Int64;
	vmallocTotal @29 :Int64;
	vmallocUsed @30 :Int64;
	vmallocChunk @31 :Int64;
	hardwareCorrupted @32 :Int64;
	anonHugePages @33 :Int64;
	hugePagesTotal @34 :Int64;
	hugePagesFree @35 :Int64;
	hugePagesRsvd @36 :Int64;
	hugepagesize @37 :Int64;
	directMap4k @38 :Int64;
	directMap2M @39 :Int64;
}

struct Message {
    iD @0 :Data;
    destID @1 :UInt32;
    type @2 :Int8;
    kind @3 :Int16;
    data @4 :Data;
}

struct RedditAccount {
    iD @0 :Text;
    name @1 :Text;
    kind @2 :Text;
    data @3 :AccountData;
}

struct AccountData {
    commentKarma @0 :Int64;
    hasMail @1 :Bool;
    hasModMail @2 :Bool;
    iD @3 :Text;
    inboxCount @4 :Int64;
    isFriend @5 :Bool;
    isGold @6 :Bool;
    linkKarma @7 :Int64;
    modHash @8 :Text;
    name @9 :Text;
    over18 @10 :Bool;
}

struct CPUInfo {
	cpus @0 :List(CPU);
}

struct CPU {
	processor @0 :Int16;
	vendorID @1 :Text;
	cPUFamily @2 :Text;
	model @3 :Text;
	modelName @4 :Text;
	stepping @5 :Text;
	microcode @6 :Text;
	cPUMHz @7 :Float32;
	cacheSize @8 :Text;
	physicalID @9 :Int16;
	siblings @10 :Int16;
	coreID @11 :Int16;
	cPUCores @12 :Int16;
	apicID @13 :Int16;
	initialApicID @14 :Int16;
	fPU @15 :Text;
	fPUException @16 :Text;
	cPUIDLevel @17 :Text;
	wP @18 :Text;
	flags @19 :Text;
	bogoMIPS @20 :Float32;
	cLFlushSize @21 :Text;
	cacheAlignment @22 :Text;
	addressSizes @23 :Text;
	powerManagement @24 :Text;
}
