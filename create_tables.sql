-- 启用UUID生成扩展
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- 1. 用户表
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(32) UNIQUE NOT NULL,
    password_hash VARCHAR(128) NOT NULL,
    nickname VARCHAR(32),
    avatar_url TEXT,
    bio VARCHAR(128),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 2. 路线表
CREATE TABLE routes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(64) NOT NULL,
    description TEXT,
    days INT NOT NULL,
    start_city VARCHAR(32),
    end_city VARCHAR(32),
    is_hot BOOLEAN DEFAULT FALSE,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 3. 行程节点表
CREATE TABLE route_nodes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    route_id UUID REFERENCES routes(id) ON DELETE CASCADE,
    day INT NOT NULL,
    order_in_day INT NOT NULL,
    name VARCHAR(64) NOT NULL,
    type VARCHAR(32),
    location VARCHAR(128),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    start_time TIME,
    end_time TIME,
    open_time TIME,
    close_time TIME,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 4. 用户行程表
CREATE TABLE user_trips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    route_id UUID REFERENCES routes(id) ON DELETE CASCADE,
    start_date DATE NOT NULL,
    status VARCHAR(16) DEFAULT 'ongoing',
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 5. 用户行程节点进度表
CREATE TABLE user_trip_nodes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_trip_id UUID REFERENCES user_trips(id) ON DELETE CASCADE,
    route_node_id UUID REFERENCES route_nodes(id) ON DELETE CASCADE,
    status VARCHAR(16) DEFAULT 'pending',
    checkin_time TIMESTAMPTZ,
    comment TEXT,
    rating INT CHECK (rating >= 0 AND rating <= 5),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 6. 动态/社区表
CREATE TABLE moments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    content TEXT,
    image_urls JSONB,
    location VARCHAR(128),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 7. 评论表
CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    moment_id UUID REFERENCES moments(id) ON DELETE CASCADE,
    content TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- 8. 关注表
CREATE TABLE follows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    follower_id UUID REFERENCES users(id) ON DELETE CASCADE,
    followee_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE (follower_id, followee_id)
);

-- 9. 用户定位历史表（可选）
CREATE TABLE user_locations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    recorded_at TIMESTAMPTZ DEFAULT now(),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
); 