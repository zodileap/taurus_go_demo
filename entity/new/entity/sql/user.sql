
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
    ALTER TABLE "public"."post" DROP CONSTRAINT IF EXISTS "fk_author_id" CASCADE;
END IF;

IF EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_schema = 'public'
    AND table_name = 'post'
    AND constraint_name = 'fk_blog_id'
) THEN
    ALTER TABLE "public"."post" DROP CONSTRAINT IF EXISTS "fk_blog_id" CASCADE;
END IF;

IF EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_schema = 'public'
    AND table_name = 'post'
    AND constraint_name = 'fk_blog_id'
) THEN
    ALTER TABLE "public"."post" DROP CONSTRAINT IF EXISTS "fk_blog_id" CASCADE;
END IF;

IF EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_schema = 'public'
    AND table_name = 'post'
    AND constraint_name = 'fk_author_id'
) THEN
    ALTER TABLE "public"."post" DROP CONSTRAINT IF EXISTS "fk_author_id" CASCADE;
END IF;

END
$$;


-- ********
-- Sequence author_id_seq
-- ********
DO $$ 
BEGIN     
    -- 创建随机种子序列
    CREATE SEQUENCE IF NOT EXISTS "public"."author_id_seq_seed"
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END $$;
CREATE OR REPLACE FUNCTION "public".author_id_seq() 
RETURNS BIGINT AS $$
DECLARE
    timestamp_part BIGINT;
    sequence_part BIGINT;
    random_part BIGINT;
BEGIN
    -- 获取当前时间戳（毫秒）
    timestamp_part := (extract(epoch from current_timestamp) * 1000)::BIGINT;
    
    -- 获取序列号
    sequence_part := nextval('"public".author_id_seq') % 512;
    
    -- 获取随机数部分
    random_part := nextval('"public".author_id_seq_seed') % 512;
    
    -- 组合TSID：41位时间戳 + 9位序列号 + 9位随机数
    RETURN (timestamp_part << 18) | (sequence_part << 9) | random_part;
END;
$$ LANGUAGE plpgsql;
-- ********
-- Table "author"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
    v_unique_constraint_name TEXT; 
    v_check_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'author') THEN
        -- 删除所有CHECK约束
        FOR v_check_constraint_name IN 
            SELECT conname
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'author'
                AND con.contype = 'c'
        LOOP
            EXECUTE 'ALTER TABLE "public"."author" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_check_constraint_name);
        END LOOP;

        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type 
            FROM information_schema.columns tbl 
            WHERE table_schema = 'public' 
            AND table_name = 'author' 
        LOOP
            IF column_rec.column_name NOT IN ('id','name') THEN
                EXECUTE 'ALTER TABLE "public"."author" DROP COLUMN IF EXISTS ' || 
                        quote_ident(column_rec.column_name) || ' CASCADE';
            END IF;
        END LOOP;

        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'author' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."author" ADD COLUMN "id" int8 NOT NULL DEFAULT author_id_seq();
        ELSE
            
            ALTER TABLE "public"."author" ALTER COLUMN "id" SET NOT NULL; 
            ALTER TABLE "public"."author" ALTER COLUMN "id" SET DEFAULT author_id_seq(); ALTER TABLE "public"."author" ALTER COLUMN "id" TYPE int8 USING "id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'author' AND column_name = 'name' ) THEN
            ALTER TABLE "public"."author" ADD COLUMN "name" varchar NOT NULL;
        ELSE
            
            ALTER TABLE "public"."author" ALTER COLUMN "name" SET NOT NULL; 
            ALTER TABLE "public"."author" ALTER COLUMN "name" DROP DEFAULT; ALTER TABLE "public"."author" ALTER COLUMN "name" TYPE varchar USING "name"::varchar;
        END IF;

        -- Search for existing unique and primary key constraints and drop them
        -- 查找并删除现有的唯一约束和主键约束
        BEGIN
            -- Drop primary key constraint
            -- 删除主键约束
            SELECT conname INTO v_constraint_name
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'author'
                AND con.contype = 'p';
            IF v_constraint_name IS NOT NULL THEN
                EXECUTE 'ALTER TABLE "public"."author" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name) || ' CASCADE';
            END IF;

            -- Drop unique constraints
            -- 删除唯一约束
            FOR v_unique_constraint_name IN 
                SELECT conname
                FROM pg_constraint con
                JOIN pg_class rel ON rel.oid = con.conrelid
                JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
                WHERE nsp.nspname = 'public'
                    AND rel.relname = 'author'
                    AND con.contype = 'u'
            LOOP
                BEGIN
                    EXECUTE 'ALTER TABLE "public"."author" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_unique_constraint_name);
                EXCEPTION WHEN OTHERS THEN
                    RAISE NOTICE 'Error dropping unique constraint %: %', v_unique_constraint_name, SQLERRM;
                END;
            END LOOP;
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'Error during dropping primary key or unique constraints: %', SQLERRM;
        END;

        -- 添加所有字段的CHECK约束
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."author" (
            "id" int8 NOT NULL DEFAULT author_id_seq(),
            "name" varchar NOT NULL
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."author"."id" IS  'Author primary key';
    

    -- Primary Key.
    -- 主键。
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'author'
                AND con.contype = 'p'
        ) THEN
            BEGIN
                ALTER TABLE "public"."author" ADD CONSTRAINT author_pkey PRIMARY KEY ("id");
            EXCEPTION 
                WHEN duplicate_table THEN
                    RAISE NOTICE 'Primary key constraint already exists';
                WHEN OTHERS THEN
                    RAISE NOTICE 'Error adding primary key constraint: %', SQLERRM;
            END;
        END IF;
    END;

    -- Add unique constraints
    -- 添加唯一约束
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding unique constraints: %', SQLERRM;
    END;

    -- Add indexes
    -- 添加索引
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding indexes: %', SQLERRM;
    END;
