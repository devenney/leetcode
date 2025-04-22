# Last updated: 4/22/2025, 6:15:03 PM
# Solo solution.
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        o = 0
        for outer in nums:
            i = o + 1
            for inner in nums[i:]:
                if outer + inner == target:
                    return [o, i] 
                i += 1
            o += 1