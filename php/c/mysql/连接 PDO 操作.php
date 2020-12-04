<?php

/**
 * @link https://juejin.im/post/5b18a57d6fb9a01e82262010
 */
$dsn = 'mysql:host=127.0.0.1;port=3306;dbname=homestead';
$username = 'username';
$password = 'password';

try {
    $pdo = new \PDO($dsn, $username, $password);
    $pdo->setAttribute(\PDO::ATTR_ERRMODE, \PDO::ERRMODE_EXCEPTION);
} catch (\PDOException $e) {
    die($e->getMessage());
}

$rows = $pdo->query('SELECT * FROM mc_merchants LIMIT 5;');
foreach ($rows->fetchAll(\PDO::FETCH_ASSOC) as $row) {
    print_r($row);
}
