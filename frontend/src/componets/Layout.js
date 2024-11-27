
import React from 'react';
import Header from './Header';
import { Outlet } from 'react-router-dom'; // Outlet is where the child routes will be rendered

function Layout({ isLoggedIn, onLogout }) {
    return (
        <div>
            <Header isLoggedIn={isLoggedIn} onLogout={onLogout} />
            {/* The Outlet renders the current route */}
            <Outlet />
        </div>
    );
}

export default Layout;
