package com.mwl.leetcode;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * @author mawenlong
 * @date 2018/11/20
 */
public class LeetCode08 {
    public int myAtoi(String str) {
        String pa = "[ ]*[+,-]?\\d+";
        Pattern pattern = Pattern.compile(pa);
        Matcher matcher = pattern.matcher(str);
        if (matcher.find() && matcher.start() == 0) {
            String num = matcher.group();
            try {
                return Integer.valueOf(num.trim());
            } catch (Exception e) {
                return num.contains("-") ? Integer.MIN_VALUE : Integer.MAX_VALUE;
            }
        } else {
            return 0;
        }
    }
}
