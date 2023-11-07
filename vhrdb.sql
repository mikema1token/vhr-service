create table department
(
	id int auto_increment
		primary key,
	name varchar(32) null comment '部门名称',
	parentId int null,
	depPath varchar(255) null,
	enabled tinyint(1) default 1 null,
	isParent tinyint(1) default 0 null
)
charset=utf8mb3;

create table hr
(
	id int auto_increment comment 'hrID'
		primary key,
	name varchar(32) null comment '姓名',
	phone char(11) null comment '手机号码',
	telephone varchar(16) null comment '住宅电话',
	address varchar(64) null comment '联系地址',
	enabled tinyint(1) default 1 null,
	username varchar(255) null comment '用户名',
	password varchar(255) null comment '密码',
	userface varchar(255) null,
	remark varchar(255) null
)
charset=utf8mb3;

create table joblevel
(
	id int auto_increment
		primary key,
	name varchar(32) null comment '职称名称',
	titleLevel enum('正高级', '副高级', '中级', '初级', '员级') null,
	createDate timestamp default CURRENT_TIMESTAMP null,
	enabled tinyint(1) default 1 null
)
charset=utf8mb3;

create table mail_send_log
(
	msgId varchar(255) null,
	empId int null,
	status int default 0 null comment '0发送中，1发送成功，2发送失败',
	routeKey varchar(255) null,
	exchange varchar(255) null,
	count int null comment '重试次数',
	tryTime date null comment '第一次重试时间',
	createTime date null,
	updateTime date null
)
collate=utf8mb4_general_ci;

create table menu
(
	id int auto_increment
		primary key,
	url varchar(64) null,
	path varchar(64) null,
	component varchar(64) null,
	name varchar(64) null,
	iconCls varchar(64) null,
	keepAlive tinyint(1) null,
	requireAuth tinyint(1) null,
	parentId int null,
	enabled tinyint(1) default 1 null,
	constraint menu_ibfk_1
		foreign key (parentId) references menu (id)
)
charset=utf8mb3;

create index parentId
	on menu (parentId);

create table msgcontent
(
	id int auto_increment
		primary key,
	title varchar(64) null,
	message varchar(255) null,
	createDate timestamp default CURRENT_TIMESTAMP not null
)
charset=utf8mb3;

create table nation
(
	id int auto_increment
		primary key,
	name varchar(32) null
)
charset=utf8mb3;

create table oplog
(
	id int auto_increment
		primary key,
	addDate date null comment '添加日期',
	operate varchar(255) null comment '操作内容',
	hrid int null comment '操作员ID',
	constraint oplog_ibfk_1
		foreign key (hrid) references hr (id)
)
charset=utf8mb3;

create index hrid
	on oplog (hrid);

create table politicsstatus
(
	id int auto_increment
		primary key,
	name varchar(32) null
)
charset=utf8mb3;

create table position
(
	id int auto_increment
		primary key,
	name varchar(32) null comment '职位',
	createDate timestamp default CURRENT_TIMESTAMP null,
	enabled tinyint(1) default 1 null,
	constraint name
		unique (name)
)
charset=utf8mb3;

create table employee
(
	id int auto_increment comment '员工编号'
		primary key,
	name varchar(10) null comment '员工姓名',
	gender char(4) null comment '性别',
	birthday date null comment '出生日期',
	idCard char(18) null comment '身份证号',
	wedlock enum('已婚', '未婚', '离异') null comment '婚姻状况',
	nationId int null comment '民族',
	nativePlace varchar(20) null comment '籍贯',
	politicId int null comment '政治面貌',
	email varchar(20) null comment '邮箱',
	phone varchar(11) null comment '电话号码',
	address varchar(64) null comment '联系地址',
	departmentId int null comment '所属部门',
	jobLevelId int null comment '职称ID',
	posId int null comment '职位ID',
	engageForm varchar(8) null comment '聘用形式',
	tiptopDegree enum('博士', '硕士', '本科', '大专', '高中', '初中', '小学', '其他') null comment '最高学历',
	specialty varchar(32) null comment '所属专业',
	school varchar(32) null comment '毕业院校',
	beginDate date null comment '入职日期',
	workState enum('在职', '离职') default '在职' null comment '在职状态',
	workID char(8) null comment '工号',
	contractTerm double null comment '合同期限',
	conversionTime date null comment '转正日期',
	notWorkDate date null comment '离职日期',
	beginContract date null comment '合同起始日期',
	endContract date null comment '合同终止日期',
	workAge int null comment '工龄',
	constraint employee_ibfk_1
		foreign key (departmentId) references department (id),
	constraint employee_ibfk_2
		foreign key (jobLevelId) references joblevel (id),
	constraint employee_ibfk_3
		foreign key (posId) references position (id),
	constraint employee_ibfk_4
		foreign key (nationId) references nation (id),
	constraint employee_ibfk_5
		foreign key (politicId) references politicsstatus (id)
)
charset=utf8mb3;

create table adjustsalary
(
	id int auto_increment
		primary key,
	eid int null,
	asDate date null comment '调薪日期',
	beforeSalary int null comment '调前薪资',
	afterSalary int null comment '调后薪资',
	reason varchar(255) null comment '调薪原因',
	remark varchar(255) null comment '备注',
	constraint adjustsalary_ibfk_1
		foreign key (eid) references employee (id)
)
charset=utf8mb3;

create index pid
	on adjustsalary (eid);

