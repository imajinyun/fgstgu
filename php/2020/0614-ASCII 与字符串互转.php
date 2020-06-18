<?php

/*
97 - a
98 - b
99 - c
100 - d
101 - e
102 - f
103 - g
104 - h
105 - i
106 - j
107 - k
108 - l
109 - m
110 - n
111 - o
112 - p
113 - q
114 - r
115 - s
116 - t
117 - u
118 - v
119 - w
120 - x
121 - y
122 - z
*/
for ($i = 97; $i <= 122; $i++) {
    $char = chr($i);
    echo "{$i} - {$char}\n";
}

/*
A - 48
B - 49
C - 50
D - 51
E - 52
F - 53
G - 54
H - 55
I - 56
J - 57
K - 49
L - 49
M - 49
N - 49
O - 49
P - 49
Q - 49
R - 49
S - 49
T - 49
U - 50
V - 50
W - 50
X - 50
Y - 50
Z - 50
*/
$s = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
for ($i = 0; $i < strlen($s); $i++) {
    $ascii = ord($i);
    echo "{$s[$i]} - {$ascii}\n";
}
