main: main.o
	cc main.o -o main # Runs third

main.o: main.c
	cc -c main.c -o main.o # Runs second

main.c:
	echo "#include <stdio.h>\n\nint main(int argc, char *argv[]) {\n\tprintf(\"🌈 Hello C, Hello Makefile!\\\\n\");\n\treturn 0;\n}" > main.c # Runs first

clean:
	rm -f main*
