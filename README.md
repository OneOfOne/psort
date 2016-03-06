# psort
Parallel sort, based on the std lib.

This is a PoC parellel version of the standard library's sort.Sort, shows quiet an improvement however it uses a lot more allocations.

It is slower for < ~10k elements as demonstrated in the stdlib's sort benchmarks, however for >10k the improvements are over 15%.

## index/suffixarray

``` sh
# go test -v -bench=NewIndex -benchmem -count 5
➜ benchstat {old,new}.txt
name              old time/op    new time/op    delta
NewIndexRandom-8     203ms ± 0%     208ms ± 2%    +2.49%  (p=0.008 n=5+5)
NewIndexRepeat-8     339ms ± 2%     402ms ± 2%   +18.86%  (p=0.008 n=5+5)

name              old alloc/op   new alloc/op   delta
NewIndexRandom-8    16.1MB ± 0%    16.1MB ± 0%    +0.08%  (p=0.008 n=5+5)
NewIndexRepeat-8    39.0MB ± 0%    39.0MB ± 0%    +0.00%  (p=0.008 n=5+5)

name              old allocs/op  new allocs/op  delta
NewIndexRandom-8      33.6 ±10%     299.2 ± 1%  +790.48%  (p=0.008 n=5+5)
NewIndexRepeat-8      73.8 ±11%     103.2 ± 2%   +39.84%  (p=0.008 n=5+5)
```


## sort
``` sh
# go test -bench=kSort -count 5 -benchmem
➜ benchstat {old,new}.txt
name            old time/op    new time/op    delta
SortString1K-8     157µs ± 1%     198µs ± 2%   +26.64%          (p=0.008 n=5+5)
SortInt1K-8       73.2µs ± 1%   115.2µs ± 0%   +57.30%          (p=0.008 n=5+5)
SortInt64K-8      6.84ms ± 2%    4.11ms ± 3%   -39.93%          (p=0.008 n=5+5)
Sort1e2-8         42.8µs ± 1%    42.5µs ± 1%      ~             (p=0.421 n=5+5)  # this falls back to the standard sort.
Sort1e4-8         8.60ms ± 1%    7.23ms ± 1%   -15.99%          (p=0.008 n=5+5)
Sort1e6-8          1.35s ± 7%     0.98s ± 7%   -27.46%          (p=0.008 n=5+5)

name            old alloc/op   new alloc/op   delta
SortString1K-8     32.0B ± 0%     80.4B ± 1%  +151.25%          (p=0.008 n=5+5)
SortInt1K-8        32.0B ± 0%     80.0B ± 0%  +150.00%          (p=0.008 n=5+5)
SortInt64K-8       32.0B ± 0%     80.0B ± 0%  +150.00%          (p=0.008 n=5+5)
Sort1e2-8           224B ± 0%      224B ± 0%      ~     (all samples are equal)
Sort1e4-8           224B ± 0%      560B ± 0%  +150.00%          (p=0.008 n=5+5)
Sort1e6-8           224B ± 0%      560B ± 0%  +150.00%          (p=0.008 n=5+5)

name            old allocs/op  new allocs/op  delta
SortString1K-8      1.00 ± 0%      2.00 ± 0%  +100.00%          (p=0.008 n=5+5)
SortInt1K-8         1.00 ± 0%      2.00 ± 0%  +100.00%          (p=0.008 n=5+5)
SortInt64K-8        1.00 ± 0%      2.00 ± 0%  +100.00%          (p=0.008 n=5+5)
Sort1e2-8           7.00 ± 0%      7.00 ± 0%      ~     (all samples are equal)
Sort1e4-8           7.00 ± 0%     14.00 ± 0%  +100.00%          (p=0.008 n=5+5)
Sort1e6-8           7.00 ± 0%     14.00 ± 0%  +100.00%          (p=0.008 n=5+5)

```