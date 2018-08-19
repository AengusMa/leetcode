package com.mwl.leetcode;

import java.util.HashMap;
import java.util.Map;

/**
 * @author mawenlong
 * @date 2018/08/10
 * describe:
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
 * You may assume nums1 and nums2 cannot be both empty.
 * <p>
 * Example 1:
 * nums1 = [1, 3]
 * nums2 = [2]
 * The median is 2.0
 */
public class LeetCode04 {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.length;
        int n = nums2.length;
        int len = m + n;
        int left = -1, right = -1;
        int aStart = 0, bStart = 0;
        for (int i = 0; i <= len / 2; i++) {
            left = right;
            boolean flag = aStart < m && (bStart >= n || nums1[aStart] < nums2[bStart]);
            if (flag) {
                right = nums1[aStart++];
            } else {
                right = nums2[bStart++];
            }
        }
        if ((len & 1) == 0) {
            return (left + right) / 2.0;
        } else {
            return right;
        }

    }
}
