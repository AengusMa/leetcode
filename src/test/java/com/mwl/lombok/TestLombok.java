package com.mwl.lombok;

import java.util.Map;

/**
 * @author mawenlong
 * @date 2018/08/10
 * describe:
 */
public class TestLombok {
    public static void main(String[] args) {
        Student s = new Student();
        s.setAge(3);
        System.out.println(s.getAge()+s.toString());
    }
}
