-- 开始初始化数据 ;
insert into `sys_api`  values (1,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysMenu.GetPage-fm', '菜单列表', ' /api/v1/sys/menu', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (2,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysMenu.GetMenuRole-fm', '角色菜单【顶部左侧菜单】', ' /api/v1/sys/menu/menu-role', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (3,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysMenu.GetMenuTreeSelect-fm', '菜单权限列表【角色配置菜单使用】', '/api/v1/sys/menu/menu-tree/:roleId', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (4,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysMenu.Get-fm', '获取菜单详情', '/api/v1/sys/menu/:id', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (5,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.GetPage-fm ', '用户列表', '/api/v1/sys/user', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (6,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.GetProfile-fm', '用户详情', '/api/v1/sys/user/profile', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (7,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.Get-fm', '获取用户信息', '/api/v1/sys/user/:id', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (8,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysDept.GetPage-fm', '部门列表', '/api/v1/sys/dept', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (9,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysDept.Get-fm', '获取部门信息', '/api/v1/sys/dept/:deptId', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (10,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysPost.GetPage-fm', '岗位列表', '/api/v1/sys/post', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (11,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysPost.Get-fm', '获取部门信息', '/api/v1/sys/post/:id', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (12,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysRole.GetPage-fm', '角色列表', '/api/v1/sys/role', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (13,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysRole.Get-fm ', '获取角色信息', '/api/v1/sys/role/:id', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (14,'github.com/xiaodulala/admin-layout/internal/application/router.LoadRouter.func3', '获取版本信息', '/sys-version ', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (15,'github.com/xiaodulala/admin-layout/internal/application/router.LoadRouter.func2', '获取服务健康状态', '/healthz', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (16,'github.com/appleboy/gin-jwt/v2.(*GinJWTMiddleware).LogoutHandler-fm', '登出', '/logout', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (17,'github.com/appleboy/gin-jwt/v2.(*GinJWTMiddleware).RefreshHandler-fm', '刷新token', '/refresh-token', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (18,'github.com/xiaodulala/admin-layout/internal/application/router.LoadRouter.func4', '解析token载荷(仅限测试使用)', '/claims', 'SYS', 'GET', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (19,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.Insert-fm', '创建用户', '/api/v1/sys/user', 'SYS', 'POST', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (20,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.InsetAvatar-fm', '上传用户头像', '/api/v1/sys/user/avatar', 'SYS', 'POST', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (21,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysDept.Insert-fm', '创建部门', '/api/v1/sys/dept', 'SYS', 'POST', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (22,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysMenu.Insert-fm', '创建部门', '/api/v1/sys/menu', 'SYS', 'POST', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (23,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysPost.Insert-fm', '创建岗位', '/api/v1/sys/post', 'SYS', 'POST', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (24,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysRole.Insert-fm', '创建角色', '/api/v1/sys/role', 'SYS', 'POST', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (25,'github.com/appleboy/gin-jwt/v2.(*GinJWTMiddleware).LoginHandler-fm', '登录', '/login', 'SYS', 'POST', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (26,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.Update-fm', '修改用户', '/api/v1/sys/user', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (27,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.UpdatePwd-fm', '修改密码', '/api/v1/sys/user/pwd/set', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (28,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.ResetPwd-fm', '重置密码', '/api/v1/sys/user/pwd/reset', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (29,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.UpdateStatus-fm', '修改用户状态', '/api/v1/sys/user/status', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (30,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysRole.Update2Status-fm', '修改角色状态', '/api/v1/sys/role/role-status', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (31,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysRole.Update2DataScope-fm', '修改角色数据权限', '/api/v1/sys/role/role-scope', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (32,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysRole.Update-fm', '修改角色', '/api/v1/sys/role/:id', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (33,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysDept.Update-fm', '修改部门', '/api/v1/sys/dept/:deptId', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (34,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysMenu.Update-fm', '修改菜单', '/api/v1/sys/menu/:id', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (35,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysPost.Update-fm', '修改岗位', '/api/v1/sys/post/:id ', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (36,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysDept.Delete-fm', '删除部门', '/api/v1/sys/dept', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (37,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysMenu.Delete-fm', '删除菜单', '/api/v1/sys/menu', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (38,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysPost.Delete-fm', '删除岗位', '/api/v1/sys/post', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (39,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysRole.Delete-fm ', '删除角色', '/api/v1/sys/role', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);
insert into `sys_api`  values (40,'github.com/xiaodulala/admin-layout/internal/application/sys/apis.SysUser.Delete-fm', '删除用户', '/api/v1/sys/user/:userId', 'SYS', 'PUT', '2022-03-10 11:02:27.000', null, null, null, null);


INSERT INTO sys_menu VALUES (2, 'Admin', '系统管理', 'api-server', '/admin', '/0/2', 'M', '无', '', 0, 1, '', 'Layout', 10, '0', '1', 0, 1, '2021-05-20 21:58:45.679', '2021-06-17 11:48:40.703', NULL);

INSERT INTO sys_role VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_user VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'fskj', '13818888888', 1, '', '', '1', '1@qq.com', 1, 1, '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
-- 数据完成 ;
