package com.a.regex;

public class Main0001 {

    public static void main(String[] args) {
        // Java 语言中的正则表达式都是以字符串的形式给出的，所以字符串转义和正则表达式转义都必须考虑。

        // -> true
        System.out.println("1".matches("\\d"));

        // -> false
        System.out.println("1a".matches("\\d"));

        // -> true
        System.out.println("\\".matches("\\\\"));
    }

}
