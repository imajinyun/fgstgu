#include <assert.h>
#include <errno.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_DATA 512
#define MAX_ROWS 100

struct Student {
  int id;
  bool set;
  char name[MAX_DATA];
  char email[MAX_DATA];
};

struct Database {
  struct Student rows[MAX_ROWS];
};

struct Connection {
  FILE *file;
  struct Database *db;
};

void die(const char *message) {
  if (errno) {
    perror(message);
  } else {
    printf("ERROR: %s\n", message);
  }
  exit(EXIT_FAILURE);
}

void showStudent(struct Student *stu) { printf("Id: %d, Name: %s, Email: %s\n", stu->id, stu->name, stu->email); }

void dbLoad(struct Connection *conn) {
  int row = fread(conn->db, sizeof(struct Database), 1, conn->file);
  if (row != 1) { die("Failed to load database"); }
}

struct Connection *dbOpen(const char *filename, char mode) {
  struct Connection *conn = malloc(sizeof(struct Connection));
  if (!conn) { die("Memory error"); }
  conn->db = malloc(sizeof(struct Database));
  if (!conn->db) { die("Memory error"); }
  if (mode == 'c') {
    conn->file = fopen(filename, "w");
  } else {
    conn->file = fopen(filename, "r+");
    if (conn->file) { dbLoad(conn); }
  }

  if (!conn->file) { die("Failed to open the file"); }
  return conn;
}

void dbClose(struct Connection *conn) {
  if (conn) {
    if (conn->file) { fclose(conn->file); }
    if (conn->db) { free(conn->db); }
    free(conn);
  }
}

void dbWrite(struct Connection *conn) {
  rewind(conn->file);
  int row = fwrite(conn->db, sizeof(struct Database), 1, conn->file);
  if (row != 1) { die("Failed to write database"); }
  row = fflush(conn->file);
  if (row == -1) { die("Connot flush database"); }
}

void dbCreate(struct Connection *conn) {
  for (int i = 0; i < MAX_ROWS; i++) {
    // make a prototype to initialize it
    struct Student stu = {.id = i, .set = false};
    // the just assign it
    conn->db->rows[i] = stu;
  }
}

void dbSet(struct Connection *conn, int id, const char *name, const char *email) {
  struct Student *stu = &conn->db->rows[id];
  if (stu->set) { die("Already set, delete it first"); }
  stu->set = true;
  char *res = strncpy(stu->name, name, MAX_DATA);
  if (!res) { die("Name copy failed"); }
  res = strncpy(stu->email, email, MAX_DATA);
  if (!res) { die("Email copy failed"); }
}

void dbGet(struct Connection *conn, int id) {
  struct Student *stu = &conn->db->rows[id];
  if (stu->set) {
    showStudent(stu);
  } else {
    die("Id is not set");
  }
}

void dbDelete(struct Connection *conn, int id) {
  struct Student stu = {.id = id, .set = false};
  conn->db->rows[id] = stu;
}

void dbList(struct Connection *conn) {
  struct Database *db = conn->db;
  for (int i = 0; i < MAX_ROWS; i++) {
    struct Student *stu = &db->rows[i];
    if (stu->set) { showStudent(stu); }
  }
}

/**
 * 内存分配之堆和栈。
 *
 * -> 如果变量不是用 malloc 直接获取的，也不是在函数内通过 malloc 间接获取的，那么这个变量就是在栈上的。
 *
 * -> /path/to/a.store.0009 /tmp/db.dat c
 * -> /path/to/a.store.0009 /tmp/db.dat s 1 jack jack@gmail.com
 * -> /path/to/a.store.0009 /tmp/db.dat l
 * -> /path/to/a.store.0009 /tmp/db.dat g 1
 * -> /path/to/a.store.0009 /tmp/db.dat d 1
 * -> /path/to/a.store.0009 /tmp/db.dat l
 */
int main(int argc, char *argv[]) {
  if (argc < 3) { die("Usage: /path/to/a.store.0009 <dbfile> <action> [action parameters]"); }
  char *filename = argv[1];
  char action = argv[2][0];
  struct Connection *conn = dbOpen(filename, action);
  int id = 0;
  if (argc > 3) { id = atoi(argv[3]); }
  if (id >= MAX_ROWS) { die("There's not that many records"); }
  switch (action) {
    case 'c':
      dbCreate(conn);
      dbWrite(conn);
      break;
    case 'g':
      if (argc != 4) { die("Need an id to get"); }
      dbGet(conn, id);
      break;
    case 's':
      if (argc != 6) { die("Need id, name, email to set"); }
      dbSet(conn, id, argv[4], argv[5]);
      dbWrite(conn);
      break;
    case 'd':
      if (argc != 4) { die("Need id to delete"); }
      dbDelete(conn, id);
      dbWrite(conn);
      break;
    case 'l':
      dbList(conn);
      break;
    default:
      die("Invalid action: c=create, g=get, s=set, d=delete, l=list");
  }
  dbClose(conn);
  return 0;
}
