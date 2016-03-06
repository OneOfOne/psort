# psort
Parallel sort, based on the std lib.

This is a PoC parellel version of the standard library's sort.Sort, shows quiet an improvement however it uses a lot more allocations.

It is slower for < ~10k elements as demonstrated in the stdlib's sort benchmarks, however for >5k the improvements are over 15%.

* updated benchmarks due to golang/go#14677

## index/suffixarray

``` sh
# go test -v -bench=NewIndex -benchmem -count 5
➜ benchstat {old,new}.txt
name              old time/op    new time/op    delta
➜ benchstat  /tmp/{old,new}.txt
name              old time/op    new time/op    delta
NewIndexRandom-8     202ms ± 1%     177ms ± 2%   -12.29%  (p=0.008 n=5+5)
NewIndexRepeat-8     339ms ± 2%     279ms ± 2%   -17.82%  (p=0.008 n=5+5)

name              old alloc/op   new alloc/op   delta
NewIndexRandom-8    16.1MB ± 0%    16.1MB ± 0%    +0.08%  (p=0.008 n=5+5)
NewIndexRepeat-8    39.0MB ± 0%    39.0MB ± 0%    +0.00%  (p=0.008 n=5+5)

name              old allocs/op  new allocs/op  delta
NewIndexRandom-8      32.8 ± 4%     291.8 ± 1%  +789.63%  (p=0.008 n=5+5)
NewIndexRepeat-8      66.4 ± 5%      96.8 ± 2%   +45.78%  (p=0.008 n=5+5)
```


## sort
``` sh
# go test -bench=kSort -count 5 -benchmem
➜ benchstat {old,new}.txt
name            old time/op    new time/op    delta
SortString1K-8     152µs ± 1%     157µs ± 0%    +3.51%          (p=0.008 n=5+5)
SortInt1K-8       70.3µs ± 1%    75.2µs ± 1%    +6.94%          (p=0.008 n=5+5)
SortInt64K-8      6.44ms ± 1%    2.60ms ± 1%   -59.70%          (p=0.008 n=5+5)
Sort1e2-8         40.7µs ± 1%    41.5µs ± 3%      ~             (p=0.056 n=5+5)
Sort1e4-8         8.62ms ± 3%    5.31ms ± 0%   -38.33%          (p=0.008 n=5+5)
Sort1e6-8          1.30s ± 2%     0.69s ± 2%   -46.52%          (p=0.008 n=5+5)

name            old alloc/op   new alloc/op   delta
SortString1K-8     32.0B ± 0%     80.0B ± 0%  +150.00%          (p=0.008 n=5+5)
SortInt1K-8        32.0B ± 0%     80.0B ± 0%  +150.00%          (p=0.008 n=5+5)
SortInt64K-8       32.0B ± 0%     80.0B ± 0%  +150.00%          (p=0.016 n=5+4)
Sort1e2-8           224B ± 0%      224B ± 0%      ~     (all samples are equal)
Sort1e4-8           224B ± 0%      560B ± 0%  +150.00%          (p=0.016 n=5+4)
Sort1e6-8           224B ± 0%      560B ± 0%  +150.00%          (p=0.008 n=5+5)

name            old allocs/op  new allocs/op  delta
SortString1K-8      1.00 ± 0%      2.00 ± 0%  +100.00%          (p=0.008 n=5+5)
SortInt1K-8         1.00 ± 0%      2.00 ± 0%  +100.00%          (p=0.008 n=5+5)
SortInt64K-8        1.00 ± 0%      2.00 ± 0%  +100.00%          (p=0.008 n=5+5)
Sort1e2-8           7.00 ± 0%      7.00 ± 0%      ~     (all samples are equal)
Sort1e4-8           7.00 ± 0%     14.00 ± 0%  +100.00%          (p=0.008 n=5+5)
Sort1e6-8           7.00 ± 0%     14.00 ± 0%  +100.00%          (p=0.008 n=5+5)


```