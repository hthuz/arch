import sqlite3
conn = sqlite3.connect('db.sqlite3')
c = conn.cursor()

c.execute('''
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
)
''')

c.execute('''
CREATE TABLE login_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    login_time TIMESTAMP,
    ip TEXT
)
''')

c.execute("INSERT INTO users (username, password) VALUES (?, ?)", ('admin', '123456'))
conn.commit()
conn.close()

