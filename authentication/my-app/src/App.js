import './App.css';
import React, { useState, useEffect } from 'react';
import  Signup  from './components/Signup';

import { onAuthStateChanged } from "firebase/auth";
import { auth } from "./firebase";

const App = () => {
  // stateとしてログイン状態を管理する。ログインしていないときはnullになる。
  const [loginUser, setLoginUser] = useState(auth.currentUser);

  // ログイン状態を監視して、stateをリアルタイムで更新する
  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, user => {
      setLoginUser(user);
    });
    return () => unsubscribe();
  }, []);

  return (
    <>
  
      
      <Signup />
      {/* ログインしていないと見られないコンテンツは、loginUserがnullの場合表示しない */}
      {loginUser ? <div>ログインユーザー専用コンテンツ</div> : null}
    </>
  );
};

export default App;
