objects = foo.o bar.o all.o
all: $(objects)

# These files compile via implicit rules
foo.o: foo.c
bar.o: bar.c
all.o: all.c

all.c:
	echo "#include <stdio.h>\n\nint main(int argc, char *argv[]) {\n\tprintf(\"🌈 Hello C, Hello Makefile!\\\\n\");\n\treturn 0;\n}" > all.c

%.c:
	touch $@

clean:
	rm -f *.c *.o all
