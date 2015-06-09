-- #df:entity#
-- !df:pmb!
-- !!int64 loginId!!
-- !!string db!!
SELECT
 db_name
from dbflute
where login_id = /*pmb.loginId*/3
and db = /*pmb.db*/'PostgreSQL'
and del_flag = 0
group by db_name
order by db_name