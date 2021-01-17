cmake_minimum_required(VERSION 3.17)
project(c C)

set(CMAKE_C_STANDARD 11)

# Configure Source & Binary Directories
set(PROJECT_ROOT "${PROJECT_SOURCE_DIR}")
set(PROJECT_SOURCE_DIR "${PROJECT_SOURCE_DIR}/src")
set(PROJECT_BINARY_DIR "${PROJECT_BINARY_DIR}/bin")
set(EXECUTABLE_OUTPUT_PATH ${PROJECT_BINARY_DIR})

message("Source path: ${PROJECT_SOURCE_DIR}")
message("Binary path: ${PROJECT_BINARY_DIR}")

# Add executable for a directory
add_executable(a.hello-world.main ${PROJECT_SOURCE_DIR}/a/hello-world/main.c)
add_executable(a.array.a10.main ${PROJECT_SOURCE_DIR}/a/array/a10/main.c)
add_executable(a.pointer.a01.main ${PROJECT_SOURCE_DIR}/a/pointer/a01/main.c)

# Add executable for b directory
add_executable(b.student.main ${PROJECT_SOURCE_DIR}/b/student/main.c)

# Add executable for c directory
