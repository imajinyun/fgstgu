#include <stdio.h>

#define MSG "I am a symbolic string constant."
#define MAXLENGTH 81

int main(int argc, char *argv[]) {
  char words[MAXLENGTH] = "I am a string in an array.";
  const char *p = "Something is pointing at me.";
  puts("Here are some strings:");
  puts(MSG);
  puts(words);
  puts(p);
  words[8] = 'p';
  puts(words);
  return 0;
}
