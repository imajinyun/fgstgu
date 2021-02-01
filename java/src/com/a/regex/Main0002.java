package com.a.regex;

public class Main0002 {

  public static void main(String[] args) {
    // Java 语言中所有的字符都采用 UTF-16 编码，所以在字符组中可以直接使用中文，因而不会出现多字节字符错误匹配的问题。

    boolean items[] = {
      "你".matches("."), // true
      "遭".matches("[正则]"), // false

      // 匹配所有的小写辅音字母
      "a".matches("[[a-z]&&[^aeiou]]"), // false
      "b".matches("[[a-z]&&[^aeiou]]"), // true

      // 匹配英文大小写字母和数字（不包括下划线）
      "a".matches("[\\w&&[^_]]"), // true
      "0".matches("[\\w&&[^_]]"), // true
      "_".matches("[\\w&&[^_]]"), // false
    };
    for (int i = 0; i < items.length; i++) {
      System.out.println(items[i]);
    }
  }

}
