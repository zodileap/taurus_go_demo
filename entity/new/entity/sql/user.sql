
/*
    Server Type: PostgreSQL
    Catalogs: user
    Schema: public
*/

-- ********
-- Sequence leapnote_tier_id_seq
-- ********
DO $$
BEGIN




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

    
    -- Foreign Key.
    -- 外键。
    ALTER TABLE "public"."post"
    ADD CONSTRAINT fk_blog_id FOREIGN KEY ("blog_id")
    REFERENCES "public"."blog" ("id");
    ALTER TABLE "public"."post"
    ADD CONSTRAINT fk_author_id FOREIGN KEY ("author_id")
    REFERENCES "public"."author" ("id");
END
$$;

