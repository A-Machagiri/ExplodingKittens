import React from 'react';
import Header from './Header';
import { Link,useNavigate,useParams } from 'react-router-dom';
import './Start.css';
const StartPage = () => {
  const { username } = useParams();
  const navigate=useNavigate();
  const handleChange=()=>{
    navigate(`/submit/${username}`)
  }
  return (
    <div>
      <Header />
      <div className="start-page">
        <h2>This is the Start Page</h2>
        <button className="start-button" onClick={handleChange}>Start</button>
      </div>
    </div>
  );
}

export default StartPage;
