# Last updated: 4/23/2025, 12:15:22 PM
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        longest_substring = ''
        working_strings = []

        # Iterate through the input string.
        for char in s:
            new_working_strings = []

            for idx, candidate in enumerate(working_strings):
                # Check for a repeat - this is invalid.
                if char in candidate:
                    # Before we discard it, check if it's longest.
                    if len(candidate) > len(longest_substring):
                        longest_substring = candidate
                # Otherwise we have a non-repeating string to continue working on.
                else:
                    new_working_strings.append(concat(candidate, char))

            # Always add the character as the base of a new working string.
            new_working_strings.append(char)

            working_strings = new_working_strings

        # Check if any of our working strings after the final iteration are longest.
        for s in working_strings:
            if len(s) > len(longest_substring):
                longest_substring = s

        return(len(longest_substring))