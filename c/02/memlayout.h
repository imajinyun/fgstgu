// Memory layout

#define EXTMEM 0x100000     // Start of extended memory
#define PHYSTOP 0xE000000   // Top physical memory
#define DEVSPACE 0xFE000000 // Other devices are at high address

// Key address for address space layout(see kmap in vm.c for layout)
#define KERNBASE 0x80000000          // First kernel virtual address
#define KERNLINK (KERNBASE + EXTMEM) // Address where kernel is linked
#define V2P(a) (((uint)(a)) - KERNBASE)
#define P2V(a) ((void *) (((char *) (a)) + KERNBASE))
#define V2P_WO(x) ((x) -KERNBASE)
#define P2V_WO(x) ((x) + KERNBASE)
