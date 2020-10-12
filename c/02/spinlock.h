#ifndef C_EXERCISES_SPINLOCK_H
#define C_EXERCISES_SPINLOCK_H

struct spinlock {
  uint locked;

  char *name;
  struct cpu *cpu;
  uint pcs[10];
};

#endif //C_EXERCISES_SPINLOCK_H
