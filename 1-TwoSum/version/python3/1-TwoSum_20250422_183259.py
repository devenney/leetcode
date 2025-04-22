# Last updated: 4/22/2025, 6:32:59 PM
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # track num: index
        seen = {}
        for i, num in enumerate(nums):
            # calculate diff between our number and the target
            diff = target - num

            # if we've seen this diff as a previous number, we have solved
            if diff in seen:
                return [seen[diff], i]

            # otherwise just track this number (key) and its index (value)
            seen[num] = i