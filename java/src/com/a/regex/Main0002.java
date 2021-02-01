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

      // 相对于 && 的交运算，Java 也支持在字符组内用 || 进行并运算，不过因为并集的意思就是『给字符组内部添加字符』，所以并不需要设定特殊的运算符
      "3".matches("[[0-4][6-9]]"), // true
      "5".matches("[[0-4][6-9]]"), // false
      "3".matches("[[0-4]]"), // true
      "[".matches("[[0-4]]"), // false

      // Java 中可以使用 Unicode 码值
      "我".matches("[\u4E00-\u9FFF]"), // true
      "中国".matches("[\u4E00-\u9FFF]"), // false,
      "🇨🇳".matches("[\u4E00-\u9FFF]"), // false,

      // POSIX 字符组都只能匹配 ASCII 字符
      "0".matches("\\p{Digit}"), // true
      "f".matches("\\p{XDigit}"), // true

      // Java 对 Unicode 字符组的支持比较好，它支持 Unicode Property，以下实例为全角下输入
      "１".matches("\\p{N}"), // true
      "，".matches("\\p{P}"), // true

      // Java 也支持 Unicode Block，其记法是以 In 为前缀的，比如匹配中文字符的 Unicode Block 就应当写为 InCJK_Compatibility_Ideographs
      "🎉".matches("\\p{InCJK_COMPATIBILITY_IDEOGRAPHS}"), // false
    };
    for (int i = 0; i < items.length; i++) {
      System.out.println(items[i]);
    }
  }

}
