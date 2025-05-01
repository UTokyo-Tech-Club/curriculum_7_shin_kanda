import React, { useEffect, useState } from "react";

type User = {
  id: string;
  name: string;
  age: number;
};

function App() {
  const [users, setUsers] = useState<User[]>([]);
  const [name, setName] = useState("");
  const [age, setAge] = useState("");

  const fetchUsers = async () => {
    try {
      const res = await fetch("http://localhost:8000/user");
      const data = await res.json();
      setUsers(data);
    } catch (error) {
      alert("ユーザー取得に失敗しました");
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name || !age || isNaN(Number(age))) {
      alert("正しい名前と年齢を入力してください");
      return;
    }

    try {
      const res = await fetch("http://localhost:8000/user", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name, age: Number(age) }),
      });

      if (!res.ok) {
        throw new Error("送信に失敗しました");
      }

      setName("");
      setAge("");
      fetchUsers();
    } catch (err) {
      alert("ユーザー登録に失敗しました");
    }
  };

  return (
    <div className="flex flex-col min-h-screen bg-gray-100">
      {/* ヘッダー部分 */}
      <div className="bg-gray-800 text-white py-6 px-4">
        <h1 className="text-2xl font-bold text-center">User Register</h1>
      </div>

      <div className="container mx-auto px-4 py-6">
        {/* フォーム部分 */}
        <div className="mb-8">
          <div className="flex flex-col space-y-4 max-w-md mx-auto">
            <div className="flex justify-between items-center">
              <label className="font-medium">Name:</label>
              <input
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                className="border rounded px-3 py-2 w-64"
              />
            </div>
            
            <div className="flex justify-between items-center">
              <label className="font-medium">Age:</label>
              <input
                type="number"
                value={age}
                onChange={(e) => setAge(e.target.value)}
                className="border rounded px-3 py-2 w-64"
              />
            </div>
            
            <button
              onClick={handleSubmit}
              className="bg-gray-200 hover:bg-gray-300 text-black font-medium py-2 rounded transition"
            >
              POST
            </button>
          </div>
        </div>

        {/* ユーザーリスト部分 */}
        <div className="space-y-3">
          {users.map((user) => (
            <div 
              key={user.id}
              className="bg-blue-100 p-4 rounded flex justify-center items-center"
            >
              <span className="text-gray-800 font-medium">
                {user.name}, {user.age}
              </span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;









