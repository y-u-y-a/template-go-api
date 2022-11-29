BEGIN;

INSERT INTO users (id, created_at, updated_at, given_name, family_name, email, birthday, gender, company_id)
VALUES
-- company_id=1
('gwernowerhfo11', NOW(), NOW(), 'taro01', 'test01', 'test01@gmail.com', '1990-01-01', 3, 1),
('gwernowerhfo22', NOW(), NOW(), 'taro02', 'test02', 'test02@gmail.com', '1990-02-01', 3, 1),
('gwernowerhfo33', NOW(), NOW(), 'taro03', 'test03', 'test03@gmail.com', '1990-03-01', 3, 1),
('gwernowerhfo44', NOW(), NOW(), 'taro04', 'test04', 'test04@gmail.com', '1990-04-01', 3, 1),
('gwernowerhfo55', NOW(), NOW(), 'taro05', 'test05', 'test05@gmail.com', '1990-05-01', 3, 1),
('gwernowerhfo66', NOW(), NOW(), 'taro06', 'test06', 'test06@gmail.com', '1990-06-01', 3, 1),
('gwernowerhfo77', NOW(), NOW(), 'taro07', 'test07', 'test07@gmail.com', '1990-07-01', 3, 1);
COMMIT;