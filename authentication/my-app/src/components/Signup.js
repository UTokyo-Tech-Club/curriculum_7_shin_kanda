import { useState } from 'react';
import { createUserWithEmailAndPassword } from 'firebase/auth';
import { auth } from '../firebase';

const Signup = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit=(event)=>{
        event.preventDefault();
        console.log(email, password);
        createUserWithEmailAndPassword(auth, email, password)
    };
    const handleChangeEmail=(event)=>{
        const {name, value}=event.target;
        if(name==='email'){
            setEmail(value);
        }
        
    }
    const handleChangePassword=(event)=>{
        const {name, value}=event.target;
        if(name==='password'){
            setPassword(value);
        }
    }   

    return(
        <>
        <h1>新規ユーザ登録</h1>
        <form onSubmit={handleSubmit}>
            <div>
                <label>メールアドレス</label>
                <input type="email" name="email" placeholder="email@example.com" value={email} onChange={handleChangeEmail} />
            </div>
            <div>
                <label>パスワード</label>
                <input type="password" name="password" placeholder="６文字以上のパスワードを入力してください" value={password} onChange={handleChangePassword} />
            </div>
            <button type="submit">登録</button>
        </form>
        </>
    );
    
};

export default Signup;