# README: Setting Up Docker and PostgreSQL on Windows

## Prerequisites
- **Operating System**: Windows 10/11
- **Administrator Privileges**

---

## Step 1: Install Docker

1. **Download Docker Desktop**:
   - Visit the official Docker website: [Docker Desktop for Windows](https://www.docker.com/products/docker-desktop/).
   - Click **Download Docker Desktop**.

2. **Install Docker**:
   - Run the installer.
   - Follow the installation wizard.
   - Ensure that the "Enable WSL 2" option is selected.

3. **Verify Installation**:
   - Open **PowerShell** or **Command Prompt**.
   - Run:
     ```bash
     docker --version
     ```
   - You should see the installed Docker version.

4. **Start Docker**:
   - Open Docker Desktop and ensure it is running.

---

## Step 2: Install PostgreSQL Using Docker

1. **Pull PostgreSQL Image**:
   - Open **PowerShell** or **Command Prompt**.
   - Run:
     ```bash
     docker pull postgres:15
     ```

2. **Run PostgreSQL Container**:
   - Start a PostgreSQL container:
     ```bash
     docker run --name postgres_container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=your_password -e POSTGRES_DB=your_db -p 5432:5432 -d postgres:15
     ```
     - Replace `your_password` with your desired password.
     - Replace `your_db` with the name of your database.

3. **Verify PostgreSQL is Running**:
   - List running containers:
     ```bash
     docker ps
     ```
   - You should see your PostgreSQL container in the list.

4. **Access PostgreSQL**:
   - Use a PostgreSQL client (e.g., pgAdmin or `psql`) to connect to the database:
     - **Host**: `localhost`
     - **Port**: `5432`
     - **Username**: `postgres`
     - **Password**: Your specified password

---

## Notes
- Ensure Docker Desktop is running before starting the PostgreSQL container.
- Use `docker stop postgres_container` to stop the container when not in use.
- Use `docker start postgres_container` to start it again later.

For additional help, refer to the [Docker Documentation](https://docs.docker.com/) or the [PostgreSQL Documentation](https://www.postgresql.org/docs/).
