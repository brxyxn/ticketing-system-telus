-- join teams and tiers
select teams.name,
    tiers.name
from teams
    join tiers on teams.tier_id = tiers.tier_id;
--------------------------------------------------------------------------------
-- users
-- customers
select customer_id,
    companies.name
from customers
    join companies on customers.company_id = companies.company_id;
-- join user + customer + company
select usr.user_id, usr.email as uemail, com.name as cname
from users usr
inner join customers cus on cus.customer_id = usr.customer_id
inner join companies com on cus.company_id = com.company_id;
-- join profile + user + customer + company
select p.profile_id, p.first_name, p.last_name, p.user_id,
u.email, u.status, u.customer_id,
c.name
from profiles p 
inner join users u on u.user_id = p.user_id
inner join customers cx on cx.customer_id = u.customer_id 
inner join companies c on c.company_id = cx.company_id;