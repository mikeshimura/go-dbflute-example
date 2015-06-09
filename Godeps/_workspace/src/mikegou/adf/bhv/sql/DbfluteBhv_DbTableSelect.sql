-- #df:entity#
-- !df:pmb!
-- !!int64 loginId!!
-- !!string db!!
-- !!string dbName!!
SELECT
 db_table
from dbflute
where login_id = /*pmb.loginId*/3
and db = /*pmb.db*/'PostgreSQL'
and db_name = /*pmb.dbName*/'mikegou'
and del_flag = 0
group by db_table
order by db_table 