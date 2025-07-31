-- データベース内のテーブル一覧を確認
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public' 
ORDER BY table_name;

-- 各テーブルの構造を確認
\d teacher_tests
\d courses
\d subjects
\d users 