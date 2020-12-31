#include <stdint.h>

typedef uintptr_t person_t;
person_t newPerson(char *name, int age);
void deletePerson(person_t p);
void setPerson(person_t p, char *name, int age);
char *getPersonName(person_t p, char *buf, int size);
int getPersonAge(person_t p);
