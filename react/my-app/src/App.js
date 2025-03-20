import React from "react";
import Form from "./Form";
import logo from "./logo.svg";
import "./App.css";

const App = () => {
  // handleSubmitを定義（name, email を受け取る）
  const handleSubmit = (name, email) => {
    console.log("onSubmit:", name, email);
  };
  
  return (
    <div>
      <h1>My React App</h1>
      <Form onSubmit={handleSubmit} />
    </div>
  );
};

export default App;