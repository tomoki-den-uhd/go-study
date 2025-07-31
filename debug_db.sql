-- データベース内のテーブル一覧を確認
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public' 
ORDER BY table_name;

-- usersテーブルの構造確認
\d users;

-- teacher_testsテーブルの構造確認
\d teacher_tests;

-- coursesテーブルの構造確認
\d courses;

-- サンプルデータの確認
SELECT user_id, name, role FROM users LIMIT 5;
SELECT teacher_test_id, title, created_by, course_id FROM teacher_tests LIMIT 5; 