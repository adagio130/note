#flashcards/questions 
==a==
1. count the min subarray
2. count the max subarray
3. count the total
4. If `minSum` == `totalSum` return `maxSum`, otherwise return `max(maxSum, totalSum - minSum)`
5. we can use total minus min to find the circular max sum
6. but if total equals to min which means all the elements are negative, so that we just return the max sum
7. otherwise, we return the max(maxSum, totalSum-minSum)