create table appraise
(
	id int auto_increment
		primary key,
	eid int null,
	appDate date null comment '考评日期',
	appResult varchar(32) null comment '考评结果',
	appContent varchar(255) null comment '考评内容',
	remark varchar(255) null comment '备注',
	constraint appraise_ibfk_1
		foreign key (eid) references employee (id)
)
charset=utf8mb3;

create index pid
	on appraise (eid);

create index departmentId
	on employee (departmentId);

create index dutyId
	on employee (posId);

create index jobId
	on employee (jobLevelId);

create index nationId
	on employee (nationId);

create index politicId
	on employee (politicId);

create index workID_key
	on employee (workID);

create table employeeec
(
	id int auto_increment
		primary key,
	eid int null comment '员工编号',
	ecDate date null comment '奖罚日期',
	ecReason varchar(255) null comment '奖罚原因',
	ecPoint int null comment '奖罚分',
	ecType int null comment '奖罚类别，0：奖，1：罚',
	remark varchar(255) null comment '备注',
	constraint employeeec_ibfk_1
		foreign key (eid) references employee (id)
)
charset=utf8mb3;

create index pid
	on employeeec (eid);

create table employeeremove
(
	id int auto_increment
		primary key,
	eid int null,
	afterDepId int null comment '调动后部门',
	afterJobId int null comment '调动后职位',
	removeDate date null comment '调动日期',
	reason varchar(255) null comment '调动原因',
	remark varchar(255) null,
	constraint employeeremove_ibfk_1
		foreign key (eid) references employee (id)
)
charset=utf8mb3;

create index pid
	on employeeremove (eid);

create table employeetrain
(
	id int auto_increment
		primary key,
	eid int null comment '员工编号',
	trainDate date null comment '培训日期',
	trainContent varchar(255) null comment '培训内容',
	remark varchar(255) null comment '备注',
	constraint employeetrain_ibfk_1
		foreign key (eid) references employee (id)
)
charset=utf8mb3;

create index pid
	on employeetrain (eid);

create table role
(
	id int auto_increment
		primary key,
	name varchar(64) null,
	nameZh varchar(64) null comment '角色名称'
)
charset=utf8mb3;

create table hr_role
(
	id int auto_increment
		primary key,
	hrid int null,
	rid int null,
	constraint hr_role_ibfk_1
		foreign key (hrid) references hr (id)
			on delete cascade,
	constraint hr_role_ibfk_2
		foreign key (rid) references role (id)
)
charset=utf8mb3;

create index rid
	on hr_role (rid);

create table menu_role
(
	id int auto_increment
		primary key,
	mid int null,
	rid int null,
	constraint menu_role_ibfk_1
		foreign key (mid) references menu (id),
	constraint menu_role_ibfk_2
		foreign key (rid) references role (id)
)
charset=utf8mb3;

create index mid
	on menu_role (mid);

create index rid
	on menu_role (rid);

create table salary
(
	id int auto_increment
		primary key,
	basicSalary int null comment '基本工资',
	bonus int null comment '奖金',
	lunchSalary int null comment '午餐补助',
	trafficSalary int null comment '交通补助',
	allSalary int null comment '应发工资',
	pensionBase int null comment '养老金基数',
	pensionPer float null comment '养老金比率',
	createDate timestamp null comment '启用时间',
	medicalBase int null comment '医疗基数',
	medicalPer float null comment '医疗保险比率',
	accumulationFundBase int null comment '公积金基数',
	accumulationFundPer float null comment '公积金比率',
	name varchar(32) null
)
charset=utf8mb3;

create table empsalary
(
	id int auto_increment
		primary key,
	eid int null,
	sid int null,
	constraint eid
		unique (eid),
	constraint empsalary_ibfk_1
		foreign key (eid) references employee (id),
	constraint empsalary_ibfk_2
		foreign key (sid) references salary (id)
)
charset=utf8mb3;

create table sysmsg
(
	id int auto_increment
		primary key,
	mid int null comment '消息id',
	type int default 0 null comment '0表示群发消息',
	hrid int null comment '这条消息是给谁的',
	state int default 0 null comment '0 未读 1 已读',
	constraint sysmsg_ibfk_1
		foreign key (mid) references msgcontent (id),
	constraint sysmsg_ibfk_2
		foreign key (hrid) references hr (id)
)
charset=utf8mb3;

create index hrid
	on sysmsg (hrid);

create definer = root@localhost procedure addDep(IN depName varchar(32), IN parentId int, IN enabled tinyint(1), OUT result int, OUT result2 int)
begin
  declare did int;
  declare pDepPath varchar(64);
  insert into department set name=depName,parentId=parentId,enabled=enabled;
  select row_count() into result;
  select last_insert_id() into did;
  set result2=did;
  select depPath into pDepPath from department where id=parentId;
  update department set depPath=concat(pDepPath,'.',did) where id=did;
  update department set isParent=true where id=parentId;
end;

create definer = root@localhost procedure deleteDep(IN did int, OUT result int)
begin
  declare ecount int;
  declare pid int;
  declare pcount int;
  declare a int;
  select count(*) into a from department where id=did and isParent=false;
  if a=0 then set result=-2;
  else
  select count(*) into ecount from employee where departmentId=did;
  if ecount>0 then set result=-1;
  else
  select parentId into pid from department where id=did;
  delete from department where id=did and isParent=false;
  select row_count() into result;
  select count(*) into pcount from department where parentId=pid;
  if pcount=0 then update department set isParent=false where id=pid;
  end if;
  end if;
  end if;
end;