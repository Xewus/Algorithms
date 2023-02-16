// (S1) Контест: простые задачи (SQL, 15 баллов)
// Полное решение: 15 баллов

SELECT p.id, p.contest_id, p.code
FROM problems AS p
JOIN submissions AS sm ON(sm.problem_id=p.id)
JOIN users AS u ON(u.id=sm.user_id)
WHERE sm.success=true
GROUP BY p.id
HAVING COUNT(DISTINCT u.id) > 1
ORDER BY p.id;
