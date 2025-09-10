<?php
// MySQL connection
$pdo = new PDO(
    'mysql:host=' . getenv('DB_HOST') . ';dbname=' . getenv('DB_NAME'),
    getenv('DB_USER'),
    getenv('DB_PASSWORD')
);

// Redis connection
$redis = new Redis();
$redis->connect('redis', 6379);
?>