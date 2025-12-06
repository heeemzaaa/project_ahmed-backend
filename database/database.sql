
CREATE TABLE IF NOT EXISTS user (
    id TEXT PRIMARY KEY UNIQUE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    centre TEXT NOT NULL,
    filiere TEXT NOT NULL,
    year TEXT DEFAULT "",
    access_premiere_annees BOOLEAN DEFAULT 0,
    access_deuxieme_annees BOOLEAN DEFAULT 0,
    access_concours_francais BOOLEAN DEFAULT 0,
    access_concours_maroc BOOLEAN DEFAULT 0,
    password TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    expiration_of_access DATETIME,
    isAdmin BOOLEAN DEFAULT 0,
    is_google BOOLEAN DEFAULT 0,
    confirmation_code TEXT,
    is_confirmed BOOLEAN DEFAULT 0
);

CREATE TABLE IF NOT EXISTS user_sessions (
    id TEXT PRIMARY KEY UNIQUE,
    user_id TEXT NOT NULL,
    device_id TEXT NOT NULL,
    device_type CHECK(device_type IN ('desktop', 'mobile')) NOT NULL,
    refresh_token TEXT NOT NULL,
    last_used_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_devices (
    id TEXT PRIMARY KEY UNIQUE,
    user_id TEXT NOT NULL,
    device_type CHECK(device_type IN ('desktop', 'mobile')) NOT NULL,
    device_Id TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS videos (
    id TEXT PRIMARY KEY UNIQUE,
    title TEXT NOT NULL,
    category TEXT NOT NULL, 
    vdocipherVideoId TEXT NOT NULL,
    orderIndex INTEGER NOT NULL,
    folderOrder INTEGER
);
