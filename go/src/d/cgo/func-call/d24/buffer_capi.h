typedef struct Buffer_T Buffer_T;

Buffer_T *NewBuffer(int size);
void DeleteBuffer(Buffer_T *p);
char *BufferData(Buffer_T *p);
int BufferSize(Buffer_T *p);
