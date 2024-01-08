
1. cache-aside (lazy loading) 
	1. a reactive approach
	2. updated after the data is requested
		1. When your application needs to read data from the database, it checks the cache first to determine whether the data is available.
		2. If the data is available (_a cache hit_), the cached data is returned, and the response is issued to the caller.
		3. If the data isn’t available (_a cache miss_), the database is queried for the data. The cache is then populated with the data that is retrieved from the database, and the data is returned to the caller.
	3. pros and cons
		1. pros
			1. The cache contains only data that the application actually requests, which helps keep the cache size cost-effective.
			2. Implementing this approach is straightforward and produces immediate performance gains, whether you use an application framework that encapsulates lazy caching or your own custom application logic.
		2. cons
			1. some overhead is added to the initial response time because additional roundtrips to the cache and database are needed.
2. write-through
	1. a proactive approach
	2. updated immediately when the primary database is updated