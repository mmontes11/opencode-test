 The program is a minimal HTTP server that connects to a MariaDB database.
  - main.go – starts the server on port 8080. It initializes the DB connection (db.Init()), creates a
  router (router.NewRouter()), and then listens for HTTP requests.
  - db/db.go – sets up a global *sql.DB that points to a MariaDB instance. The DSN is hard‑coded but can be overridden with MARIADB_DSN. It opens the connection and pings it to verify connectivity.
  - router/router.go – uses gorilla/mux to create a router. The only route defined is /health, which is
  handled by handler.HealthCheck.
  - handler/handler.go – implements the health‑check handler, writing a JSON response {"status":"ok"}.

  So when you run go run ., you get a server that reports "status":"ok" at http://localhost:8080/health
  and logs any DB or server errors.
