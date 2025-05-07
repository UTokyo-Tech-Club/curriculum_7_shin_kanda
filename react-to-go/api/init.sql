-- テーブルが存在する場合は削除
DROP TABLE IF EXISTS users;

-- テーブルの作成
CREATE TABLE users (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INTEGER NOT NULL
);

-- サンプルデータの挿入
INSERT INTO users (id, name, age) VALUES
    ('01HXSAXGZJ8K4P2Q3R5T7V9WXZ', '太郎', 25),
    ('01HXSAXGZJ8K4P2Q3R5T7V9WXY', '花子', 30);