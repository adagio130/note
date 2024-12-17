
- the complexity needs to be O(n) so that it wouldn't reach the TL
- only can traverse the nums on time
- keep the only largest number of index in the temp indices array.
- need to pop out the largest number of index which is out of the k range
- loop the temp indices to check the numbers if greater than the current idx.
- if it's less than the current number. pull it out from the tail.
- put the current index into the temp indices.
- if current indices is larger than k - 1, put the first index into the result.