END
$$;

-- ********
-- Sequence blog_id_seq
-- ********
DO $$ 
BEGIN     
    -- 创建随机种子序列
    CREATE SEQUENCE IF NOT EXISTS "public"."blog_id_seq_seed"
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END $$;
CREATE OR REPLACE FUNCTION "public".blog_id_seq() 
RETURNS BIGINT AS $$
DECLARE
    timestamp_part BIGINT;
    sequence_part BIGINT;
    random_part BIGINT;
BEGIN
    -- 获取当前时间戳（毫秒）
    timestamp_part := (extract(epoch from current_timestamp) * 1000)::BIGINT;
    
    -- 获取序列号
    sequence_part := nextval('"public".blog_id_seq') % 512;
    
    -- 获取随机数部分
    random_part := nextval('"public".blog_id_seq_seed') % 512;
    
    -- 组合TSID：41位时间戳 + 9位序列号 + 9位随机数
    RETURN (timestamp_part << 18) | (sequence_part << 9) | random_part;
END;
$$ LANGUAGE plpgsql;
-- ********
-- Table "blog"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
    v_unique_constraint_name TEXT; 
    v_check_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'blog') THEN
        -- 删除所有CHECK约束
        FOR v_check_constraint_name IN 
            SELECT conname
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'blog'
                AND con.contype = 'c'
        LOOP
            EXECUTE 'ALTER TABLE "public"."blog" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_check_constraint_name);
        END LOOP;

        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type 
            FROM information_schema.columns tbl 
            WHERE table_schema = 'public' 
            AND table_name = 'blog' 
        LOOP
            IF column_rec.column_name NOT IN ('id','uuid','description','created_time') THEN
                EXECUTE 'ALTER TABLE "public"."blog" DROP COLUMN IF EXISTS ' || 
                        quote_ident(column_rec.column_name) || ' CASCADE';
            END IF;
        END LOOP;

        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "id" int8 NOT NULL DEFAULT blog_id_seq();
        ELSE
            
            ALTER TABLE "public"."blog" ALTER COLUMN "id" SET NOT NULL; 
            ALTER TABLE "public"."blog" ALTER COLUMN "id" SET DEFAULT blog_id_seq(); ALTER TABLE "public"."blog" ALTER COLUMN "id" TYPE int8 USING "id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'uuid' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "uuid" uuid NOT NULL;
        ELSE
            
            ALTER TABLE "public"."blog" ALTER COLUMN "uuid" SET NOT NULL; 
            ALTER TABLE "public"."blog" ALTER COLUMN "uuid" DROP DEFAULT; ALTER TABLE "public"."blog" ALTER COLUMN "uuid" TYPE uuid USING "uuid"::uuid;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'description' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "description" varchar;
        ELSE
            
            ALTER TABLE "public"."blog" ALTER COLUMN "description" DROP NOT NULL; 
            ALTER TABLE "public"."blog" ALTER COLUMN "description" DROP DEFAULT; ALTER TABLE "public"."blog" ALTER COLUMN "description" TYPE varchar USING "description"::varchar;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'blog' AND column_name = 'created_time' ) THEN
            ALTER TABLE "public"."blog" ADD COLUMN "created_time" timestamptz(6) DEFAULT CURRENT_TIMESTAMP;
        ELSE
            
            ALTER TABLE "public"."blog" ALTER COLUMN "created_time" DROP NOT NULL; 
            ALTER TABLE "public"."blog" ALTER COLUMN "created_time" SET DEFAULT CURRENT_TIMESTAMP; ALTER TABLE "public"."blog" ALTER COLUMN "created_time" TYPE timestamptz(6) USING "created_time"::timestamptz(6);
        END IF;

        -- Search for existing unique and primary key constraints and drop them
        -- 查找并删除现有的唯一约束和主键约束
        BEGIN
            -- Drop primary key constraint
            -- 删除主键约束
            SELECT conname INTO v_constraint_name
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'blog'
                AND con.contype = 'p';
            IF v_constraint_name IS NOT NULL THEN
                EXECUTE 'ALTER TABLE "public"."blog" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name) || ' CASCADE';
            END IF;

            -- Drop unique constraints
            -- 删除唯一约束
            FOR v_unique_constraint_name IN 
                SELECT conname
                FROM pg_constraint con
                JOIN pg_class rel ON rel.oid = con.conrelid
                JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
                WHERE nsp.nspname = 'public'
                    AND rel.relname = 'blog'
                    AND con.contype = 'u'
            LOOP
                BEGIN
                    EXECUTE 'ALTER TABLE "public"."blog" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_unique_constraint_name);
                EXCEPTION WHEN OTHERS THEN
                    RAISE NOTICE 'Error dropping unique constraint %: %', v_unique_constraint_name, SQLERRM;
                END;
            END LOOP;
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'Error during dropping primary key or unique constraints: %', SQLERRM;
        END;

        -- 添加所有字段的CHECK约束
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."blog" (
            "id" int8 NOT NULL DEFAULT blog_id_seq(),
            "uuid" uuid NOT NULL,
            "description" varchar,
            "created_time" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."blog"."id" IS  'Blog primary key';
    
    
    

    -- Primary Key.
    -- 主键。
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'blog'
                AND con.contype = 'p'
        ) THEN
            BEGIN
                ALTER TABLE "public"."blog" ADD CONSTRAINT blog_pkey PRIMARY KEY ("id");
            EXCEPTION 
                WHEN duplicate_table THEN
                    RAISE NOTICE 'Primary key constraint already exists';
                WHEN OTHERS THEN
                    RAISE NOTICE 'Error adding primary key constraint: %', SQLERRM;
            END;
        END IF;
    END;

    -- Add unique constraints
    -- 添加唯一约束
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding unique constraints: %', SQLERRM;
    END;

    -- Add indexes
    -- 添加索引
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding indexes: %', SQLERRM;
    END;
