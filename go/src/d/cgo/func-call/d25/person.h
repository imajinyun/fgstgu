extern "C" {
#include "./person_capi.h"
}

struct Person {
  person_t p_;
  Person(const char *name, int age) { this->p_ = newPerson((char *)name, age); }
  ~Person() { deletePerson(this->p_); }
  void Set(char *name, int age) { setPerson(this->p_, name, age); }
  char *GetName(char *buf, int size) {
    return getPersonName(this->p_, buf, size);
  }
  int GetAge() { return getPersonAge(this->p_); }
};
