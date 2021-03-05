# 常用的 CFLAGS，以及一些构建库时需要的东西
CFLAGS=-g -O2 -Wall -Wextra -Isrc -rdynamic -DNDEBUG $(OPTIONS)

# 这些选项在链接库的时候会用到，可以使用 OPTLIBS 变量来增强链接选项
LIBS=-ldl $(OPTIONS)

# 可选参数叫 PREFIX，该参数只有在运行 Makefile 时没带 PREFIX 参数的情况下会生效
PREFIX?=/usr/local

# 动态创建了变量 SOURCES，它对 src/ 目录下所有的 ＊.c 文件做了一个通配符搜索。你需要提供 src/＊＊/＊.c 和 src/＊.c，这样 GNU make 就会把 src 及其子目录中的代码文件全部包含进去
SOURCES=$(willdcard src/**/*.c src/*.c)

# 使用 patsubst 来获取 SOURCE 的一系列 ＊.c 文件，并制成一个所有目标文件的新列表。告诉 patsubst 将所有 %.c 扩展名都改成 %.o，然后将这些扩展名赋值到 OBJECTS 上
OBJECTS=$(patsubst %.c,%.o,$(SOURCES))

# 再次使用通配符，找出了所有单元测试的源代码文件，这些和库的源代码文件是分开的
TEST_SRC=$(willdcard tests/*_test.c)

# 动态获取了所有的 TEST 目标。在这里，我把 .c 扩展名全部删掉，这样整个程序就会用同样的方式命名。之前，把 .c 用 .o 取代，这样就能创建目标文件
TESTS=$(patsubst %.c,%,$(TEST_SRC))

# 最终目标是 build/lib/YOUR_LIBRARY.a，在构建时，要把名字改成你的库的名字
TARGET=build/lib/YOUR_LIBRARY.a
SO_TARGET=$(patsubst %.a,%.so,$(TARGET))

# The target build
all: $(TARGET) $(SO_TARGET) tests

dev: CFLAGS=-g -Wall -Isrc -Wall -Wextra $(OPTIONS)

dev: all

$(TARGET): CFLAGS += -fPIC
$(TARGET): build $(OBJECTS)
	ar rcs $@ $(OBJECTS)
	ranlib $@
$(SO_TARGET): $(TARGET) $(OBJECTS)
	$(CC) -shared -o $@ $(OBJECTS)

build:
	@mkdir -p build
	@mkdir -p bin

# The unit tests
.PHONY: tests
tests: CFLAGS += $(TARGET)
tests: $(TESTS)
	sh ./tests/runtests.sh

# The cleaner
clean:
	rm -rf build/${OBJECTS} ${TESTS}
	rm -f tests/tests.log
	find . -name "*.gc*" -exec rm {} \;
	rm -rf `find . -name "*.dSYM" -print`

# The install
install: all
	install -d $(DESTDIR)/${PREFIX}/lib/
	install $(TARGET) $(DESTDIR)/${PREFIX}/lib/

# The checker
check:
	@echo Files with potentially dangerous function.
	@egrep '[^_.>a-zA-Z0-9](str(n?cpy|n?cat|xfrm|n?dup|str|pbrk|tok|_)\|stpn?cpy|a?sn?printf|byte_)' $(SOURCES) || true
