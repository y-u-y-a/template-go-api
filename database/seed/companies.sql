BEGIN;

INSERT INTO companies (id, created_at, updated_at, name)
VALUES
(1, NOW(), NOW(), '株式会社ABCDE'),
(2, NOW(), NOW(), '株式会社AtoZ'),
(3, NOW(), NOW(), 'Amazon'),
(4, NOW(), NOW(), 'Apple'),
(5, NOW(), NOW(), 'Google'),
(6, NOW(), NOW(), 'Meta');
COMMIT;