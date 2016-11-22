use db_sso;

select * from t_user;
select * from t_token;
select * from db_sso.t_email_verify;

INSERT INTO t_token(t_token_user_id,t_token_token) VALUES (2,"xxxxxxxx");

INSERT INTO t_user(t_user_username,t_user_password) VALUES ("loifsd","xxxxxxxx");

delete from t_user;
delete from t_token;
delete from t_email_verify;

SET SQL_SAFE_UPDATES = 0;