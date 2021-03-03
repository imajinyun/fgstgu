cmake_minimum_required(VERSION 3.18)
project(c C)

set(CMAKE_C_STANDARD 11)

# Version
set(VERSION_MAJOR 1)
set(VERSION_MINOR 0)
set(VERSION_FEATURE 0)
set(VERSION_PATCH 0)
set(VERSION "${VERSION_MAJOR}.${VERSION_MINOR}")
set(VERSION "${VERSION}.${VERSION_FEATURE}.${VERSION_PATCH}")
message("Version: ${VERSION}")

# Configure Source & Binary Directories
set(PROJECT_ROOT "${PROJECT_SOURCE_DIR}")
set(PROJECT_SOURCE_DIR "${PROJECT_SOURCE_DIR}/src")
set(PROJECT_BINARY_DIR "${PROJECT_BINARY_DIR}/bin")
set(EXECUTABLE_OUTPUT_PATH ${PROJECT_BINARY_DIR})

message("Source path: ${PROJECT_SOURCE_DIR}")
message("Binary path: ${PROJECT_BINARY_DIR}")

add_executable(a.ds.extra.linkedlist ${PROJECT_SOURCE_DIR}/a/ds/extra/linkedlist.c ${PROJECT_SOURCE_DIR}/a/ds/extra/linkedlist.c)
add_executable(a.ds.extra.queue ${PROJECT_SOURCE_DIR}/a/ds/extra/queue.c ${PROJECT_SOURCE_DIR}/a/ds/extra/queue.c)
add_executable(a.preprocessor.extra.name ${PROJECT_SOURCE_DIR}/a/preprocessor/extra/name_main.c ${PROJECT_SOURCE_DIR}/a/preprocessor/extra/name_st.c)

# Add executable for all directory
file(GLOB_RECURSE files "${PROJECT_SOURCE_DIR}/*/[0-9]*.c")
foreach (file ${files})
    string(REGEX MATCH "src/(.+)/[a-z].+/" path ${file})
    string(SUBSTRING ${path} 4 -1 path)
    string(REPLACE "/" "." path ${path})

    string(REGEX REPLACE ".+/(.+)\\..*" "\\1" name ${file})
    string(SUBSTRING ${name} 0 4 prefix)
    string(REGEX MATCH "[0-9].?$" suffix ${name})

    if (${suffix} MATCHES "[0-9].?")
        string(CONCAT out ${path} ${prefix} "." ${suffix})
    else ()
        string(CONCAT out ${path} ${prefix})
    endif ()
    add_executable(${out} ${file})
    message(\ \ --\ ${file}\ will\ be\ compiled\ to\ ${PROJECT_BINARY_DIR}/${out})
endforeach ()
