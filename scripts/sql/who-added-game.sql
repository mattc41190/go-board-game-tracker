SELECT g.title, u.email 
FROM games AS g 
JOIN users AS u 
ON g.user_id = u.id;