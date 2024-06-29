
/*
    Server Type: PostgreSQL
    Catalogs: user
    Schema: public
*/

-- ********
-- Delete Foreign Key
-- ********
DO $$
BEGIN

IF EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_schema = 'public'
    AND table_name = 'post'
    AND constraint_name = 'fk_author_id'
) THEN
    ALTER TABLE "public"."post" DROP CONSTRAINT "fk_author_id";
END IF;

IF EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_schema = 'public'
    AND table_name = 'post'
    AND constraint_name = 'fk_blog_id'
) THEN
    ALTER TABLE "public"."post" DROP CONSTRAINT "fk_blog_id";
END IF;

IF EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_schema = 'public'
    AND table_name = 'post'
    AND constraint_name = 'fk_blog_id'
) THEN
    ALTER TABLE "public"."post" DROP CONSTRAINT "fk_blog_id";
END IF;

IF EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_schema = 'public'
    AND table_name = 'post'
    AND constraint_name = 'fk_author_id'
) THEN
    ALTER TABLE "public"."post" DROP CONSTRAINT "fk_author_id";
END IF;

END
$$;


-- ********
-- Sequence author_id_seq
-- ********
DO $$
BEGIN
    CREATE  SEQUENCE IF NOT EXISTS "public"."author_id_seq" 
    INCREMENT 1
    MINVALUE  1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END
