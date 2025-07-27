document.addEventListener("DOMContentLoaded", () => {
    const path = window.location.pathname;

    if (path === "/") {
        // 登录页面逻辑
        const form = document.getElementById("login-form");
        form.addEventListener("submit", async (e) => {
            e.preventDefault();
            const username = document.getElementById("username").value;
            const password = document.getElementById("password").value;

            const res = await fetch("/api/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username, password })
            });

            if (res.ok) {
                window.location.href = "/dashboard";
            } else {
                const err = await res.json();
                document.getElementById("message").textContent = err.error || "Login failed";
            }
        });
    }

    if (path === "/dashboard") {
        // 主页面逻辑
        fetch("/api/profile").then(async (res) => {
            if (res.ok) {
                const data = await res.json();
                document.getElementById("welcome").textContent = `Welcome, ${data.username}!`;
            } else {
                window.location.href = "/";
            }
        });

        document.getElementById("logout").addEventListener("click", async () => {
            await fetch("/api/logout");
            window.location.href = "/";
        });
    }
});

