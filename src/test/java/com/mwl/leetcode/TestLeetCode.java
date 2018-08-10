package com.mwl.leetcode;

import org.junit.Test;

/**
 * @author mawenlong
 * @date 2018/08/10
 * describe:
 */
public class TestLeetCode {
    @Test
    public void testLeetCode03(){
        LeetCode03 l = new LeetCode03();
        int res = l.lengthOfLongestSubstring1("abcabcbb");

        //StringBuffer sb = new StringBuffer("wke");
        System.out.println(res);
    }
}
