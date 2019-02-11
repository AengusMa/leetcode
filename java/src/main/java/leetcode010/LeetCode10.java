package leetcode010;

/**
 * @author mawenlong
 * @date 2018/11/21
 */
public class LeetCode10 {
    // 若p为空----------------->返回s.isEmpty();
    // 若p长度为1-------------->判断首字符是否相等(也可以去除第一个字符递归调用)
    // 若 p的第二个字符不为*---->判断首字符是否相等，去除第一个字符递归调用
    // 若 p的第二个字符为*------>递归
    public boolean isMatch(String s, String p) {
        if (p.isEmpty()) {
            return s.isEmpty();
        }
        if (p.length() == 1) {
            return !s.isEmpty() && s.length() == 1 && (s.charAt(0) == p.charAt(0) || p.charAt(0) == '.');
        }
        if (p.charAt(1) != '*') {
            return !s.isEmpty() && (s.charAt(0) == p.charAt(0) || p.charAt(0) == '.')
                    && isMatch(s.substring(1), p.substring(1));
        }
        while (!s.isEmpty() && (s.charAt(0) == p.charAt(0) || p.charAt(0) == '.')) {
            // s="aaaaab",p="a*b"
            // p后面的如果匹配字符串返回true，如果不匹配循环
            if (isMatch(s, p.substring(2))) {
                return true;
            }
            s = s.substring(1);
        }
        return isMatch(s, p.substring(2));

    }

    public boolean isMatch1(String s, String p) {
        if (p.isEmpty()) {
            return s.isEmpty();
        }
        if (p.length() > 1 && p.charAt(1) == '*') {
            return isMatch(s, p.substring(2)) && (p.charAt(0) == s.charAt(0) || p.charAt(0) == '.')
                    && isMatch(s.substring(1), p.substring(1));
        } else {
            // p长度为1或者第二个字母不为*
            return !s.isEmpty() && (p.charAt(0) == s.charAt(0) || p.charAt(0) == '.')
                    && isMatch(s.substring(1), p.substring(1));
        }
    }

    public boolean isMatchDP(String s, String p) {
        if (s == null || p == null) {
            return false;
        }
        boolean dp[][] = new boolean[s.length() + 1][p.length() + 1];
        dp[0][0] = true;
        // 初始化第0列
        for (int i = 0; i < p.length(); i++) {
            if (p.charAt(i) == '*' && dp[0][i - 1]) {
                dp[0][i + 1] = true;
            }
        }
        for (int i = 1; i <= s.length(); i++) {
            for (int j = 1; j <= p.length(); j++)
                if (p.charAt(j - 1) == '*')
                    // *对应空或者当前字符匹配*之前的字符 ，才能传递dp[i-1][j]真值
                    dp[i][j] = dp[i][j - 2] || (s.charAt(i - 1) == p.charAt(j - 2) || p.charAt(j - 2) == '.')
                            && dp[i - 1][j];
                else
                    // 当前字符完全匹配，才能传递dp[i-1][j-1] 真值
                    dp[i][j] = (s.charAt(i - 1) == p.charAt(j - 1) || p.charAt(j - 1) == '.')
                            && dp[i - 1][j - 1];
        }
        return dp[s.length()][p.length()];
    }

}
