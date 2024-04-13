
GCS:
1. Type:

| Storage Class    | Name for APIs and CLIs | Minimum storage duration | Retrieval fees) | Typical monthly availability1                                      |
| ---------------- | ---------------------- | ------------------------ | --------------- | ------------------------------------------------------------------ |
| Standard storage | `STANDARD`             | None                     | None            | - >99.99% in multi-regions and dual-regions<br>- 99.99% in regions |
| Nearline storage | `NEARLINE`             | 30 days                  | Yes             | - 99.95% in multi-regions and dual-regions<br>- 99.9% in regions   |
| Coldline storage | `COLDLINE`             | 90 days                  | Yes             | - 99.95% in multi-regions and dual-regions<br>- 99.9% in regions   |
| Archive storage  | `ARCHIVE`              | 365 days                 | Yes             | - 99.95% in multi-regions and dual-regions<br>- 99.9% in regions   |
3. Object Versioning:
	1. This feature maintains a history of versions for each object in the bucket.
	2. This is critical for protecting against accidental deletions and data corruption.
4. Bucket lock:
	1. set a retention policy on a bucket that prohibits data from being deleted for a set period of time.
5. Lifecycle Management:
	1. Set up lifecycle rules on your buckets to manage object versions and storage costs effectively.