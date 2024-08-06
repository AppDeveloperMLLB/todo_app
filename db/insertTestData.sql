-- テストデータの投入

-- users テーブルにデータを投入
INSERT INTO users (google_id, email, username, profile_picture, created_at, updated_at) VALUES
('google-123', 'user1@example.com', 'user1', 'https://example.com/profile1.jpg', NOW(), NOW()),
('google-456', 'user2@example.com', 'user2', 'https://example.com/profile2.jpg', NOW(), NOW()),
('google-789', 'user3@example.com', 'user3', 'https://example.com/profile3.jpg', NOW(), NOW());

-- todos テーブルにデータを投入
INSERT INTO todos (user_id, title, description, status, created_at, updated_at) VALUES
(1, 'Todo1', 'Fiction', 'todo', NOW(), NOW()),
(1, 'Todo2', 'Non-Fiction', 'in_progress',  NOW(), NOW()),
(1, 'Todo3', 'Science Fiction', 'completed', NOW(), NOW());
