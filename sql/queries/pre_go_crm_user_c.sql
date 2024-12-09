-- name: GetUserByEmailSQLC :one
SELECT usr_email, usr_id 
FROM `pre_go_crm_user_c` 
WHERE usr_email = ? 
LIMIT 1;

-- name: UpdateUserStatusByUserId :exec
UPDATE `pre_go_crm_user_c`
SET usr_status = $2,
    usr_updated_at = $3
WHERE usr_id = $1;
