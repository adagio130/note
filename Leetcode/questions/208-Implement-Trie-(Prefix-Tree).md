1. A Trie (pronounced "try") or Prefix Tree is a tree data structure used to efficiently retrieve a key from a dataset of strings. Here are some of its key applications:
	1. auto complete
	2. spell chacker
	3. IP routing
2. Tries offer advantages in certain operations:
	1. Prefix Searches: Tries excel in finding all keys with a common prefix.
	2. Lexicographical Ordering: They allow efficient enumeration of keys in lexicographical order.
	3. Space Efficiency: For keys with common prefixes, Tries use less space compared to hash tables, which can suffer from hash collisions and increased search times as they grow.
3. data structure
```
type Trie struct {
	Children map[rune]*Trie
	IsEnd bool
}
```
5. 
