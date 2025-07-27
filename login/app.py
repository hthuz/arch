from flask import Flask, request, jsonify, session, send_from_directory
import sqlite3
import datetime
import os

app = Flask(__name__, static_folder='static', static_url_path='')
app.secret_key = 'secret_key'
DATABASE = 'db.sqlite3'

def get_db():
    db = sqlite3.connect(DATABASE)
    db.row_factory = sqlite3.Row
    return db


@app.route('/api/login', methods=['POST'])
def login():
    data = request.get_json()
    username = data.get('username')
    password = data.get('password')

    db = get_db()
    user = db.execute("SELECT * FROM users WHERE username=? AND password=?", 
                      (username, password)).fetchone()

    if user:
        session['username'] = username
        db.execute("INSERT INTO login_history (username, login_time, ip) VALUES (?, ?, ?)",
                   (username, datetime.datetime.now(), request.remote_addr))
        db.commit()
        return jsonify({'success': True})
    else:
        return jsonify({'success': False, 'error': 'Invalid credentials'}), 401


@app.route('/api/logout')
def logout():
    session.clear()
    return jsonify({'success': True})


@app.route('/api/profile')
def profile():
    if 'username' not in session:
        return jsonify({'error': 'Not logged in'}), 401
    return jsonify({'username': session['username']})


# Serve static HTML files
@app.route('/')
def root():
    return send_from_directory('static', 'login.html')

@app.route('/dashboard')
def dashboard():
    return send_from_directory('static', 'dashboard.html')


if __name__ == '__main__':
    app.run(debug=True)

