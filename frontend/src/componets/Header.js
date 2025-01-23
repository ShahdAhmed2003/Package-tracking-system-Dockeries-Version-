import React from 'react';
import { Link } from 'react-router-dom';
import '../assets/styles.css';
import LogoutButton from './Log-out';

function Header({ isLoggedIn, onLogout }) {
    return (
        <header className="header">
            <div className="logo">
                <Link to="/" className="logo-link">Bosta</Link>
            </div>
            <nav className="nav-links">
                <Link to="/" className="nav-item">Home</Link>
                {isLoggedIn ? (
                    <>
                        <Link to="/order" className="nav-item">Create Order</Link>
                        <Link to="/my-orders" className="nav-item">My Orders</Link>
                        <LogoutButton onLogout={onLogout} />
                    </>
                ) : (
                    <>
                        <Link to="/register" className="login-button">Register</Link>
                        <Link to="/login" className="login-button">Login</Link>
                    </>
                )}
            </nav>
        </header>
    );
}

export default Header;
