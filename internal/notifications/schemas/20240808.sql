create or replace function set_web_app_verification_time()
returns trigger as $$
	begin
		if NEW.is_verified = true then
		   NEW.verified_at = now()::timestamp;
  		end if;
  	  	return NEW;
  	end;
	$$
	language 'plpgsql';

create or replace trigger on_web_app_is_verified_update  
	after update of is_verified on web_apps
  	for each row
	execute procedure set_web_app_verification_time();
