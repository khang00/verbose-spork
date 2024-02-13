'use client';
import React, {useState} from 'react';
import Input from './Input';
import Button from './Button';
import Error from './Error';
import {Signin, SigninResp} from '../fetch';

interface SigninProps {
    onSignin: (signin: SigninResp) => void
}

export default function ({onSignin}: SigninProps) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [signingIn, setSigningIn] = useState(false);

    const [error, setError] = useState(null);
    const onClick = async (e: React.MouseEvent) => {
        setSigningIn(true)
        const signinResp = await Signin({
            username: username,
            password: password,
        })
        onSignin(signinResp)
        setSigningIn(false)
    }

    const onSubmit = (e: React.FormEvent) => {
        e.preventDefault()
    }

    return (
        <form className="bg-white rounded-lg px-8 py-6" onSubmit={onSubmit}>
            <h2 className="text-2xl font-semibold mb-4">Login</h2>
            <Input type="text" label="Username" value={username} onChange={(e) => setUsername(e.target.value)}/>
            <Input type="password" label="Password" value={password} onChange={(e) => setPassword(e.target.value)}/>
            <div className="flex items-center justify-center">
                <Button onClick={onClick} isLoading={signingIn} type="button">Login</Button>
            </div>
            {error && <Error message={error}/>}
        </form>
    );
}
