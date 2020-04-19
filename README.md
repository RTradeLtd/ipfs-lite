# ipfs-litex


`ipfs-litex` is a modified version of `hsanjuan/ipfs-lite` to showcase how you can use TemporalX as a replacement for common internals used by all IPFS implementations to get maximal performance, without having to use TemporalX directly. The code in this repository has been modified to use TemporalX as a replacement for the `ipld.DAGService` and `blockstore.Blockstore` interfaces.