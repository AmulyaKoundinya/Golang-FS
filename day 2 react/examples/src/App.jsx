import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Greetings from './Greetings';
import Navbar from './Navbar';

function App() {
  

  return (
    <>
      <Navbar/>
      <h1><marquee>Welcome to REACT</marquee></h1>
      <Greetings/>
      
    </>
  );
}

export default App
