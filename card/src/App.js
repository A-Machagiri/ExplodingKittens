import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import LoginForm from "./Components/LoginForm";
import StartPage from "./Components/StartPage";
import Card from './Components/Card';
import Leaderboard from "./Components/Leaderboard";
function App() {
  return (
    <div className="app">
      <Router>
        <Routes>
          <Route exact path="/" element={<LoginForm />} />
          <Route path="/start/:username" element={<StartPage />} />
          <Route path="/submit/:username" element={<Card/>} />
          <Route path="/leaderboard" element={<Leaderboard/>} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
