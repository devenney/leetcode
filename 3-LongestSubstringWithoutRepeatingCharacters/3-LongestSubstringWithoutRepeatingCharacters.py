# Last updated: 24/07/2025, 12:22:31
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        start = 0
        max_len = 0
        max_substring = ""
        seen = {}

        for end, char in enumerate(s):
            if char in seen and seen[char] >= start:
                start = seen[char] + 1
            seen[char] = end
            if end - start + 1 > max_len:
                max_len = end - start + 1
                max_substring = s[start:end+1]

        return len(max_substring)