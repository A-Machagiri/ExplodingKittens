import React, { useState, useEffect } from "react";
import "./Leaderboard.css";
import Header from "./Header";
import axios from "axios";

const Leaderboard = () => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(()=>{
    (async ()=>{
      try {
        const response = await axios.get("http://localhost:8080/leaderboard-desc");
        console.log(response.data);
        setUsers(response.data);
        setLoading(false); // Set loading to false after data is fetched
      } catch (error) {
        console.error(error);
        setError("Error fetching data."); // Set error state if there's an error
        setLoading(false); // Set loading to false if there's an error
      }
    })();
  },[]);

  useEffect(() => {
    console.log(users);
  }, [users]);

  if (loading) {
    return (
      <div className="middlecontent">
        <h1 className="loading-text">Loading</h1>
      </div>
    );
  }

  if (error) {
    return (
      <div className="middlecontent">
        <div className="error-text">Error: {error}</div>
      </div>
    );
  }

  return (
    <div>
      <Header />
      <div className="leaderboard-container">
        <h2>Leaderboard</h2>
        <table>
          <thead>
            <tr>
              <th>Username</th>
              <th>Wins</th>
              <th>Losses</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user, index) => (
              <tr key={index}>
                <td>{user.username}</td>
                <td>{user.wins}</td>
                <td>{user.losses}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Leaderboard;
