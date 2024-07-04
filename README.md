<p align="center"><a href="https://laravel.com" target="_blank"><img src="https://raw.githubusercontent.com/laravel/art/master/logo-lockup/5%20SVG/2%20CMYK/1%20Full%20Color/laravel-logolockup-cmyk-red.svg" width="400" alt="Laravel Logo"></a></p>

# 🏢 Job Jive - Human Resource Management System

The HR department in a company is responsible for managing employees, including attendance, storing employee biodata, and making decisions regarding employees. When a company has a large number of employees, manual administration can significantly increase the workload. Therefore, we developed **Job Jive** as a Human Resource Management System to streamline these tasks.

## ✨ Features

In this application, we developed five main features:

1. **Authentication and Authorization**: Secure login and role-based access control for manager and staff accounts.
2. **Employee Attendance Tracking**: Efficiently track employee attendance and manage records.
3. **Employee Data Management**: Store and manage employee biodata, job details, and other relevant information.
4. **Chart & Stats**: Visualize data and metrics through interactive charts for better decision-making.
5. **Motivation Letter Generate by AI**:AI generating personalized motivation letters.

## 💻 Technologies Used

### Front-end

- **Framework**: PHP Laravel with Blade templating engine.
- **CSS**: Tailwind CSS for responsive and modern design.
- **UI Components**: Integrated components from Flowbite for enhanced UI/UX.

### Back-end

- **Language**: GOLANG for robust and efficient back-end processes.
- **Features**: Handling authentication, authorization, web server, database, and all back-end funtionality.

## 📌 Installation Instructions

To set up the application locally, follow these steps:

### Prerequisites
- PHP 8.2 or higher
- Golang
- Composer
- Node.js and npm
- MySQL server

## Front-End Installation

1. **Clone the repository:**
   ```
   git clone https://github.com/bryanfks-dev/Job-Jive.git
   cd Job-Jive/laravel-app/
   ```

2. **Install PHP dependencies:**
    ```
    composer install
    ```

3. **Install JavaScript dependencies:**
    ```
    npm install
    ```

4. **Set up the environment file:**
    ```
    cp .env.example .env
    ```

5. **Generate application key:**
    ```
    php artisan key:generate
    ```

6. **Run migrations and seed the database:**
    ```
    php artisan migrate --seed
    ```

## Back-End Installation

7. **Copy BE environment variable**
    ```
   # Go to root dir
    
    cd go-worker
    copy .env.example .env
    ```

8. **Build addAdmin.exe**
    ```  
    go build -o addAdmin.exe ./utils/addAdmin/main.go
    ```

9. **Create admin account**
    ```
    ./addAdmin <your-username> <your-password>
    ```

## Run the program

10. **Start the Back-End server**
    ```
    go run .
    ```

11. **Build assets**
    ```
    # Go to root folder

    cd laravel-app
    npm run dev
    ```

12. **Start the development server:**
    ```
    php artisan serve
    ```

13. **Access the application:**
Open your web browser and navigate to 'http://localhost:8000'
