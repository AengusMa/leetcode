package leetcode003;

import org.junit.Test;

/**
 * @author mawenlong
 * @date 2019/02/11
 */
public class LeetCode03Test {

    @Test
    public void test01(){
        LeetCode03 code03 = new LeetCode03();
        System.out.println(code03.lengthOfLongestSubstring("abcabcbb"));
        System.out.println(code03.lengthOfLongestSubstring("bbbbb"));
        System.out.println(code03.lengthOfLongestSubstring("pwwkew"));
    }
}
