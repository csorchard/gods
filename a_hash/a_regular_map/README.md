## HashMap

<!-- TOC -->
  * [HashMap](#hashmap)
    * [Collision Handling](#collision-handling)
      * [Chain address method](#chain-address-method)
      * [Open addressing method (`collision handling`)](#open-addressing-method-collision-handling)
        * [Free bucket `address detection` method in Open Addressing](#free-bucket-address-detection-method-in-open-addressing)
    * [Hash Map's](#hash-maps)
        * [C++ std::unordered_map/boost::unordered_map](#c-stdunordered_mapboostunordered_map)
        * [go map](#go-map)
        * [Swiss Table](#swiss-table)
        * [ClickHouse hashtable](#clickhouse-hashtable)
<!-- TOC -->

### Collision Handling

Different keys are mapped to the same bucket through the hash function, which is called hash collision. The most common
collision handling mechanisms in various implementations are chaining and open-addressing.

#### Chain address method

In a hash table, each bucket stores a linked list, and different elements with the same hash value are placed in the
linked list. This is the approach typically used by **C++ standard containers** and **Golang Map**.

Advantage:

- The most simple and intuitive implementation
- Less space wasted

#### Open addressing method (`collision handling`)

If a collision occurs during insertion, starting from the hash bucket where the collision occurred, a free bucket will
be found in a certain order.

Advantage:

- There is only one pointer jump for each insertion or search operation, which is more friendly to CPU cache.
- All data is stored in a continuous memory, with less memory fragmentation


##### Free bucket `address detection` method in Open Addressing

In the open addressing method, if the bucket returned by the hash function is already occupied by other keys, you need
to use preset rules to find free buckets in adjacent buckets. The most common methods are as follows (assuming there are
|T| buckets in total and the hash function is H(k)):

- Linear probing: search one by one from next bucket.
- Quadratic probing: search in power of 2 from next bucket
- Double hashing: Use two different hash functions to detect H(k, i) = (H 1 (k) + i * H 2 (k)) mod |T| in sequence.

Compared with other methods, the linear detection method requires the largest number of buckets to be detected on
average. However, the linear detection method always accesses memory sequentially and continuously, which is the most
cache-friendly. Therefore, when the probability of conflict is low (max load factor is small), the linear detection
method is the fastest way.

There are other more sophisticated detection methods, such as cuckoo hashing, hopscotch hashing, and robin hood
hashing. However, they are all designed for situations where the max load factor is large (above 0.6). When max load
factor = 0.5, the performance in actual tests is not as good as the most original and direct linear detection.
[Good article](https://zhuanlan.zhihu.com/p/277732297)

### Hash Map's

##### C++ std::unordered_map/boost::unordered_map
The chain address method is used to handle collisions. Default max load factor = 1, growth factor = 2.

##### go map
Built-in map in golang uses the chain address method.

##### Swiss Table


##### ClickHouse hashtable

- open addressing
- Linear detection
- max load factor = 0.5，growth factor = 4
- Integer hash function based on CRC32 instructions

Details
- https://github.com/ClickHouse/ClickHouse/pull/5417
- https://www.researchgate.net/publication/339879042_SAHA_A_String_Adaptive_Hash_Table_for_Analytical_Databases