$$;
-- ********
-- Table "author"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'author') THEN
        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type FROM information_schema.columns tbl WHERE table_schema = 'public' AND table_name = 'author' LOOP
            IF column_rec.column_name NOT IN ('id','name') THEN
                EXECUTE 'ALTER TABLE "public"."author" DROP COLUMN IF EXISTS ' || quote_ident(column_rec.column_name) || ' CASCADE;';
            END IF;
        END LOOP;
        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'author' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."author" ADD COLUMN "id" int8 NOT NULL DEFAULT nextval('author_id_seq'::regclass);
        ELSE
            
            ALTER TABLE "public"."author" ALTER COLUMN "id" SET NOT NULL; 
            ALTER TABLE "public"."author" ALTER COLUMN "id" SET DEFAULT nextval('author_id_seq'::regclass); ALTER TABLE "public"."author" ALTER COLUMN "id" TYPE int8 USING "id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'author' AND column_name = 'name' ) THEN
            ALTER TABLE "public"."author" ADD COLUMN "name" varchar(255) NOT NULL;
        ELSE
            
            ALTER TABLE "public"."author" ALTER COLUMN "name" SET NOT NULL; 
            ALTER TABLE "public"."author" ALTER COLUMN "name" DROP DEFAULT; ALTER TABLE "public"."author" ALTER COLUMN "name" TYPE varchar(255) USING "name"::varchar(255);
        END IF;
        -- Search for the name of any existing primary key constraints. 
        -- If found, delete them first, then add new primary key constraints.
        -- 查找现有的主键约束名称，如果找到了先删除它， 添加新的主键约束。
        SELECT conname INTO v_constraint_name
        FROM pg_constraint con
        JOIN pg_class rel ON rel.oid = con.conrelid
        JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
        WHERE nsp.nspname = 'public'
            AND rel.relname = 'author'
            AND con.contype = 'p';
        IF v_constraint_name IS NOT NULL THEN
            EXECUTE 'ALTER TABLE "public"."author" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name);
        END IF;
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."author" (
            "id" int8 NOT NULL DEFAULT nextval('author_id_seq'::regclass),
            "name" varchar(255) NOT NULL
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."author"."id" IS  'Author primary key';
    

    -- Primary Key.
     -- 主键。
    ALTER TABLE "public"."author" ADD CONSTRAINT author_pkey PRIMARY KEY ("id");
END
$$;

-- ********
-- Sequence blog_id_seq
-- ********
DO $$
BEGIN
    CREATE  SEQUENCE IF NOT EXISTS "public"."blog_id_seq" 
    INCREMENT 1
    MINVALUE  1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END
$$;
-- ********
-- Table "blog"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'blog') THEN
        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type FROM information_schema.columns tbl WHERE table_schema = 'public' AND table_name = 'blog' LOOP
            IF column_rec.column_name NOT IN ('id','uuid','desc','created_time') THEN
                EXECUTE 'ALTER TABLE "public"."blog" DROP COLUMN IF EXISTS ' || quote_ident(column_rec.column_name) || ' CASCADE;';
            END IF;
        END LOOP;
        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "id" int8 NOT NULL DEFAULT nextval('blog_id_seq'::regclass);
        ELSE
            
            ALTER TABLE "public"."blog" ALTER COLUMN "id" SET NOT NULL; 
            ALTER TABLE "public"."blog" ALTER COLUMN "id" SET DEFAULT nextval('blog_id_seq'::regclass); ALTER TABLE "public"."blog" ALTER COLUMN "id" TYPE int8 USING "id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'uuid' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "uuid" uuid NOT NULL;
        ELSE
            
            ALTER TABLE "public"."blog" ALTER COLUMN "uuid" SET NOT NULL; 
            ALTER TABLE "public"."blog" ALTER COLUMN "uuid" DROP DEFAULT; ALTER TABLE "public"."blog" ALTER COLUMN "uuid" TYPE uuid USING "uuid"::uuid;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'desc' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "desc" varchar(255);
        ELSE
            
            ALTER TABLE "public"."blog" ALTER COLUMN "desc" DROP NOT NULL; 
            ALTER TABLE "public"."blog" ALTER COLUMN "desc" DROP DEFAULT; ALTER TABLE "public"."blog" ALTER COLUMN "desc" TYPE varchar(255) USING "desc"::varchar(255);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'created_time' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "created_time" timestamptz(6) DEFAULT CURRENT_TIMESTAMP;
        ELSE
            
            ALTER TABLE "public"."blog" ALTER COLUMN "created_time" DROP NOT NULL; 
            ALTER TABLE "public"."blog" ALTER COLUMN "created_time" SET DEFAULT CURRENT_TIMESTAMP; ALTER TABLE "public"."blog" ALTER COLUMN "created_time" TYPE timestamptz(6) USING "created_time"::timestamptz(6);
        END IF;
        -- Search for the name of any existing primary key constraints. 
        -- If found, delete them first, then add new primary key constraints.
        -- 查找现有的主键约束名称，如果找到了先删除它， 添加新的主键约束。
        SELECT conname INTO v_constraint_name
        FROM pg_constraint con
        JOIN pg_class rel ON rel.oid = con.conrelid
        JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
        WHERE nsp.nspname = 'public'
            AND rel.relname = 'blog'
            AND con.contype = 'p';
        IF v_constraint_name IS NOT NULL THEN
            EXECUTE 'ALTER TABLE "public"."blog" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name);
        END IF;
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."blog" (
            "id" int8 NOT NULL DEFAULT nextval('blog_id_seq'::regclass),
            "uuid" uuid NOT NULL,
            "desc" varchar(255),
            "created_time" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."blog"."id" IS  'Blog primary key';
    
    
    

    -- Primary Key.
     -- 主键。
    ALTER TABLE "public"."blog" ADD CONSTRAINT blog_pkey PRIMARY KEY ("id");
END
$$;

-- ********
-- Table "field_demo"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'field_demo') THEN
        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type FROM information_schema.columns tbl WHERE table_schema = 'public' AND table_name = 'field_demo' LOOP
            IF column_rec.column_name NOT IN ('int64_f','var_f','bool_f','int_array_f','int_array2_f','bool_array_f','time_f','time_array_f') THEN
                EXECUTE 'ALTER TABLE "public"."field_demo" DROP COLUMN IF EXISTS ' || quote_ident(column_rec.column_name) || ' CASCADE;';
            END IF;
        END LOOP;
        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'int64_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "int64_f" int8 NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "int64_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "int64_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "int64_f" TYPE int8 USING "int64_f"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'var_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "var_f" varchar(255) NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "var_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "var_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "var_f" TYPE varchar(255) USING "var_f"::varchar(255);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'bool_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "bool_f" boolean NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "bool_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "bool_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "bool_f" TYPE boolean USING "bool_f"::boolean;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'int_array_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "int_array_f" int8[] NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "int_array_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "int_array_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "int_array_f" TYPE int8[] USING "int_array_f"::int8[];
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'int_array2_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "int_array2_f" int8[][] NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "int_array2_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "int_array2_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "int_array2_f" TYPE int8[][] USING "int_array2_f"::int8[][];
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'bool_array_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "bool_array_f" boolean[] NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "bool_array_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "bool_array_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "bool_array_f" TYPE boolean[] USING "bool_array_f"::boolean[];
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'time_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "time_f" timestamptz(6) NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "time_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "time_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "time_f" TYPE timestamptz(6) USING "time_f"::timestamptz(6);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'time_array_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "time_array_f" timestamptz[] NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "time_array_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "time_array_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "time_array_f" TYPE timestamptz[] USING "time_array_f"::timestamptz[];
        END IF;
        -- Search for the name of any existing primary key constraints. 
        -- If found, delete them first, then add new primary key constraints.
        -- 查找现有的主键约束名称，如果找到了先删除它， 添加新的主键约束。
        SELECT conname INTO v_constraint_name
        FROM pg_constraint con
        JOIN pg_class rel ON rel.oid = con.conrelid
        JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
        WHERE nsp.nspname = 'public'
            AND rel.relname = 'field_demo'
            AND con.contype = 'p';
        IF v_constraint_name IS NOT NULL THEN
            EXECUTE 'ALTER TABLE "public"."field_demo" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name);
        END IF;
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."field_demo" (
            "int64_f" int8 NOT NULL,
            "var_f" varchar(255) NOT NULL,
            "bool_f" boolean NOT NULL,
            "int_array_f" int8[] NOT NULL,
            "int_array2_f" int8[][] NOT NULL,
            "bool_array_f" boolean[] NOT NULL,
            "time_f" timestamptz(6) NOT NULL,
            "time_array_f" timestamptz[] NOT NULL
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."field_demo"."int64_f" IS  'Int64 field';
    COMMENT ON COLUMN "public"."field_demo"."var_f" IS  'Varchar field';
    COMMENT ON COLUMN "public"."field_demo"."bool_f" IS  'Bool field';
    COMMENT ON COLUMN "public"."field_demo"."int_array_f" IS  'Int array field';
    COMMENT ON COLUMN "public"."field_demo"."int_array2_f" IS  'Int array2 field';
    COMMENT ON COLUMN "public"."field_demo"."bool_array_f" IS  'Bool array field';
    COMMENT ON COLUMN "public"."field_demo"."time_f" IS  'Time field';
    COMMENT ON COLUMN "public"."field_demo"."time_array_f" IS  'Time array field';

    -- Primary Key.
     -- 主键。
    ALTER TABLE "public"."field_demo" ADD CONSTRAINT field_demo_pkey PRIMARY KEY ("int64_f");
END
$$;

-- ********
-- Sequence geo_id_seq
-- ********
DO $$
BEGIN
    CREATE  SEQUENCE IF NOT EXISTS "public"."geo_id_seq" 
    INCREMENT 1
    MINVALUE  1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END
$$;
-- ********
-- Table "geo_demo"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'geo_demo') THEN
        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type FROM information_schema.columns tbl WHERE table_schema = 'public' AND table_name = 'geo_demo' LOOP
            IF column_rec.column_name NOT IN ('id','point','line_string','polygon','multi_point','multi_line_string','multi_polygon','circular_string') THEN
                EXECUTE 'ALTER TABLE "public"."geo_demo" DROP COLUMN IF EXISTS ' || quote_ident(column_rec.column_name) || ' CASCADE;';
            END IF;
        END LOOP;
        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "id" int8 NOT NULL DEFAULT nextval('geo_id_seq'::regclass);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "id" SET NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "id" SET DEFAULT nextval('geo_id_seq'::regclass); ALTER TABLE "public"."geo_demo" ALTER COLUMN "id" TYPE int8 USING "id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'point' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "point" geometry(Point, 4326) NOT NULL;
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "point" SET NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "point" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "point" TYPE geometry(Point, 4326) USING "point"::geometry(Point, 4326);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'line_string' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "line_string" geometry(LineString, 0) NOT NULL;
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string" SET NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string" TYPE geometry(LineString, 0) USING "line_string"::geometry(LineString, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'polygon' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "polygon" geometry(Polygon, 0) NOT NULL;
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "polygon" SET NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "polygon" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "polygon" TYPE geometry(Polygon, 0) USING "polygon"::geometry(Polygon, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'multi_point' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "multi_point" geometry(MultiPoint, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_point" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_point" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_point" TYPE geometry(MultiPoint, 0) USING "multi_point"::geometry(MultiPoint, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'multi_line_string' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "multi_line_string" geometry(MultiLineString, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_line_string" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_line_string" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_line_string" TYPE geometry(MultiLineString, 0) USING "multi_line_string"::geometry(MultiLineString, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'multi_polygon' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "multi_polygon" geometry(MultiPolygon, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_polygon" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_polygon" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_polygon" TYPE geometry(MultiPolygon, 0) USING "multi_polygon"::geometry(MultiPolygon, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'circular_string' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "circular_string" geometry(CircularString, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "circular_string" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "circular_string" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "circular_string" TYPE geometry(CircularString, 0) USING "circular_string"::geometry(CircularString, 0);
        END IF;
        -- Search for the name of any existing primary key constraints. 
        -- If found, delete them first, then add new primary key constraints.
        -- 查找现有的主键约束名称，如果找到了先删除它， 添加新的主键约束。
        SELECT conname INTO v_constraint_name
        FROM pg_constraint con
        JOIN pg_class rel ON rel.oid = con.conrelid
        JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
        WHERE nsp.nspname = 'public'
            AND rel.relname = 'geo_demo'
            AND con.contype = 'p';
        IF v_constraint_name IS NOT NULL THEN
            EXECUTE 'ALTER TABLE "public"."geo_demo" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name);
        END IF;
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."geo_demo" (
            "id" int8 NOT NULL DEFAULT nextval('geo_id_seq'::regclass),
            "point" geometry(Point, 4326) NOT NULL,
            "line_string" geometry(LineString, 0) NOT NULL,
            "polygon" geometry(Polygon, 0) NOT NULL,
            "multi_point" geometry(MultiPoint, 0),
            "multi_line_string" geometry(MultiLineString, 0),
            "multi_polygon" geometry(MultiPolygon, 0),
            "circular_string" geometry(CircularString, 0)
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."geo_demo"."id" IS  '主键。';
    COMMENT ON COLUMN "public"."geo_demo"."point" IS  '点';
    COMMENT ON COLUMN "public"."geo_demo"."line_string" IS  '线';
    COMMENT ON COLUMN "public"."geo_demo"."polygon" IS  '多边形';
    COMMENT ON COLUMN "public"."geo_demo"."multi_point" IS  '多点';
    COMMENT ON COLUMN "public"."geo_demo"."multi_line_string" IS  '多线';
    COMMENT ON COLUMN "public"."geo_demo"."multi_polygon" IS  '多多边形';
    COMMENT ON COLUMN "public"."geo_demo"."circular_string" IS  '圆弧';
    -- Table Comment.
    -- 表备注。
    COMMENT ON TABLE "public"."geo_demo" IS 'Geo的类型测试';

    -- Primary Key.
     -- 主键。
    ALTER TABLE "public"."geo_demo" ADD CONSTRAINT geo_demo_pkey PRIMARY KEY ("id");
END
$$;

-- ********
-- Sequence post_id_seq
-- ********
DO $$
BEGIN
    CREATE  SEQUENCE IF NOT EXISTS "public"."post_id_seq" 
    INCREMENT 1
    MINVALUE  1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END
$$;
-- ********
-- Table "post"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'post') THEN
        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type FROM information_schema.columns tbl WHERE table_schema = 'public' AND table_name = 'post' LOOP
            IF column_rec.column_name NOT IN ('id','content','blog_id','author_id') THEN
                EXECUTE 'ALTER TABLE "public"."post" DROP COLUMN IF EXISTS ' || quote_ident(column_rec.column_name) || ' CASCADE;';
            END IF;
        END LOOP;
        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'post' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."post" ADD COLUMN "id" int8 NOT NULL DEFAULT nextval('post_id_seq'::regclass);
        ELSE
            
            ALTER TABLE "public"."post" ALTER COLUMN "id" SET NOT NULL; 
            ALTER TABLE "public"."post" ALTER COLUMN "id" SET DEFAULT nextval('post_id_seq'::regclass); ALTER TABLE "public"."post" ALTER COLUMN "id" TYPE int8 USING "id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'post' AND column_name = 'content' ) THEN
            ALTER TABLE "public"."post" ADD COLUMN "content" varchar(255) NOT NULL;
        ELSE
            
            ALTER TABLE "public"."post" ALTER COLUMN "content" SET NOT NULL; 
            ALTER TABLE "public"."post" ALTER COLUMN "content" DROP DEFAULT; ALTER TABLE "public"."post" ALTER COLUMN "content" TYPE varchar(255) USING "content"::varchar(255);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'post' AND column_name = 'blog_id' ) THEN
            ALTER TABLE "public"."post" ADD COLUMN "blog_id" int8 NOT NULL;
        ELSE
            
            ALTER TABLE "public"."post" ALTER COLUMN "blog_id" SET NOT NULL; 
            ALTER TABLE "public"."post" ALTER COLUMN "blog_id" DROP DEFAULT; ALTER TABLE "public"."post" ALTER COLUMN "blog_id" TYPE int8 USING "blog_id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'post' AND column_name = 'author_id' ) THEN
            ALTER TABLE "public"."post" ADD COLUMN "author_id" int8 NOT NULL;
        ELSE
            
            ALTER TABLE "public"."post" ALTER COLUMN "author_id" SET NOT NULL; 
            ALTER TABLE "public"."post" ALTER COLUMN "author_id" DROP DEFAULT; ALTER TABLE "public"."post" ALTER COLUMN "author_id" TYPE int8 USING "author_id"::int8;
        END IF;
        -- Search for the name of any existing primary key constraints. 
        -- If found, delete them first, then add new primary key constraints.
        -- 查找现有的主键约束名称，如果找到了先删除它， 添加新的主键约束。
        SELECT conname INTO v_constraint_name
        FROM pg_constraint con
        JOIN pg_class rel ON rel.oid = con.conrelid
        JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
        WHERE nsp.nspname = 'public'
            AND rel.relname = 'post'
            AND con.contype = 'p';
        IF v_constraint_name IS NOT NULL THEN
            EXECUTE 'ALTER TABLE "public"."post" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name);
        END IF;
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."post" (
            "id" int8 NOT NULL DEFAULT nextval('post_id_seq'::regclass),
            "content" varchar(255) NOT NULL,
            "blog_id" int8 NOT NULL,
            "author_id" int8 NOT NULL
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."post"."id" IS  'Post primary key';
    
    
    

    -- Primary Key.
     -- 主键。
    ALTER TABLE "public"."post" ADD CONSTRAINT post_pkey PRIMARY KEY ("id");
END
$$;





-- ********
-- Add Foreign Key
-- ********
DO $$
BEGIN

ALTER TABLE "public"."post"
ADD CONSTRAINT fk_blog_id FOREIGN KEY ("blog_id")
REFERENCES "public"."blog" ("id");

ALTER TABLE "public"."post"
ADD CONSTRAINT fk_author_id FOREIGN KEY ("author_id")
REFERENCES "public"."author" ("id");

END
$$;

