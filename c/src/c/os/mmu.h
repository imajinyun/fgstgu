// Eflags register
#define FL_IF 0x0000200 // Interrupt Enable

// Control register flags
#define CR0_PE 0x00000001 // Protection Enable
#define CR0_WP 0x00010000 // Write Protect
#define CR0_PG 0x80000000 // Paging

#define CR4_PSE 0x00000010 // Page size extension

// various segment selectors.
#define SEG_KCODE 1 // kernel code
#define SEG_KDATA 2 // kernel data+stack
#define SEG_UCODE 3 // user code
#define SEG_UDATA 4 // user data+stack
#define SEG_TSS 5 // this process's task state

// cpu->gdt[NSEGS] holds the above segments.
#define NSEGS 6

#ifndef __ASSEMBLER__
// Segment Descriptor
struct segdesc {
  uint lim_15_0: 16; // Low bits of segment limit
  uint base_15_0: 16; // Low bits of segment base address
  uint base_23_16: 8; // Middle bits of segment base address
  uint type: 4; // Segment type
  uint s: 1; // 0 = system, 1 = application
  uint dpl: 2; // Descriptor Privilege Level
  uint p: 1; // Present
  uint lim_19_16: 4; // High bits of segment limit
  uint avl: 1; // Unused (available for software use)
  uint rsv1: 1; // 0 = 16-bit segment, 1 = 32-bit segment
  uint db: 1;
  uint g: 1;
  uint base_31_24: 8;
};

// Normal segment
#define SEG(type, base, lim, dpl) (struct segdesc)    \
{ ((lim) >> 12) & 0xffff, (uint)(base) & 0xffff,      \
  ((uint)(base) >> 16) & 0xff, type, 1, dpl, 1,       \
  (uint)(lim) >> 28, 0, 0, 1, 1, (uint)(base) >> 24 }
#define SEG16(type, base, lim, dpl) (struct segdesc)  \
{ (lim) & 0xffff, (uint)(base) & 0xffff,              \
  ((uint)(base) >> 16) & 0xff, type, 1, dpl, 1,       \
  (uint)(lim) >> 16, 0, 0, 1, 0, (uint)(base) >> 24 }
#endif

#define DPL_USER 0x3 // User DPL

// Application segment type bits
#define STA_X 0x8
#define STA_W 0x2
#define STA_R 0x2

// System segment type bits
#define STS_T32A    0x9     // Available 32-bit TSS
#define STS_IG32    0xE     // 32-bit Interrupt Gate
#define STS_TG32    0xF     // 32-bit Trap Gate

// A virtual address 'la' has a three-part structure as follows:
//
// +--------10------+-------10-------+---------12----------+
// | Page Directory |   Page Table   | Offset within Page  |
// |      Index     |      Index     |                     |
// +----------------+----------------+---------------------+
//  \--- PDX(va) --/ \--- PTX(va) --/

// page directory index
#define PDX(va) (((uint)(va) >> PDXSHIFT) & 0x3FF)

// page table index
