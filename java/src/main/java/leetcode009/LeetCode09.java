package leetcode009;

/**
 * @author mawenlong
 * @date 2018/11/20
 */
public class LeetCode09 {
    public boolean isPalindrome(int x) {
        if (x == 0) {
            return true;
        }
        if (x < 0) {
            return false;
        }
        int tag = 0;
        int org = x;
        while (x != 0) {
            tag = tag * 10 + x % 10;
            x = x / 10;
        }
        if (tag == org) {
            return true;
        }
        return false;
    }
}
