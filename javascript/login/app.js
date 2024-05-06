const express = require('express');
const session = require('express-session');
const bodyParser = require('body-parser');

const app = express();
const port = 3000;

// Middleware
app.use(bodyParser.urlencoded({ extended: true }));
app.use(session({
    secret: 'secret',
    resave: true,
    saveUninitialized: true
}));

// Middleware to check if user is logged in
const requireLogin = (req, res, next) => {
    if (!req.session.loggedin) {
        res.redirect('/login');
    } else {
        next();
    }
};

// Routes
app.get('/', requireLogin, (req, res) => {
    res.send(`
        <h1>Welcome back, ${req.session.username}!</h1>
        <form action="/logout" method="post">
            <button type="submit">Logout</button>
        </form>
    `);
});

app.get('/login', (req, res) => {
    if (req.session.loggedin) {
        res.redirect('/');
    } else {
        const errorMessage = req.session.errorMessage;
        req.session.errorMessage = ''; // Clear error message
        res.send(`
            <h1>Login Page</h1>
            ${errorMessage ? `<p style="color: red;">${errorMessage}</p>` : ''}
            <form action="/login" method="post">
                <input type="text" name="username" placeholder="Username" required><br>
                <input type="password" name="password" placeholder="Password" required><br>
                <button type="submit">Login</button>
            </form>
        `);
    }
});

app.post('/login', (req, res) => {
    const { username, password } = req.body;
    if (username && password && username === 'admin' && password === 'admin') {
        req.session.loggedin = true;
        req.session.username = username;
        res.redirect('/');
    } else {
        req.session.errorMessage = 'Incorrect username or password.';
        res.redirect('/login');
    }
});

app.post('/logout', (req, res) => {
    req.session.destroy(err => {
        if (err) {
            console.error('Error destroying session:', err);
        }
        res.redirect('/login');
    });
});

// Server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
