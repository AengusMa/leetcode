package com.mwl.leetcode;

/**
 * @author mawenlong
 * @date 2018/9/2
 * describe:
 */
public class LeetCode07 {
    public String convert(String s, int numRows) {
        if (numRows == 1) {
            return s;
        }
        int len = s.length();
        int rowcy = numRows * 2 - 2;
        StringBuffer sb = new StringBuffer();
        for (int i = 0; i < numRows; i++) {
            for (int j = 0; j + i < len; j += rowcy) {
                sb.append(s.charAt(j + i));
                if (i != 0 && i != numRows - 1 && j + rowcy - i < len) {
                    sb.append(s.charAt(j + rowcy - i));
                }
            }

        }
        return sb.toString();
    }
}