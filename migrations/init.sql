-- ERP Finance Module - Database Init
-- Run automatically by Docker on first start

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(50) DEFAULT 'viewer',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS budgets (
    id SERIAL PRIMARY KEY,
    department VARCHAR(100) NOT NULL,
    fiscal_year INT NOT NULL,
    allocated_amount NUMERIC(15,2) NOT NULL,
    spent_amount NUMERIC(15,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS expenses (
    id SERIAL PRIMARY KEY,
    budget_id INT REFERENCES budgets(id),
    description TEXT NOT NULL,
    amount NUMERIC(15,2) NOT NULL,
    category VARCHAR(100),
    status VARCHAR(50) DEFAULT 'pending',
    submitted_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS invoices (
    id SERIAL PRIMARY KEY,
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    vendor_name VARCHAR(100) NOT NULL,
    amount NUMERIC(15,2) NOT NULL,
    due_date DATE,
    status VARCHAR(50) DEFAULT 'unpaid',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS payroll (
    id SERIAL PRIMARY KEY,
    employee_id VARCHAR(50) NOT NULL,
    employee_name VARCHAR(100) NOT NULL,
    department VARCHAR(100),
    basic_salary NUMERIC(15,2) NOT NULL,
    allowances NUMERIC(15,2) DEFAULT 0,
    deductions NUMERIC(15,2) DEFAULT 0,
    net_salary NUMERIC(15,2) GENERATED ALWAYS AS (basic_salary + allowances - deductions) STORED,
    pay_period VARCHAR(20),
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW()
);

-- Sample Data
INSERT INTO users (name, email, password_hash, role) VALUES
('Admin User', 'admin@erp.com', '$2a$10$examplehashedpassword', 'admin'),
('Finance Manager', 'finance@erp.com', '$2a$10$examplehashedpassword', 'manager');

INSERT INTO budgets (department, fiscal_year, allocated_amount) VALUES
('Engineering', 2025, 500000.00),
('Marketing', 2025, 200000.00),
('HR', 2025, 150000.00),
('Operations', 2025, 300000.00);

INSERT INTO expenses (budget_id, description, amount, category, status, submitted_by) VALUES
(1, 'Server infrastructure upgrade', 45000.00, 'Infrastructure', 'approved', 1),
(1, 'Software licenses', 12000.00, 'Software', 'pending', 2),
(2, 'Marketing campaign Q1', 30000.00, 'Advertising', 'approved', 2),
(3, 'Team training workshop', 8000.00, 'Training', 'pending', 1);

INSERT INTO invoices (invoice_number, vendor_name, amount, due_date, status) VALUES
('INV-2025-001', 'AWS Cloud Services', 8500.00, '2025-06-30', 'unpaid'),
('INV-2025-002', 'Office Supplies Co.', 1200.00, '2025-05-15', 'paid'),
('INV-2025-003', 'Consulting Firm Ltd.', 25000.00, '2025-07-01', 'unpaid');

INSERT INTO payroll (employee_id, employee_name, department, basic_salary, allowances, deductions, pay_period, status) VALUES
('EMP001', 'John Silva', 'Engineering', 85000.00, 5000.00, 8500.00, '2025-04', 'processed'),
('EMP002', 'Amara Perera', 'Marketing', 65000.00, 3000.00, 6500.00, '2025-04', 'processed'),
('EMP003', 'Kamal Fernando', 'HR', 55000.00, 2500.00, 5500.00, '2025-04', 'pending');
