#include "name_st.h"

#include "name_st.h" // 仅仅为了测试重复包含文件而正常编译通过运行而已
#include <stdio.h>

int main(int argc, char *argv[]) {
  names myname;
  get_names(&myname);
  printf("🎉 Let's welcome ");
  show_names(&myname);
  printf(" to this program.\n");
  return 0;
}
