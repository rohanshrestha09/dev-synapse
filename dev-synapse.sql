CREATE TABLE models (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME
);

CREATE TABLE users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    bio TEXT,
    password VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    imageName VARCHAR(255),
    provider ENUM('GOOGLE', 'EMAIL') DEFAULT 'EMAIL' NOT NULL
);

CREATE TABLE projects (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    published BOOLEAN DEFAULT true NOT NULL,
    image VARCHAR(255),
    imageName VARCHAR(255),
    estimatedDuration INT UNSIGNED NOT NULL,
    startDate DATETIME,
    endDate DATETIME,
    status ENUM('OPEN', 'IN_PROGRESS', 'COMPLETED') DEFAULT 'OPEN' NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE requests (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    project_id INT UNSIGNED NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    status ENUM('APPROVED', 'PENDING', 'REJECTED') DEFAULT 'PENDING' NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE notifications (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    description VARCHAR(255) NOT NULL,
    initiator_id INT UNSIGNED NOT NULL,
    listener_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (initiator_id) REFERENCES users(id),
    FOREIGN KEY (listener_id) REFERENCES users(id)
);

CREATE TABLE developers (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    project_id INT UNSIGNED NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE chats (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    project_id INT UNSIGNED NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    message VARCHAR(255) NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
