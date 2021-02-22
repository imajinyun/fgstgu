#include "name_st.h"
#include <stdio.h>

int main(int argc, char *argv[]) {
  names myname;
  get_names(&myname);
  printf("Let's welcome ");
  show_names(&myname);
  printf(" to this program.\n");
  return 0;
}