END
$$;

-- ********
-- Table "field_demo"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
    v_unique_constraint_name TEXT; 
    v_check_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'field_demo') THEN
        -- 删除所有CHECK约束
        FOR v_check_constraint_name IN 
            SELECT conname
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'field_demo'
                AND con.contype = 'c'
        LOOP
            EXECUTE 'ALTER TABLE "public"."field_demo" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_check_constraint_name);
        END LOOP;

        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type 
            FROM information_schema.columns tbl 
            WHERE table_schema = 'public' 
            AND table_name = 'field_demo' 
        LOOP
            IF column_rec.column_name NOT IN ('int64_f','var_f','bool_f','int_array_f','int_array2_f','string_array_f','bool_array_f','time_f','time_array_f','json_f') THEN
                EXECUTE 'ALTER TABLE "public"."field_demo" DROP COLUMN IF EXISTS ' || 
                        quote_ident(column_rec.column_name) || ' CASCADE';
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
            ALTER TABLE "public"."field_demo" ADD COLUMN "var_f" varchar NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "var_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "var_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "var_f" TYPE varchar USING "var_f"::varchar;
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
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'string_array_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "string_array_f" varchar[] NOT NULL;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "string_array_f" SET NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "string_array_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "string_array_f" TYPE varchar[] USING "string_array_f"::varchar[];
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
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'field_demo' AND column_name = 'json_f' ) THEN
            ALTER TABLE "public"."field_demo" ADD COLUMN "json_f" json;
        ELSE
            
            ALTER TABLE "public"."field_demo" ALTER COLUMN "json_f" DROP NOT NULL; 
            ALTER TABLE "public"."field_demo" ALTER COLUMN "json_f" DROP DEFAULT; ALTER TABLE "public"."field_demo" ALTER COLUMN "json_f" TYPE json USING "json_f"::json;
        END IF;

        -- Search for existing unique and primary key constraints and drop them
        -- 查找并删除现有的唯一约束和主键约束
        BEGIN
            -- Drop primary key constraint
            -- 删除主键约束
            SELECT conname INTO v_constraint_name
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'field_demo'
                AND con.contype = 'p';
            IF v_constraint_name IS NOT NULL THEN
                EXECUTE 'ALTER TABLE "public"."field_demo" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name) || ' CASCADE';
            END IF;

            -- Drop unique constraints
            -- 删除唯一约束
            FOR v_unique_constraint_name IN 
                SELECT conname
                FROM pg_constraint con
                JOIN pg_class rel ON rel.oid = con.conrelid
                JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
                WHERE nsp.nspname = 'public'
                    AND rel.relname = 'field_demo'
                    AND con.contype = 'u'
            LOOP
                BEGIN
                    EXECUTE 'ALTER TABLE "public"."field_demo" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_unique_constraint_name);
                EXCEPTION WHEN OTHERS THEN
                    RAISE NOTICE 'Error dropping unique constraint %: %', v_unique_constraint_name, SQLERRM;
                END;
            END LOOP;
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'Error during dropping primary key or unique constraints: %', SQLERRM;
        END;

        -- 添加所有字段的CHECK约束
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."field_demo" (
            "int64_f" int8 NOT NULL,
            "var_f" varchar NOT NULL,
            "bool_f" boolean NOT NULL,
            "int_array_f" int8[] NOT NULL,
            "int_array2_f" int8[][] NOT NULL,
            "string_array_f" varchar[] NOT NULL,
            "bool_array_f" boolean[] NOT NULL,
            "time_f" timestamptz(6) NOT NULL,
            "time_array_f" timestamptz[] NOT NULL,
            "json_f" json
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."field_demo"."int64_f" IS  'Int64 field';
    COMMENT ON COLUMN "public"."field_demo"."var_f" IS  'Varchar field';
    COMMENT ON COLUMN "public"."field_demo"."bool_f" IS  'Bool field';
    COMMENT ON COLUMN "public"."field_demo"."int_array_f" IS  'Int array field';
    COMMENT ON COLUMN "public"."field_demo"."int_array2_f" IS  'Int array2 field';
    COMMENT ON COLUMN "public"."field_demo"."string_array_f" IS  'String array field';
    COMMENT ON COLUMN "public"."field_demo"."bool_array_f" IS  'Bool array field';
    COMMENT ON COLUMN "public"."field_demo"."time_f" IS  'Time field';
    COMMENT ON COLUMN "public"."field_demo"."time_array_f" IS  'Time array field';
    COMMENT ON COLUMN "public"."field_demo"."json_f" IS  'Json field';

    -- Primary Key.
    -- 主键。
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'field_demo'
                AND con.contype = 'p'
        ) THEN
            BEGIN
                ALTER TABLE "public"."field_demo" ADD CONSTRAINT field_demo_pkey PRIMARY KEY ("int64_f");
            EXCEPTION 
                WHEN duplicate_table THEN
                    RAISE NOTICE 'Primary key constraint already exists';
                WHEN OTHERS THEN
                    RAISE NOTICE 'Error adding primary key constraint: %', SQLERRM;
            END;
        END IF;
    END;

    -- Add unique constraints
    -- 添加唯一约束
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding unique constraints: %', SQLERRM;
    END;

    -- Add indexes
    -- 添加索引
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding indexes: %', SQLERRM;
    END;
