#!/bin/bash

# Exit script if any command fails
set -e

echo "========================================="
echo "  Setting up PostgreSQL on EC2 Instance"
echo "========================================="

# Update package lists
echo "[1/4] Updating package lists..."
sudo apt update -y

# Install PostgreSQL
echo "[2/4] Installing PostgreSQL..."
sudo apt install postgresql postgresql-contrib -y

# Ensure PostgreSQL service is started and enabled on boot
echo "[3/4] Starting PostgreSQL service..."
sudo systemctl enable postgresql
sudo systemctl start postgresql

# Set the password for the 'postgres' user to '4207'
echo "[4/4] Setting password for 'postgres' user..."
sudo -u postgres psql -c "ALTER USER postgres PASSWORD '4207';"

echo "========================================="
echo "  PostgreSQL Setup Complete!"
echo "  You can now start your Go backend."
echo "========================================="
