#include "./buffer.h"

extern "C" {
#include "./buffer_capi.h"
}

struct Buffer_T : Buffer {
  Buffer_T(int size) : Buffer(size) {}
  ~Buffer_T() {}
};

Buffer_T *NewBuffer(int size) {
  auto p = new Buffer_T(size);
  return p;
}

void DeleteBuffer(Buffer_T *p) { delete p; }

char *BufferData(Buffer_T *p) { return p->Data(); }

int BufferSize(Buffer_T *p) { return p->Size(); }
