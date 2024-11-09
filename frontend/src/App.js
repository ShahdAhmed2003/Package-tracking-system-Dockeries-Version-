// App.js
import React, { useState, useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Header from './componets/Header';
import Home from './Pages/Home';
import LoginForm from './componets/LoginForm';
import RegistrationForm from './componets/RegisterationForm';
import OrderForm from './componets/OrderForm';
import MyOrders from './componets/MyOrders';
import OrderDetails from './componets/OrderDetails';
import AssignedOrders from'./componets/CourierAssignedOrders';
import ManageOrders from './componets/CourierManageOrder';
import AdminManageOrders from './componets/AdminManageOrders';
import Layout from './componets/Layout';


function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    useEffect(() => {
        const token = localStorage.getItem("token");
        const user = localStorage.getItem("user");
        setIsLoggedIn(!!(token && user));
    }, []);

    const handleLoginStatusChange = () => {
        const token = localStorage.getItem("token");
        const user = localStorage.getItem("user");
        setIsLoggedIn(!!(token && user));
    };

const handleLogout = () => {
        localStorage.clear();
        setIsLoggedIn(false);
    };
    return (
        <BrowserRouter>
           
            <Routes>
            <Route path="/" element={<Layout isLoggedIn={isLoggedIn} onLogout={handleLogout} />}>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<LoginForm onLogin={handleLoginStatusChange} />} />
                <Route path="/register" element={<RegistrationForm />} />
                <Route path="/order" element={<OrderForm />} />
                <Route path="/my-orders" element={<MyOrders />} />
                <Route path="/order-details/:orderId" element={<OrderDetails />} />

            </Route>
              {/* Routes that don't need the Header */}
            <Route path="/courier/manageOrders" element={<ManageOrders onLogout={handleLogout} />} />
            <Route path="/courier/assignedOrders" element={<AssignedOrders onLogout={handleLogout} />} />
            <Route path="/admin/manageOrders" element={<AdminManageOrders onLogout={handleLogout}/>} />
            </Routes>
            
        </BrowserRouter>
    );
}

export default App;