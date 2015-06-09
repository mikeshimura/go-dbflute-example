-- !df:pmb!
-- !!int64 loginId!!
-- !!string database!!
delete
from dbflute
where login_id = /*pmb.loginId*/1
and database = /*pmb.database*/'test'