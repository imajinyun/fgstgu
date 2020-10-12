#include "defs.h"
#include "type.h"
#include "param.h"
#include "memlayout.h"
#include "mmu.h"
#include "spinlock.h"

struct {
  struct spinlock lock;
  int use_lock;
  struct run *freelist;
} kmem;

void kinit1(void *vstart, void *vend) {
  initlock(&kmem.lock, "kmem");
  kmem.use_lock = 0;
  freerange(vstart, vend);
}
