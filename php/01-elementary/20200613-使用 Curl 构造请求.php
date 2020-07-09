<?php

$data = ['phone' => 13600000000, 'type' => 'password'];
print_r(request_by_curl('http://apicorefz.xiaohe.com/api/member/image/captchas', $data));
echo PHP_EOL;
print_r(request_by_curl('https://fanyi.sogou.com/', '?keyword=test&transfrom=auto&transto=zh-CHS'));


function request_by_curl(string $url, $data = null)
{
    $ch = curl_init($url);
    $options = [
        CURLOPT_SSL_VERIFYHOST => false,
        CURLOPT_SSL_VERIFYPEER => false,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_HEADER => ['Content-Type: application/json;charset=utf-8'],
        CURLOPT_VERBOSE => 0,
        CURLOPT_TIMEOUT => 5,
    ];

    foreach ($options as $key => $val) {
        curl_setopt($ch, $key, $val);
    }

    if (is_array($data)) {
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
    } elseif (is_string($data)) {
        curl_setopt($ch, CURLOPT_URL, $url . $data);
    }
    $result = curl_exec($ch);
    curl_close($ch);

    return $result;
}
