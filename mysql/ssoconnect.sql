use db_sso;

select * from t_user;
select * from t_token;
select * from t_email_verify;

SET SQL_SAFE_UPDATES = 0;

delete from t_user where t_user_username = 'liujunshi1';