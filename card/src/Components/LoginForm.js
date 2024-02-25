import React, { useState } from 'react';
import './Login.css';
import Header from './Header';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
const LoginForm = () => {
  const navigate = useNavigate(); // Use useNavigate hook for navigation
  const [data, setData] = useState({
    username: '',
    password: ''
  });

  const { username, password } = data;

  const changeHandler = e => {
    setData({ ...data, [e.target.name]: e.target.value });
  }
  const handleSubmit = async e => {
  e.preventDefault();
  try {
    const response = await axios.post(`http://localhost:8080/start`, data); // Use Axios to make POST request

      if (response.status !== 200) {
        throw new Error('Failed to login');
      }

      const responseData = response.data;
      console.log('Login successful:', responseData);
    navigate(`/start/${username}`);
  } catch (error) {
    console.error('Login error:', error);
  }
}


  return (
    <div>
      <Header />
      <div className="login-container">
        <center>
          <form onSubmit={handleSubmit}>
            <input type='text' placeholder='username' name='username' value={username} onChange={changeHandler} required/>
            <br />
            <input type='password' placeholder='password' name='password' value={password} onChange={changeHandler} required/>
            <br />
            <input type='submit' value='Submit' />
          </form>
        </center>
      </div>
    </div>
  );
}

export default LoginForm;
