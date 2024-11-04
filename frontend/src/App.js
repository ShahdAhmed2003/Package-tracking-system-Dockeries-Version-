import React from 'react';
import RegistrationForm from './componets/RegisterationForm';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Login from './componets/LoginForm';
import OrderForm from './componets/OrderForm';
import MyOrders from './componets/MyOrders';
import OrderDetails from "./componets/OrderDetails";

function App() {
    return (
        <BrowserRouter>
         <Routes>
            <Route path="/" element={<RegistrationForm />} />
            <Route path="/login" element={<Login />} />
             <Route path="/order" element={<OrderForm />} />
             <Route path="/myOrders" element={<MyOrders />} />
             <Route path="/order/:orderId" element={<OrderDetails />} />
         </Routes>
        </BrowserRouter>
    );
}

export default App;
