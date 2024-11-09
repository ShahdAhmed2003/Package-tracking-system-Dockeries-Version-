import React, { useState, useEffect } from 'react';
import { useParams,useNavigate } from 'react-router-dom';
import "../assets/styless.css";
import { decodeToken } from "../utils/auth";
import LogoutButton from './Log-out';
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
        setMessage(fetchedOrders.message);  // Set the "No assigned orders found" message
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
  
 const AcceptOrder = async (orderId) => {
        
            const response = await fetch(`http://localhost:8080/api/orders/verify?orderId=${orderId}`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.ok) {
                alert("Order Accepted successfully");
                
            } else {
                const errorData = await response.text();
                alert("Failed to accept order: " + errorData);
            }
            window.location.reload();
    };
    const DeclineOrder=async(orderId)=>{
      const response=await fetch(`http://localhost:8080/api/orders/decline/${orderId}`,{
       method:'DELETE',
       headers: {
        'Authorization': `Bearer ${token}`
    }

      })
      if (response.ok) {
        alert("Order Declined");
        
    } else {
        const errorData = await response.text();
        alert("Failed to decline order: " + errorData);
    }
    window.location.reload();
    };
  return (
    <div>
      <div><LogoutButton onLogout={onLogout} /> </div>
      <h1>Orders for Courier:{userID} </h1>
      {message ? (
        // If there's a message (no assigned orders), display it
        <p>{message}</p>
      ) : (
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
                  {order.status === "pending" && (
                    <>
                      <button onClick={() => AcceptOrder(order.ID)}>Accept</button>
                      <button className="refuse-button" onClick={() => DeclineOrder(order.ID)}>Refuse</button>
                    </>
                  )}
                  {order.status === "accepted" && <span>Already Accepted</span>}
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      )}
      <button onClick={() => navigate('/courier/manageOrders')}>Update Orders</button>
    </div>
  );
};

export default OrdersPage;
