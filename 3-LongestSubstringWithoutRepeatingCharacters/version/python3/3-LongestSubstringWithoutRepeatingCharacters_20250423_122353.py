# Last updated: 4/23/2025, 12:23:53 PM
# Needs review - optimal solution from ChatGPT.
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        char_index_map = {}
        left = 0
        max_len = 0

        for right, char in enumerate(s):
            # If character is already in map and its index is >= left boundary,
            # it means it's part of the current window - time to move left.
            if char in char_index_map and char_index_map[char] >= left:
                left = char_index_map[char] + 1
            
            # Update latest index of the current character
            char_index_map[char] = right

            # Update max length
            max_len = max(max_len, right - left + 1)

        return max_len