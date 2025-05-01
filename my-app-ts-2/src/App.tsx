import logo from "./logo.svg";
import "./App.css";
import { useState } from "react";

function App() {
  const [name, setName] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [age, setAge] = useState<number>(0);

  const handlePost = async (e:React.FormEvent<HTMLFormElement>) => {
    e.preventDefault(); // フォームのリロードを防ぐ

    try {
      const response = await fetch(
        "https://react-1-5-besa-default-rtdb.firebaseio.com/tweets.json",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ name, email, age: Number(age) }),
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      console.log("データ送信成功");

      // フォームをクリア
      setName("");
      setEmail("");
      setAge(0);
    }catch (error) {
      console.error("エラー:", error);
    }
  };
  
  

  const handleGet = async () => {
    try {
      const response = await fetch(
        "https://react-1-5-besa-default-rtdb.firebaseio.com/tweets.json",
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ name, email, age: Number(age) }),
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data: Record<string, { name: string; age: number }> = await response.json();
      const obj = Object.values(data).find((v:any) => v.name === "inada");

      if (obj) {
        setAge(Number(obj.age) + 10);
      } else {
        console.log("指定された名前が見つかりませんでした");
      }
    } catch (error) {
      console.error("エラー:", error);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>

      <form style={{ display: "flex", flexDirection: "column" }} onSubmit={handlePost}>
        <label>Name:</label>
        <input type="text" value={name} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setName(e.target.value)} />

        <label>Email:</label>
        <input type="email" value={email} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)} />

        <label>Age:</label>
        <input type="number" value={age} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setAge(Number(e.target.value))} />

        <button type="submit">Submit</button>
      </form>

      <button onClick={handleGet}>Get</button>
      <p>年齢（inadaの場合）：{age}</p>
    </div>
  );
}

export default App;