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
    @Test
    public void testLeetCode04(){
        LeetCode04 l = new LeetCode04();
        int[] nums1 = {1,2};
        int[] nums2 = {1,2,3,4,5,6,7,8,9,10};
        double res = l.findMedianSortedArrays1(nums1,nums2);

        //StringBuffer sb = new StringBuffer("wke");
        System.out.println(res);

    }
    @Test
    public void testLeetCode05(){
        LeetCode05 l = new LeetCode05();
        l.longestPalindrome("cbadabc");

    }
}
