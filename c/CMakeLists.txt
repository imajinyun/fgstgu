cmake_minimum_required(VERSION 3.17)
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

# Add executable for all directory
file(GLOB_RECURSE files ${PROJECT_SOURCE_DIR}/*/main*.c)
foreach (file ${files})
    string(REGEX REPLACE ".+/(.+)\\..*" "\\1" name ${file})
    string(REGEX MATCH "src/(.+)/[a-z].+/" path ${file})
    string(SUBSTRING ${path} 4 -1 path)
    string(REPLACE "/" "." path ${path})
    string(CONCAT out ${path} ${name})
    add_executable(${out} ${file})
    message(\ \ --\ ${file}\ will\ be\ compiled\ to\ ${PROJECT_BINARY_DIR}/${out})
endforeach ()
