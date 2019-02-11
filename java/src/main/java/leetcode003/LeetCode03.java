package leetcode003;

import java.util.HashMap;
import java.util.Map;

/**
 * @author mawenlong
 * @date 2018/08/10
 * describe:
 * Given a string, find the length of the longest substring without repeating characters.
 * Examples:
 * Given "abcabcbb", the answer is "abc", which the length is 3.
 * Given "bbbbb", the answer is "b", with the length of 1.
 * Given "pwwkew", the answer is "wke", with the length of 3.
 * Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */
public class LeetCode03 {
    public int lengthOfLongestSubstring(String s) {
        int res = 0;
        StringBuffer sb = new StringBuffer("");
        if (s.length() == 0) {
            return 0;
        }
        for (int i = 0; i < s.length(); i++) {
            int index = sb.indexOf(s.charAt(i) + "");
            if (index == -1) {
                sb.append(s.charAt(i));
                if (res < sb.length()) {
                    res = sb.length();
                }
            } else {
                sb = new StringBuffer(sb.subSequence(index + 1, sb.length()));
                sb.append(s.charAt(i));
            }
        }
        return res;
    }
    public int lengthOfLongestSubstring1(String s) {
        if (s.length() == 0) {
            return 0;
        }
        int max = 0;
        //map的key是字符，value为字符的位置。map里面的所有key不重复
        Map<Character, Integer> map = new HashMap<Character, Integer>();
        for (int i = 0, j = 0; i < s.length(); i++) {
            if(map.containsKey(s.charAt(i))){
                //j = map.get(s.charAt(i))+1是错误的，例子：abcd b a
                j = Math.max(j,map.get(s.charAt(i))+1);
            }
            map.put(s.charAt(i),i);
            max=Math.max(max,i-j+1);
        }
        return max;
    }
}
