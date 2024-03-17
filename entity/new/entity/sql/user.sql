
/*
    Server Type: PostgreSQL
    Catalogs: user
    Schema: public
*/





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
            IF column_rec.column_name NOT IN ('i_d','uuid','desc','created_time') THEN
                EXECUTE 'ALTER TABLE "public"."blog" DROP COLUMN IF EXISTS ' || quote_ident(column_rec.column_name) || ' CASCADE;';
            END IF;
        END LOOP;
        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'i_d' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "i_d" int8 NOT NULL DEFAULT nextval('blog_id_seq'::regclass);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'uuid' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "uuid" uuid NOT NULL;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'desc' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "desc" varchar(255);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'created_time' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "created_time" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP;
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
        ALTER TABLE "public"."blog" ADD CONSTRAINT blog_pkey PRIMARY KEY ("i_d");
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."blog" (
            "i_d" int8 NOT NULL DEFAULT nextval('blog_id_seq'::regclass),
            "uuid" uuid NOT NULL,
            "desc" varchar(255),
            "created_time" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
        );
        -- Field Comment.
        -- 字段备注。
        COMMENT ON COLUMN "public"."blog"."i_d" IS  'Blog primary key';
        
        
        -- Primary Key.
        -- 主键。
        ALTER TABLE "public"."blog" ADD CONSTRAINT blog_pkey PRIMARY KEY ("i_d");
    END IF;
END
$$;

