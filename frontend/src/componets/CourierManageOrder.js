import React, { useState, useEffect } from 'react';
import { useParams,useNavigate } from 'react-router-dom';
import LogoutButton from './Log-out';
import { decodeToken } from "../utils/auth";
const OrdersPage = ({onLogout}) => {
  const [orders, setOrders] = useState([]);
   const token=localStorage.getItem("token");
   const [error, setError] = useState(null);
   const [message, setMessage] = useState(""); 
    // Add more orders as needed
    let userID=null;
    if(token){
      const decodedToken=decodeToken(token);
      userID=decodedToken.userID;
  }
  const navigate = useNavigate(); 

  useEffect(() => {
    const fetchOrdersForCourier = async () => {
      // Replace this with your actual data fetching logic
      try{

      const response = await fetch("http://localhost:8080/api/orders/assigned_orders",
      {method:'GET',
      headers:{
        'Authorization':`Bearer ${token}`
      }
      });
      if (!response.ok) {
        throw new Error('Failed to fetch orders');
      }
      const fetchedOrders = await response.json();
      console.log(fetchedOrders);
      if(fetchedOrders.message)
      {
        setMessage(fetchedOrders.message);  
        setOrders([]);
      }
      else{
      setOrders(fetchedOrders);
      setMessage(""); 
      }
      
    } catch(err)
    {
      setError(err.message);
    };

  };
  fetchOrdersForCourier();
},[token]);
  

  /*const updateOrderStatus = (id, newStatus) => {
    setOrders(orders.map(order => 
      order.id === id ? { ...order, status: newStatus } : order
    ));
  };*/
  const updateOrderStatus = async (orderId,newStatus) => {
        
    const response = await fetch(`http://localhost:8080/api/orders/update-order-status/${orderId}`, {
        method: 'PUT',
        body: JSON.stringify({ status: newStatus }),
        headers: {
            'Authorization': `Bearer ${token}`
        }
    });

    if (response.ok) {
        alert("Order Updated successfully");
        
    } else {
        const errorData = await response.text();
        alert("Failed update order: " + errorData);
    }
    window.location.reload();
};
  return (
    <div>
      <div><LogoutButton onLogout={onLogout} /> </div>
      <h1>Orders for Courier:{userID}</h1>
      <table>
        <thead>
          <tr>
            <th>Order ID</th>
            <th>Pick-up location</th>
            <th>Drop-off location</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {orders.map(order => (
            <tr key={order.ID}>
              <td>{order.ID}</td>
              <td>{order.pickup_location.street_address}</td>
              <td>{order.drop_off_location.street_address}</td>
              <td>{order.status}</td>
              <td>
                <div className="button-group">
                  {order.status === "accepted" && (
                    <button className='manage-button' onClick={() => updateOrderStatus(order.ID, "picked up")}>Picked Up</button>
                  )}
                  {order.status === "picked up" && (
                    <button className='manage-button' onClick={() => updateOrderStatus(order.ID, "in transit")}>In Transit</button>
                  )}
                  {order.status === "in transit" && (
                    <button className='manage-button' onClick={() => updateOrderStatus(order.ID, "delivered")}>Delivered</button>
                  )}
                  {order.status === "delivered" && <span>Delivered</span>}
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <button onClick={() => navigate('/courier/assignedOrders')}>Assigned Orders</button>
    </div>
  );
};

export default OrdersPage;
