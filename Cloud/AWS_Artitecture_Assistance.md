## storage type:
1. S3 (object)
2. glacier
3. EBS (block)
4.

selecting storage factor:
1. capacity(size)
2. performance
3. cost

### S3:
S3 storage is distributed across 3 AZs, except 1A, which uses one zone and is less expensive.
S3 allows for automatic data classfication
files stored in s3 can be encrypted
object in S3 buckets have eventual consistency
object in Elastic Block Stores are consistent
very good use of S3 buckets is static websites
1. object storage
- 
2. getting data into s3
- API
- Amazion direct connect
- Storage gateway
- kinesis firehose
- tranfer acceleration
3. snow family
- snowball
- snoball edge
- snowmobile
### S3 operations:
1. createing and deleting buckets
2. writing objects
3. reading objects
4. deleting objects
5. managing objects proeprties
6. listing keys in buckets

buckets:
regions:
objects:
keys:
object urls:



## Four aspects:

1. Resilient
2. Performant
3. Secure
4. Cost