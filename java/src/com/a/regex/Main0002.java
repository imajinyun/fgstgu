package com.a.regex;

public class Main0002 {

    public static void main(String[] args) {
        // Java 语言中所有的字符都采用 UTF-16 编码，所以在字符组中可以直接使用中文，因而不会出现多字节字符错误匹配的问题。

        // -> true
        System.out.println("你".matches("."));

        // -> false
        System.out.println("遭".matches("[正则]"));
    }

}
