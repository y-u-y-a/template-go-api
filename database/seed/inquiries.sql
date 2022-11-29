BEGIN;

INSERT INTO inquiries (id, created_at, updated_at, name, email, tel, content, is_confirm)
VALUES
(1, NOW(), NOW(), 'visitor01', 'inquiry01@gmail.com', '080-1234-5678', 'ログインできません。', false),
(2, NOW(), NOW(), 'visitor02', 'inquiry02@gmail.com', '080-1234-5678', 'ログインできません。', false),
(3, NOW(), NOW(), 'visitor03', 'inquiry03@gmail.com', '080-1234-5678', 'ログインできません。', false),
(4, NOW(), NOW(), 'visitor04', 'inquiry04@gmail.com', '080-1234-5678', 'ログインできません。', false),
(5, NOW(), NOW(), 'visitor05', 'inquiry05@gmail.com', '080-1234-5678', 'ログインできません。', false),
(6, NOW(), NOW(), 'visitor06', 'inquiry06@gmail.com', '080-1234-5678', 'ログインできません。', false),
(7, NOW(), NOW(), 'visitor07', 'inquiry07@gmail.com', '080-1234-5678', 'ログインできません。', false);
COMMIT;