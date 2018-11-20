package com.mwl.leetcode;

/**
 * @author mawenlong
 * @date 2018/11/20
 */
public class LeetCode06 {
    public int reverse(int x) {
        long res = 0;
        if (x == 0) {
            return 0;
        }
        while (x != 0) {
            int tmp = x % 10;
            if (tmp != 0 || res != 0) {
                res = res * 10 + tmp;
            }
            x = x / 10;
        }
        if (res > Integer.MAX_VALUE) {
            return 0;
        }
        if (res < Integer.MIN_VALUE) {
            return 0;
        }
        return (int) res;
    }
}