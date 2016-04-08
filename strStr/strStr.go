// Source : https://oj.leetcode.com/problems/implement-strstr/
// Author : daoran
// Date   : 2016-04-08

/********************************************************************************** 
* 
* Implement strStr().
* Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
* 
*               
**********************************************************************************/

func strStr(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}