END
$$;

-- ********
-- Sequence geo_id_seq
-- ********
DO $$ 
BEGIN     
    -- 创建随机种子序列
    CREATE SEQUENCE IF NOT EXISTS "public"."geo_id_seq_seed"
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END $$;
CREATE OR REPLACE FUNCTION "public".geo_id_seq() 
RETURNS BIGINT AS $$
DECLARE
    timestamp_part BIGINT;
    sequence_part BIGINT;
    random_part BIGINT;
BEGIN
    -- 获取当前时间戳（毫秒）
    timestamp_part := (extract(epoch from current_timestamp) * 1000)::BIGINT;
    
    -- 获取序列号
    sequence_part := nextval('"public".geo_id_seq') % 512;
    
    -- 获取随机数部分
    random_part := nextval('"public".geo_id_seq_seed') % 512;
    
    -- 组合TSID：41位时间戳 + 9位序列号 + 9位随机数
    RETURN (timestamp_part << 18) | (sequence_part << 9) | random_part;
END;
$$ LANGUAGE plpgsql;
-- ********
-- Table "geo_demo"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
    v_unique_constraint_name TEXT; 
    v_check_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'geo_demo') THEN
        -- 删除所有CHECK约束
        FOR v_check_constraint_name IN 
            SELECT conname
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'geo_demo'
                AND con.contype = 'c'
        LOOP
            EXECUTE 'ALTER TABLE "public"."geo_demo" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_check_constraint_name);
        END LOOP;

        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type 
            FROM information_schema.columns tbl 
            WHERE table_schema = 'public' 
            AND table_name = 'geo_demo' 
        LOOP
            IF column_rec.column_name NOT IN ('id','point','line_string','polygon','multi_point','multi_line_string','multi_polygon','circular_string','point_json','line_string_json','polygon_json','multi_point_json','multi_line_string_json','multi_polygon_json') THEN
                EXECUTE 'ALTER TABLE "public"."geo_demo" DROP COLUMN IF EXISTS ' || 
                        quote_ident(column_rec.column_name) || ' CASCADE';
            END IF;
        END LOOP;

        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "id" int8 NOT NULL DEFAULT geo_id_seq();
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "id" SET NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "id" SET DEFAULT geo_id_seq(); ALTER TABLE "public"."geo_demo" ALTER COLUMN "id" TYPE int8 USING "id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'point' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "point" geometry(Point, 4326);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "point" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "point" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "point" TYPE geometry(Point, 4326) USING "point"::geometry(Point, 4326);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'line_string' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "line_string" geometry(LineString, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string" TYPE geometry(LineString, 0) USING "line_string"::geometry(LineString, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'polygon' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "polygon" geometry(Polygon, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "polygon" DROP NOT NULL; 
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
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'point_json' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "point_json" geometry(Point, 4326);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "point_json" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "point_json" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "point_json" TYPE geometry(Point, 4326) USING "point_json"::geometry(Point, 4326);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'line_string_json' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "line_string_json" geometry(LineString, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string_json" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string_json" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "line_string_json" TYPE geometry(LineString, 0) USING "line_string_json"::geometry(LineString, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'polygon_json' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "polygon_json" geometry(Polygon, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "polygon_json" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "polygon_json" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "polygon_json" TYPE geometry(Polygon, 0) USING "polygon_json"::geometry(Polygon, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'multi_point_json' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "multi_point_json" geometry(MultiPoint, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_point_json" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_point_json" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_point_json" TYPE geometry(MultiPoint, 0) USING "multi_point_json"::geometry(MultiPoint, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'multi_line_string_json' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "multi_line_string_json" geometry(MultiLineString, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_line_string_json" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_line_string_json" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_line_string_json" TYPE geometry(MultiLineString, 0) USING "multi_line_string_json"::geometry(MultiLineString, 0);
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'geo_demo' AND column_name = 'multi_polygon_json' ) THEN
            ALTER TABLE "public"."geo_demo" ADD COLUMN "multi_polygon_json" geometry(MultiPolygon, 0);
        ELSE
            
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_polygon_json" DROP NOT NULL; 
            ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_polygon_json" DROP DEFAULT; ALTER TABLE "public"."geo_demo" ALTER COLUMN "multi_polygon_json" TYPE geometry(MultiPolygon, 0) USING "multi_polygon_json"::geometry(MultiPolygon, 0);
        END IF;

        -- Search for existing unique and primary key constraints and drop them
        -- 查找并删除现有的唯一约束和主键约束
        BEGIN
            -- Drop primary key constraint
            -- 删除主键约束
            SELECT conname INTO v_constraint_name
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'geo_demo'
                AND con.contype = 'p';
            IF v_constraint_name IS NOT NULL THEN
                EXECUTE 'ALTER TABLE "public"."geo_demo" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name) || ' CASCADE';
            END IF;

            -- Drop unique constraints
            -- 删除唯一约束
            FOR v_unique_constraint_name IN 
                SELECT conname
                FROM pg_constraint con
                JOIN pg_class rel ON rel.oid = con.conrelid
                JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
                WHERE nsp.nspname = 'public'
                    AND rel.relname = 'geo_demo'
                    AND con.contype = 'u'
            LOOP
                BEGIN
                    EXECUTE 'ALTER TABLE "public"."geo_demo" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_unique_constraint_name);
                EXCEPTION WHEN OTHERS THEN
                    RAISE NOTICE 'Error dropping unique constraint %: %', v_unique_constraint_name, SQLERRM;
                END;
            END LOOP;
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'Error during dropping primary key or unique constraints: %', SQLERRM;
        END;

        -- 添加所有字段的CHECK约束
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."geo_demo" (
            "id" int8 NOT NULL DEFAULT geo_id_seq(),
            "point" geometry(Point, 4326),
            "line_string" geometry(LineString, 0),
            "polygon" geometry(Polygon, 0),
            "multi_point" geometry(MultiPoint, 0),
            "multi_line_string" geometry(MultiLineString, 0),
            "multi_polygon" geometry(MultiPolygon, 0),
            "circular_string" geometry(CircularString, 0),
            "point_json" geometry(Point, 4326),
            "line_string_json" geometry(LineString, 0),
            "polygon_json" geometry(Polygon, 0),
            "multi_point_json" geometry(MultiPoint, 0),
            "multi_line_string_json" geometry(MultiLineString, 0),
            "multi_polygon_json" geometry(MultiPolygon, 0)
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
    COMMENT ON COLUMN "public"."geo_demo"."point_json" IS  '点';
    COMMENT ON COLUMN "public"."geo_demo"."line_string_json" IS  '线';
    COMMENT ON COLUMN "public"."geo_demo"."polygon_json" IS  '多边形';
    COMMENT ON COLUMN "public"."geo_demo"."multi_point_json" IS  '多点';
    COMMENT ON COLUMN "public"."geo_demo"."multi_line_string_json" IS  '多线';
    COMMENT ON COLUMN "public"."geo_demo"."multi_polygon_json" IS  '多多边形';
    -- Table Comment.
    -- 表备注。
    COMMENT ON TABLE "public"."geo_demo" IS 'Geo的类型测试';

    -- Primary Key.
    -- 主键。
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'geo_demo'
                AND con.contype = 'p'
        ) THEN
            BEGIN
                ALTER TABLE "public"."geo_demo" ADD CONSTRAINT geo_demo_pkey PRIMARY KEY ("id");
            EXCEPTION 
                WHEN duplicate_table THEN
                    RAISE NOTICE 'Primary key constraint already exists';
                WHEN OTHERS THEN
                    RAISE NOTICE 'Error adding primary key constraint: %', SQLERRM;
            END;
        END IF;
    END;

    -- Add unique constraints
    -- 添加唯一约束
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding unique constraints: %', SQLERRM;
    END;

    -- Add indexes
    -- 添加索引
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding indexes: %', SQLERRM;
    END;
END
$$;

-- ********
-- Sequence post_id_seq
-- ********
DO $$ 
BEGIN     
    -- 创建随机种子序列
    CREATE SEQUENCE IF NOT EXISTS "public"."post_id_seq_seed"
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;
END $$;
CREATE OR REPLACE FUNCTION "public".post_id_seq() 
RETURNS BIGINT AS $$
DECLARE
    timestamp_part BIGINT;
    sequence_part BIGINT;
    random_part BIGINT;
BEGIN
    -- 获取当前时间戳（毫秒）
    timestamp_part := (extract(epoch from current_timestamp) * 1000)::BIGINT;
    
    -- 获取序列号
    sequence_part := nextval('"public".post_id_seq') % 512;
    
    -- 获取随机数部分
    random_part := nextval('"public".post_id_seq_seed') % 512;
    
    -- 组合TSID：41位时间戳 + 9位序列号 + 9位随机数
    RETURN (timestamp_part << 18) | (sequence_part << 9) | random_part;
END;
$$ LANGUAGE plpgsql;
-- ********
-- Table "post"
-- ********
DO $$
DECLARE
    column_rec RECORD;
    v_constraint_name TEXT;
    v_unique_constraint_name TEXT; 
    v_check_constraint_name TEXT;
BEGIN
    IF EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'post') THEN
        -- 删除所有CHECK约束
        FOR v_check_constraint_name IN 
            SELECT conname
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'post'
                AND con.contype = 'c'
        LOOP
            EXECUTE 'ALTER TABLE "public"."post" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_check_constraint_name);
        END LOOP;

        -- Check for any extra columns, and delete them if there are any.
        -- 检查是否有多余的列，如果有则删除。
        FOR column_rec IN SELECT tbl.column_name, tbl.data_type 
            FROM information_schema.columns tbl 
            WHERE table_schema = 'public' 
            AND table_name = 'post' 
        LOOP
            IF column_rec.column_name NOT IN ('id','content','blog_id','author_id') THEN
                EXECUTE 'ALTER TABLE "public"."post" DROP COLUMN IF EXISTS ' || 
                        quote_ident(column_rec.column_name) || ' CASCADE';
            END IF;
        END LOOP;

        -- Check for missing columns, and add them if any are missing.
        -- 检查是否缺少列，如果缺少则添加
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'post' AND column_name = 'id' ) THEN
            ALTER TABLE "public"."post" ADD COLUMN "id" int8 NOT NULL DEFAULT post_id_seq();
        ELSE
            
            ALTER TABLE "public"."post" ALTER COLUMN "id" SET NOT NULL; 
            ALTER TABLE "public"."post" ALTER COLUMN "id" SET DEFAULT post_id_seq(); ALTER TABLE "public"."post" ALTER COLUMN "id" TYPE int8 USING "id"::int8;
        END IF;
        IF NOT EXISTS (SELECT FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'post' AND column_name = 'content' ) THEN
            ALTER TABLE "public"."post" ADD COLUMN "content" varchar NOT NULL;
        ELSE
            
            ALTER TABLE "public"."post" ALTER COLUMN "content" SET NOT NULL; 
            ALTER TABLE "public"."post" ALTER COLUMN "content" DROP DEFAULT; ALTER TABLE "public"."post" ALTER COLUMN "content" TYPE varchar USING "content"::varchar;
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

        -- Search for existing unique and primary key constraints and drop them
        -- 查找并删除现有的唯一约束和主键约束
        BEGIN
            -- Drop primary key constraint
            -- 删除主键约束
            SELECT conname INTO v_constraint_name
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'post'
                AND con.contype = 'p';
            IF v_constraint_name IS NOT NULL THEN
                EXECUTE 'ALTER TABLE "public"."post" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_constraint_name) || ' CASCADE';
            END IF;

            -- Drop unique constraints
            -- 删除唯一约束
            FOR v_unique_constraint_name IN 
                SELECT conname
                FROM pg_constraint con
                JOIN pg_class rel ON rel.oid = con.conrelid
                JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
                WHERE nsp.nspname = 'public'
                    AND rel.relname = 'post'
                    AND con.contype = 'u'
            LOOP
                BEGIN
                    EXECUTE 'ALTER TABLE "public"."post" DROP CONSTRAINT IF EXISTS ' || quote_ident(v_unique_constraint_name);
                EXCEPTION WHEN OTHERS THEN
                    RAISE NOTICE 'Error dropping unique constraint %: %', v_unique_constraint_name, SQLERRM;
                END;
            END LOOP;
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'Error during dropping primary key or unique constraints: %', SQLERRM;
        END;

        -- 添加所有字段的CHECK约束
    ELSE
        -- If the table does not exist, then create the table.
        -- 如果表不存在，则创建表。
        CREATE TABLE "public"."post" (
            "id" int8 NOT NULL DEFAULT post_id_seq(),
            "content" varchar NOT NULL,
            "blog_id" int8 NOT NULL,
            "author_id" int8 NOT NULL
        );
    END IF;
    -- Field Comment.
    -- 字段备注。
    COMMENT ON COLUMN "public"."post"."id" IS  'Post primary key';
    
    
    

    -- Primary Key.
    -- 主键。
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint con
            JOIN pg_class rel ON rel.oid = con.conrelid
            JOIN pg_namespace nsp ON nsp.oid = rel.relnamespace
            WHERE nsp.nspname = 'public'
                AND rel.relname = 'post'
                AND con.contype = 'p'
        ) THEN
            BEGIN
                ALTER TABLE "public"."post" ADD CONSTRAINT post_pkey PRIMARY KEY ("id");
            EXCEPTION 
                WHEN duplicate_table THEN
                    RAISE NOTICE 'Primary key constraint already exists';
                WHEN OTHERS THEN
                    RAISE NOTICE 'Error adding primary key constraint: %', SQLERRM;
            END;
        END IF;
    END;

    -- Add unique constraints
    -- 添加唯一约束
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding unique constraints: %', SQLERRM;
    END;

    -- Add indexes
    -- 添加索引
    BEGIN
    EXCEPTION WHEN OTHERS THEN
        RAISE NOTICE 'Error during adding indexes: %', SQLERRM;
    END;
END
$$;





-- ********
-- Add Foreign Key
-- ********
DO $$
BEGIN

-- Check if principal table exists first
IF EXISTS (
    SELECT 1 FROM information_schema.tables 
    WHERE table_schema = 'public' 
    AND table_name = 'blog'
) THEN
    -- 判断是否存在唯一键，不存在添加
    IF NOT EXISTS (
        SELECT 1 
        FROM pg_constraint 
        WHERE conname = 'unique_blog_id' 
        AND conrelid = 'public.blog'::regclass
    ) THEN
        ALTER TABLE "public"."blog" 
        ADD CONSTRAINT unique_blog_id 
        UNIQUE (id);
    END IF;

    -- Add foreign key if dependent table exists
    IF EXISTS (
        SELECT 1 FROM information_schema.tables 
        WHERE table_schema = 'public' 
        AND table_name = 'post'
    ) THEN
        ALTER TABLE "public"."post"
        ADD CONSTRAINT fk_blog_id 
        FOREIGN KEY ("blog_id")
        REFERENCES "public"."blog" ("id");
    END IF;
END IF;

-- Check if principal table exists first
IF EXISTS (
    SELECT 1 FROM information_schema.tables 
    WHERE table_schema = 'public' 
    AND table_name = 'author'
) THEN
    -- 判断是否存在唯一键，不存在添加
    IF NOT EXISTS (
        SELECT 1 
        FROM pg_constraint 
        WHERE conname = 'unique_author_id' 
        AND conrelid = 'public.author'::regclass
    ) THEN
        ALTER TABLE "public"."author" 
        ADD CONSTRAINT unique_author_id 
        UNIQUE (id);
    END IF;

    -- Add foreign key if dependent table exists
    IF EXISTS (
        SELECT 1 FROM information_schema.tables 
        WHERE table_schema = 'public' 
        AND table_name = 'post'
    ) THEN
        ALTER TABLE "public"."post"
        ADD CONSTRAINT fk_author_id 
        FOREIGN KEY ("author_id")
        REFERENCES "public"."author" ("id");
    END IF;
END IF;

END
$$;

-- ********
-- Create Triggers
-- ********


