# User Management Database

Shared database untuk **Auth Service** dan **User Service**.

## 📋 Database Tables

| Table   | Owner Service       | Description                        |
| ------- | ------------------- | ---------------------------------- |
| `users` | Auth & User Service | User data (email, password, token) |

## 🚀 Setup

### 1. Install golang-migrate

```bash
# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# macOS
brew install golang-migrate

# Windows
choco install golang-migrate
```

### 2. Configure Database

```bash
# Copy .env.example to .env
cp .env.example .env

# Edit .env and update with your MySQL credentials
nano .env
```

### 3. Create Database

```bash
# Login to MySQL
mysql -u root -p

# Create database
CREATE DATABASE user_management_db;
exit;
```

## 📝 Migration Commands

```bash
# Show available commands
make help

# Run all pending migrations
make migrate-up

# Rollback last migration
make migrate-down

# Drop all tables and re-run migrations (⚠️ DESTRUCTIVE!)
make migrate-fresh

# Check current migration version
make migrate-status

# Create new migration file
make migrate-create
```

## 🔐 Environment Variables

| Variable      | Description       | Example              |
| ------------- | ----------------- | -------------------- |
| `DB_HOST`     | Database host     | `localhost`          |
| `DB_PORT`     | Database port     | `3306`               |
| `DB_USER`     | Database user     | `root`               |
| `DB_PASSWORD` | Database password | `your_password`      |
| `DB_NAME`     | Database name     | `user_management_db` |

## 📁 Directory Structure

```
user-db/
├── migrations/
│   ├── 20251222082418_create_table_user.up.sql
│   └── 20251222082418_create_table_user.down.sql
├── .env                 # Your actual config (git ignored)
├── .env.example         # Template for other developers
├── .gitignore           # Ignore sensitive files
├── Makefile             # Migration commands
└── README.md            # This file
```

## 🔒 Service Access Rules

### Auth Service

- ✅ **READ/WRITE**: `users` table
- Use for: Login, Register, Token management

### User Service

- ✅ **READ/WRITE**: `users` table
- Use for: Profile management, User preferences

## ⚠️ Important Notes

1. **Never commit `.env`** - It contains sensitive credentials
2. **Always use migrations** - Don't modify database schema manually
3. **Coordinate changes** - If both Auth and User service need schema changes, discuss first
4. **Backup before fresh** - `make migrate-fresh` will delete all data!

## 🐛 Troubleshooting

### Error: "database connection failed"

- Check if MySQL is running: `sudo systemctl status mysql`
- Verify credentials in `.env` file
- Ensure database exists: `mysql -u root -p -e "SHOW DATABASES;"`

### Error: "Dirty database version"

- Run: `migrate -path migrations -database "mysql://..." force <version>`
- Then retry your migration

### Error: "no change"

- All migrations are already applied
- Check status: `make migrate-status`
