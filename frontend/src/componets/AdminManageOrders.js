import React, { useState, useEffect } from 'react';
import "../assets/styless.css";
import LogoutButton from './Log-out';

const AdminOrdersPage = ({ onLogout }) => {
  const [orders, setOrders] = useState([]);
  const [couriers, setCouriers] = useState([]);
  const [error, setError] = useState(null);

  // Fetch orders and couriers when the component loads
  useEffect(() => {
    // Fetch orders
    const fetchOrders = async () => {
      try {
        const response = await fetch("https://backend-shahdahmed851-dev.apps.rm2.thpm.p1.openshiftapps.com/api/orders/admin", {
          method: 'GET'
        });

        if (!response.ok) {
          throw new Error('Failed to fetch orders');
        }

        const fetchedOrders = await response.json();
        setOrders(fetchedOrders);
      } catch (err) {
        setError("Failed to fetch orders.");
      }
    };

    // Fetch couriers
    const fetchCouriers = async () => {
      try {
        const response = await fetch("https://backend-shahdahmed851-dev.apps.rm2.thpm.p1.openshiftapps.com/api/couriers", {
          method: 'GET'
        });

        if (!response.ok) {
          throw new Error('Failed to fetch couriers');
        }

        const fetchedCouriers = await response.json();
        setCouriers(fetchedCouriers);
      } catch (err) {
        setError("Failed to fetch couriers.");
      }
    };

    fetchOrders();
    fetchCouriers();
  
  }, []); 

  // Assign courier to an order
  const assignCourier = async (orderId, courierId) => {
    try {
      const response = await fetch(`https://backend-shahdahmed851-dev.apps.rm2.thpm.p1.openshiftapps.com/api/orders/assign/${orderId}/${courierId}`, {
        method: 'POST', // Assuming the method is POST for assigning courier
      });

      if (!response.ok) {
        throw new Error("Failed to assign courier.");
      }

      // Refresh the page after the assignment
      window.location.reload(); // This will reload the page
    } catch (err) {
      setError(err.message);
    }
  };

  // Update the status of an order
  const updateOrderStatus = async (orderId, newStatus) => {
    try {
      const response = await fetch(`https://backend-shahdahmed851-dev.apps.rm2.thpm.p1.openshiftapps.com/api/orders/admin-update/${orderId}`, {
        method: 'PUT', 
        body: JSON.stringify({ status: newStatus }),
        headers: { 'Content-Type': 'application/json' },
      });

      if (!response.ok) {
        throw new Error("Failed to update order status.");
      }
      window.location.reload();

    } catch (err) {
      setError(err.message);
    }
  };

  // Delete an order
  const deleteOrder = async (orderId) => {
    try {
      const response = await fetch(`https://backend-shahdahmed851-dev.apps.rm2.thpm.p1.openshiftapps.com/api/orders/admin-delete/${orderId}`, {
        method: 'DELETE',
      });

      if (!response.ok) {
        throw new Error("Failed to delete the order.");
      }

      window.location.reload();
    } catch (err) {
      setError(err.message);
    }
  };

  // List of available statuses
  const statusOptions = ["pending", "accepted", "picked up", "in transit", "delivered"];

  return (
    <div>
      <h1>Admin Orders Management</h1>
      <div><LogoutButton onLogout={onLogout} /> </div>
      {error && <p>{error}</p>}
      <table>
        <thead>
          <tr>
            <th>Order ID</th>
            <th>Pick-up Location</th>
            <th>Drop-off Location</th>
            <th>Current Status</th>
            <th>Change Status</th>
            <th>Assign Courier</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {orders.map((order) => (
            <tr key={order.ID}>
              <td>{order.ID}</td>
              <td>{order.pickup_location.street_address}</td>
              <td>{order.drop_off_location.street_address}</td>
              <td>{order.status.charAt(0).toUpperCase() + order.status.slice(1)}</td>
              <td>
                <select
                  value={order.status}
                  onChange={(e) => updateOrderStatus(order.ID, e.target.value)}
                >
                  {statusOptions.map(status => (
                    <option key={status} value={status}>
                      {status.charAt(0).toUpperCase() + status.slice(1)}
                    </option>
                  ))}
                </select>
              </td>
              <td>
                <select
                  value={order.courier_id || ""}
                  onChange={(e) => assignCourier(order.ID, parseInt(e.target.value))}
                >
                  <option value="">Unassigned</option>
                  {couriers.map((courier) => (
                    <option key={courier.ID} value={courier.ID}>
                      {courier.ID}
                    </option>
                  ))}
                </select>
              </td>
              <td>
                <button className='refuse-button' onClick={() => deleteOrder(order.ID)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default AdminOrdersPage;
