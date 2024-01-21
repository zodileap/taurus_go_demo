
/*
    Server Type: PostgreSQL
    Catalogs: Permission
    Schema: public
*/





-- ********
-- Sequence user_id_seq
-- ********
DO $$
BEGIN
    CREATE  SEQUENCE IF NOT EXISTS "public"."user_id_seq" 
    INCREMENT 1
    MINVALUE  1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END
$$;

-- ********
-- Table "user"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'user') THEN
        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type FROM information_schema.columns tbl WHERE table_schema = 'public' AND table_name = 'user' LOOP
            IF column_rec.column_name NOT IN ('id','uuid','name','email','created_time','desc') THEN
                EXECUTE 'ALTER TABLE "public"."user" DROP COLUMN IF EXISTS ' || quote_ident(column_rec.column_name) || ' CASCADE;';
            END IF;
        END LOOP;
        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'user' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."user" ADD COLUMN "id" int8 NOT NULL DEFAULT nextval('user_id_seq'::regclass);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'user' AND column_name = 'uuid' ) THEN
            ALTER TABLE "public"."user" ADD COLUMN "uuid" varchar(255) NOT NULL;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'user' AND column_name = 'name' ) THEN
            ALTER TABLE "public"."user" ADD COLUMN "name" varchar(255) NOT NULL;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'user' AND column_name = 'email' ) THEN
            ALTER TABLE "public"."user" ADD COLUMN "email" varchar(255);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'user' AND column_name = 'created_time' ) THEN
            ALTER TABLE "public"."user" ADD COLUMN "created_time" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'user' AND column_name = 'desc' ) THEN
            ALTER TABLE "public"."user" ADD COLUMN "desc" varchar(255) DEFAULT 'desc';
        END IF;
        -- Search for the name of any existing primary key constraints. 
        -- If found, delete them first, then add new primary key constraints.
        -- 查找现有的主键约束名称，如果找到了先删除它， 添加新的主键约束。
        SELECT conname INTO v_constraint_name
        FROM pg_constraint con
        JOIN pg_class rel ON rel.oid = con.conrelid
        JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
        WHERE nsp.nspname = 'public'
            AND rel.relname = 'user'
            AND con.contype = 'p';
        IF v_constraint_name IS NOT NULL THEN
            EXECUTE 'ALTER TABLE "public"."user" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name);
        END IF;
        ALTER TABLE "public"."user" ADD CONSTRAINT user_pkey PRIMARY KEY ("id","uuid");
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."user" (
            "id" int8 NOT NULL DEFAULT nextval('user_id_seq'::regclass),
            "uuid" varchar(255) NOT NULL,
            "name" varchar(255) NOT NULL,
            "email" varchar(255),
            "created_time" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
            "desc" varchar(255) DEFAULT 'desc'
        );
        -- Field Comment.
        -- 字段备注。
        COMMENT ON COLUMN "public"."user"."id" IS  '用户的id';
        
        
        
        
        -- Primary Key.
        -- 主键。
        ALTER TABLE "public"."user" ADD CONSTRAINT user_pkey PRIMARY KEY ("id","uuid");
    END IF;
END
$$;

