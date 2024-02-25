import { Link, useNavigate, useParams } from "react-router-dom";
import React, { useState, useEffect } from "react";
import "./Card.css";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
// import { Link } from 'react-router-dom';
import Header from "./Header";
import axios from "axios";
const Card = () => {
  const { username } = useParams();
  console.log(useParams());
  const navigate = useNavigate();
  const elements = ["cat", "cat", "shuffle", "bomb", "defuse"];
  const [shuffledElements, setShuffledElements] = useState([]);
  const [defuseActivated, setDefuseActivated] = useState(false);
  const [gameWon, setGameWon] = useState(false);
  const [lostGame, setLostGame] = useState(false);
  const[ele,setele]=useState("")
  const [isButtonVisible, setIsButtonVisible] = useState(false);

  const handleClickToShowButton = () => {
    setIsButtonVisible(true);
  };
  const handleClicked = async () => {
    try {
      const response = await axios.post(`http://localhost:8080/leaderboard`, {
        username: username,
        gameWon: gameWon ? 1 : 0,
        lostGame: lostGame ? 1 : 0,
      });
      if (response.status !== 200) {
        throw new Error("Failed to login");
      }
      const responseData = response.data;
      console.log("Login successful:", responseData);
      
    } catch (error) {
      console.error("login error", error);
    }
  };
  

  const handleClick = (element) => {
    setele(element)
    console.log(element);
    if (element === "cat" || element === "bomb" || element === "defuse") {
      const index = shuffledElements.indexOf(element);
      document.getElementsByClassName("card")[index].classList.add("clicked");
    }
    if (element === "cat") {
      const catIndex = shuffledElements.indexOf("cat");
      if (catIndex !== -1) {
        const newElements = [
          ...shuffledElements.slice(0, catIndex),
          ...shuffledElements.slice(catIndex + 1),
        ];
        setShuffledElements(newElements);
      }
    } else if (element === "shuffle") {
      const shuffledArray = [...elements];
      for (let i = shuffledArray.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [shuffledArray[i], shuffledArray[j]] = [
          shuffledArray[j],
          shuffledArray[i],
        ];
      }
      setShuffledElements(shuffledArray);
    } else if (element === "defuse") {
      const indexOfDefuse = shuffledElements.indexOf("defuse");
      if (indexOfDefuse !== -1) {
        const newElements = [
          ...shuffledElements.slice(0, indexOfDefuse),
          ...shuffledElements.slice(indexOfDefuse + 1),
        ];
        setShuffledElements(newElements);
      }
      setDefuseActivated(true);
    } else if (element === "bomb" && !defuseActivated) {
      setLostGame(true);
      setShuffledElements([]);
    } else if (element === "bomb" && defuseActivated) {
      const bombIndex = shuffledElements.indexOf("bomb");
      if (bombIndex !== -1) {
        const newElements = [
          ...shuffledElements.slice(0, bombIndex),
          ...shuffledElements.slice(bombIndex + 1),
        ];
        setShuffledElements(newElements);
      }
    }
  };

  useEffect(() => {
    const shuffledArray = [...elements];
    for (let i = shuffledArray.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [shuffledArray[i], shuffledArray[j]] = [
        shuffledArray[j],
        shuffledArray[i],
      ];
    }
    setShuffledElements(shuffledArray);
  }, []);

  useEffect(() => {
    if (shuffledElements.length === 1 && shuffledElements[0] === "shuffle") {
      setGameWon(true);
    }
  }, [shuffledElements]);
  useEffect(() => {
    if (gameWon) {
      toast.success("Congratulations! You won the game!", {
        position: "top-center",
        autoClose: 700,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
      });
      setTimeout(()=>{ handleClicked(); navigate('/leaderboard')},1000)
    }
    else if(lostGame) {
      toast.error("Sorry! You lost the game!", {
        position: "top-center",
        autoClose: 700,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
      });
      setTimeout(()=>{ handleClicked(); navigate('/leaderboard')},1000)
    }
    
  }, [gameWon, lostGame]);

  
  return (
    <div>
      <Header />
      <ToastContainer />
      <div className="Card-Container">
        {shuffledElements.map((element, index) => (
          <div
            key={index}
            className="card"
            onClick={() => handleClick(element)}
          >
            <button style={{visibility:"hidden"}}>{element}</button>
          </div>
        ))}
      </div>
      <h2 className="ele">You have clicked :{ele}</h2>
    </div>
      
  );
};

export default Card;
