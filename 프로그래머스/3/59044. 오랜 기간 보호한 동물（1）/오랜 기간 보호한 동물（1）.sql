SELECT a.name, a.DATETIME
FROM animal_ins a 
left outer join animal_outs b on a.animal_id =b.`ANIMAL_ID`
WHERE b.animal_id is NULL
ORDER BY a.`DATETIME`
limit 0,3;