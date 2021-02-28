#include <stdio.h>

int main(int argc, char *argv[]) {
  char plaintext = 'M';                     // 明文
  char secretkey = '#';                     // 密钥
  char ciphertext = plaintext ^ secretkey;  // 加密后的密文
  char decodetext = ciphertext ^ secretkey; // 解密后的明文
  char buffer[20];

  sprintf(buffer, "%#X", plaintext);
  printf("plaintext:  %c %s\n", plaintext, buffer);

  sprintf(buffer, "%#X", secretkey);
  printf("secretkey:  %c %s\n", secretkey, buffer);

  sprintf(buffer, "%#X", ciphertext);
  printf("ciphertext: %c %s\n", ciphertext, buffer);

  sprintf(buffer, "%#X", decodetext);
  printf("decodetext: %c %s\n", decodetext, buffer);

  printf("\n");
  puts("对称加密的关键技术：\n"
       "1. 通过一次异或运算，生成密文，密文没有可读性，与原文风马牛不相及，这就是加密；\n"
       "2. 密文再经过一次异或运算，就会还原成原文，这就是解密的过程；\n"
       "3. 加密和解密需要相同的密钥，如果密钥不对，是无法成功解密的；\n");
  puts("如果加密和解密的密钥不同，则称为非对称加密算法。在非对称算法中，加密的密钥称为公钥，解密的密钥称为私钥，只知道"
       "公钥是无法解密的，还必须知道私钥");
  return 0;
}
