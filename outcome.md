### *1.不同大小数据时的set/get性能*
> 说明：<font color="gree">使用命令</font>
```
redis-benchmark -h 127.0.0.1 -p 6379 -t set,get -q
```
> 测试结果
---
+ 10B
```
SET: 96432.02 requests per second, p50=0.255 msec                   
GET: 100704.94 requests per second, p50=0.255 msec 
```
+ 20B
```
SET: 102145.05 requests per second, p50=0.255 msec                    
GET: 101112.23 requests per second, p50=0.255 msec
```
+ 50B
```
SET: 102040.81 requests per second, p50=0.255 msec                    
GET: 99206.34 requests per second, p50=0.255 msec
```
+ 100B
```
SET: 101936.80 requests per second, p50=0.255 msec                    
GET: 99206.34 requests per second, p50=0.255 msec
```
+ 200B
```
SET: 102564.10 requests per second, p50=0.255 msec                    
GET: 99403.58 requests per second, p50=0.255 msec
```
+ 1KB
```
SET: 102986.61 requests per second, p50=0.255 msec                    
GET: 100200.40 requests per second, p50=0.255 msec
```
+ 5KB
```
SET: 100401.61 requests per second, p50=0.255 msec                   
GET: 94073.38 requests per second, p50=0.263 msec
```
### 2.计算平均每个key的空间大小
> 插入key前info memory <font color="yellow">(已用空间1.07M)</fon>
```
# Memory
used_memory:1117360
used_memory_human:1.07M
used_memory_rss:5935104
used_memory_rss_human:5.66M
used_memory_peak:4129600
used_memory_peak_human:3.94M
used_memory_peak_perc:27.06%
used_memory_overhead:1079952
used_memory_startup:1062512
used_memory_dataset:37408
used_memory_dataset_perc:68.20%
allocator_allocated:1080432
allocator_active:5897216
allocator_resident:5897216
total_system_memory:8589934592
total_system_memory_human:8.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:5.46
allocator_frag_bytes:4816784
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.01
rss_overhead_bytes:37888
mem_fragmentation_ratio:5.49
mem_fragmentation_bytes:4854672
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```
> 插入key后info memory <font color="yellow">(已用空间10.84M)</font>
```
# Memory
used_memory:11371504
used_memory_human:10.84M
used_memory_rss:16711680
used_memory_rss_human:15.94M
used_memory_peak:12675040
used_memory_peak_human:12.09M
used_memory_peak_perc:89.72%
used_memory_overhead:3614240
used_memory_startup:1062512
used_memory_dataset:7757264
used_memory_dataset_perc:75.25%
allocator_allocated:11334576
allocator_active:16673792
allocator_resident:16673792
total_system_memory:8589934592
total_system_memory_human:8.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.47
allocator_frag_bytes:5339216
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.47
mem_fragmentation_bytes:5377104
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```
> 计算结果<font color="gree">如下</font>：  
> + key总数为50250 占用空间10.84-1.07=9.77MB 平均每个key占203.87B
