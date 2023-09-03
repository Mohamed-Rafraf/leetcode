## Majority Element - Boyer-Moore Majority Voting Algorithm


Given an array `nums` of size `n`, return the majority element.

The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

 

### Example 1:

Input: `nums = [3,2,3]`

Output: `3`

### Example 2:

Input: `nums = [2,2,1,1,1,2,2]`

Output: `2`
 

### Constraints:

* `n == nums.length`
* `1 <= n <= 5 * 104`
* `-109 <= nums[i] <= 109`
 

Follow-up: Could you solve the problem in linear time and in O(1) space?


---

## Explanation

The main idea behind the Boyer-Moore algorithm is quite ingenious: every time we find two different elements, we discard both of them. This doesn't affect the majority element because even if one of the discarded elements is the majority element, it still remains the majority among the remaining elements. The process repeats until a potential candidate for the majority element is found.

Here's a breakdown of the algorithm:

#### Initialization: S

tart with an initial candidate and a count.

* candidate = None
* count = 0

Iterate through the list:

For each element e in the list:
* If count == 0, set candidate = e.
* If e == candidate, increment count. Otherwise, decrement count.


#### Validation: 
At the end of this procedure, candidate may be the majority element. However, we have to iterate through the list again to confirm that it occurs more than n/2 times.

### Why it works:
Imagine pairing each majority element with a non-majority element and then discarding them. We would still be left with the majority element at the end. This is essentially what the algorithm does. The counting mechanism keeps track of the number of unpaired instances of the current candidate.