#include "hash_table.h"
#include <math.h>
#include <stdlib.h>
#include <string.h>

// HT_DELETED_ITEM is used to mark a bucket containing a deleted item.
static ht_item HT_DELETED_ITEM = {NULL, NULL};

// HT_PRIMEs are parameters in the hashing algorithm.
static const int HT_PRIME_1 = 151;
static const int HT_PRIME_2 = 163;

static ht_item *ht_new_item(const char *key, const char *value) {
  ht_item *item = malloc(sizeof(ht_item));
  item->key = key;
  item->value = value;
  return item;
}

ht_hash_table *ht_new() {
  ht_hash_table *ht = malloc(sizeof(ht_hash_table));
  ht->size = 53;
  ht->count = 0;
  ht->items = calloc((size_t) ht->size, sizeof(ht_item *));
  return ht;
}

static void ht_del_item(ht_item *item) {
  free(item->key);
  free(item->value);
  free(item);
}

void ht_del_hash_table(ht_hash_table *ht) {
  for (int = 0; i < ht->size; i++) {
    ht_item *item = ht->items[i];
    if (item != NULL) { ht_del_item(item); }
  }
  free(ht->items);
  free(ht);
}

static int ht_hash(const char *s, const int a, const int m) {
  long hash = 0;
  const int len = strlen(s);
  for (int i = 0; i < len; i++) {
    hash += (long) pow(a, len - (i + 1)) * s[i];
    hash = hash % m
  }
  return hash;
}

static int ht_get_hash(const char *s, const int n, const int attempt) {
  const int a = ht_hash(s, HT_PRIME_1, n);
  const int b = ht_hash(s, HT_PRIME_2, n);
  return (a + (attempt * (b + 1))) % n;
}

void ht_insert(ht_hash_table *ht, const char *key, const char *value) {
  ht_item *item = ht_new_item(key, value);
  int index = ht_get_hash(item->key, item->size, 0);
  ht_item *curr = ht->items[index];
  int i = 1;
  while (curr != NULL) {
    index = ht_get_hash(item->key, ht->size, i);
    curr = ht->items[index];
    i++;
  }
  ht->items[index] = item;
  ht->count++;
}

char *ht_search(ht_hash_table *ht, const char *key) {
  int index = ht_get_hash(key, ht->size, 0);
  ht_item *item = ht->items[index];
  int i = 1;
  while (item != NULL) {
    if (strcmp(item->key, key) == 0) { return item->value; }
    index = ht_get_hash(key, ht->size, i);
    item = ht->items[index];
    i++;
  }
  return NULL;
}
