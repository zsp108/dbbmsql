
create table bmsql_config (
  cfg_name    varchar2(30) PRIMARY KEY,
  cfg_value   varchar2(50)
);

create tablegroup "tpcc_group" partition by hash partitions 3;

create table bmsql_warehouse (
  w_id        NUMBER NOT NULL primary key, 
  w_ytd       DECIMAL(12,2),
  w_tax       DECIMAL(4,4),
  w_name      VARCHAR2(10),
  w_street_1  VARCHAR2(20),
  w_street_2  VARCHAR2(20),
  w_city      VARCHAR2(20),
  w_state     CHAR(2),
  w_zip       CHAR(9)
)tablegroup='tpcc_group' use_bloom_filter=true compress BLOCK_SIZE=16384  partition by hash(w_id) partitions 3;

create table bmsql_district (
  d_w_id       NUMBER       NOT NULL,
  d_id         NUMBER       NOT NULL,
  d_ytd        DECIMAL(12,2),
  d_tax        DECIMAL(4,4),
  d_next_o_id  NUMBER,
  d_name       VARCHAR2(10),
  d_street_1   VARCHAR2(20),
  d_street_2   VARCHAR2(20),
  d_city       VARCHAR2(20),
  d_state      CHAR(2),
  d_zip        CHAR(9),
  constraint pk_district PRIMARY KEY (d_w_id, d_id)
)tablegroup='tpcc_group' use_bloom_filter=true compress BLOCK_SIZE=16384  partition by hash(d_w_id) partitions 3;

create table bmsql_customer (
  c_w_id         NUMBER        NOT NULL,
  c_d_id         NUMBER        NOT NULL,
  c_id           NUMBER        NOT NULL,
  c_discount     DECIMAL(4,4),
  c_credit       CHAR(2),
  c_last         VARCHAR2(16),
  c_first        VARCHAR2(16),
  c_credit_lim   DECIMAL(12,2),
  c_balance      DECIMAL(12,2),
  c_ytd_payment  DECIMAL(12,2),
  c_payment_cnt  NUMBER,
  c_delivery_cnt NUMBER,
  c_street_1     VARCHAR2(20),
  c_street_2     VARCHAR2(20),
  c_city         VARCHAR2(20),
  c_state        CHAR(2),
  c_zip          CHAR(9),
  c_phone        CHAR(16),
  c_since        TIMESTAMP,
  c_middle       CHAR(2),
  c_data         VARCHAR2(500),
  constraint pk_customer PRIMARY KEY (c_w_id, c_d_id, c_id)
)tablegroup='tpcc_group' use_bloom_filter=true compress BLOCK_SIZE=16384 partition by hash(c_w_id) partitions 3;
create index idx_c_w_d_l_f_i on bmsql_customer (c_w_id, c_d_id, c_last, c_first, c_id) LOCAL;

create table bmsql_history (
  --auto increment
  hist_id  NUMBER,
  h_w_id   NUMBER,
  h_c_id   NUMBER,
  h_c_d_id NUMBER,
  h_c_w_id NUMBER,
  h_d_id   NUMBER,
  h_date   TIMESTAMP with local time zone,
  h_amount DECIMAL(6,2),
  h_data   VARCHAR2(24),
  constraint pk_history PRIMARY KEY (hist_id, h_w_id)
)tablegroup='tpcc_group' use_bloom_filter=true compress BLOCK_SIZE=16384 partition by hash(h_w_id) partitions 3;
create sequence history_seq
  START WITH 1
  INCREMENT BY 1
  NOMAXVALUE
  NOCYCLE
  CACHE 1000000
  NOORDER;

create table bmsql_new_order (
  no_w_id  NUMBER   NOT NULL,
  no_d_id  NUMBER   NOT NULL,
  no_o_id  NUMBER   NOT NULL,
  constraint pk_new_order PRIMARY KEY (no_w_id, no_d_id, no_o_id)
)tablegroup='tpcc_group' use_bloom_filter=true compress BLOCK_SIZE=16384  partition by hash(no_w_id) partitions 3;

create table bmsql_oorder (
  o_w_id       NUMBER      NOT NULL,
  o_d_id       NUMBER      NOT NULL,
  o_id         NUMBER      NOT NULL,
  o_c_id       NUMBER,
  o_carrier_id NUMBER,
  o_ol_cnt     DECIMAL(2,0),
  o_all_local  DECIMAL(1,0),
  o_entry_d    TIMESTAMP with time zone,
  constraint pk_oorder PRIMARY KEY (o_w_id, o_d_id, o_id)
)tablegroup='tpcc_group' use_bloom_filter=true compress BLOCK_SIZE=16384 partition by hash(o_w_id) partitions 3;
create index idx_o_w_d_c_i on bmsql_oorder (o_w_id, o_d_id, o_c_id, o_id) LOCAL;
create index idx_o_w_d_ol on bmsql_oorder (o_w_id, o_d_id, o_ol_cnt) LOCAL;

create table bmsql_order_line (
  ol_w_id         NUMBER   NOT NULL,
  ol_d_id         NUMBER   NOT NULL,
  ol_o_id         NUMBER   NOT NULL,
  ol_number       NUMBER   NOT NULL,
  ol_i_id         NUMBER   NOT NULL,
  ol_delivery_d   TIMESTAMP with time zone,
  ol_amount       DECIMAL(6,2),
  ol_supply_w_id  NUMBER,
  ol_quantity     DECIMAL(2,0),
  ol_dist_info    CHAR(24),
  constraint pk_order_line PRIMARY KEY (ol_w_id, ol_d_id, ol_o_id, ol_number)
)tablegroup='tpcc_group' use_bloom_filter=true compress BLOCK_SIZE=16384 partition by hash(ol_w_id) partitions 3;
create index idx_ol_sup_w_i on bmsql_order_line (ol_supply_w_id, ol_i_id) LOCAL;

create table bmsql_item (
  i_id     NUMBER      NOT NULL,
  i_name   VARCHAR2(24),
  i_price  DECIMAL(5,2),
  i_data   VARCHAR2(50),
  i_im_id  NUMBER,
  constraint pk_item PRIMARY KEY (i_id)
)use_bloom_filter=true compress BLOCK_SIZE=16384 locality='F,R{all_server}@z1, F,R{all_server}@z2, F,R{all_server}@z3' duplicate_scope='cluster' primary_zone='z1' partition by hash(i_id) partitions 3;

create table bmsql_stock (
  s_w_id       NUMBER       NOT NULL,
  s_i_id       NUMBER       NOT NULL,
  s_quantity   DECIMAL(4,0),
  s_ytd        DECIMAL(8,2),
  s_order_cnt  NUMBER,
  s_remote_cnt NUMBER,
  s_data       VARCHAR2(50),
  s_dist_01    CHAR(24),
  s_dist_02    CHAR(24),
  s_dist_03    CHAR(24),
  s_dist_04    CHAR(24),
  s_dist_05    CHAR(24),
  s_dist_06    CHAR(24),
  s_dist_07    CHAR(24),
  s_dist_08    CHAR(24),
  s_dist_09    CHAR(24),
  s_dist_10    CHAR(24),
  constraint pk_stock PRIMARY KEY (s_w_id, s_i_id)
)tablegroup='tpcc_group' use_bloom_filter=true compress BLOCK_SIZE=16384 partition by hash(s_w_id) partitions 3;
create index idx_s_i on bmsql_stock (s_i_id) LOCAL;

