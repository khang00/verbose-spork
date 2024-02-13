'use client';
import React, {useState} from 'react';
import Input from './Input';
import Button from './Button';
import Error from './Error';
import {Signup, SignupResp} from '../fetch';

interface SignupProps {
    onSignup: (signup: SignupResp) => void
}

export default function ({onSignup}: SignupProps) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [signingUp, setSigningUp] = useState(false);

    const [error, setError] = useState('');
    const onClick = async (e: React.MouseEvent) => {
        if (password == confirmPassword) {
            setSigningUp(true)
            const signupResp = await Signup({
                username: username,
                password: password,
            }).catch(error => setError(error))

            if (signupResp != null) {
                onSignup(signupResp)
            }

            setSigningUp(false)
        } else {
            setError('password doesn\'t match')
        }
    }
    const onSubmit = (e: React.FormEvent) => {
        e.preventDefault()
    }

    return (
        <form className="bg-white rounded-lg px-8 py-6" onSubmit={onSubmit}>
            <h2 className="text-2xl font-semibold mb-4">Sign Up</h2>
            <Input type="text" label="Username" value={username}
                   onChange={(e) => setUsername(e.target.value)}/>
            <Input type="password" label="Password" value={password}
                   onChange={(e) => setPassword(e.target.value)}/>
            <Input type="password" label="Confirm Password" value={confirmPassword}
                   onChange={(e) => setConfirmPassword(e.target.value)}/>
            <div className="flex items-center justify-center">
                <Button onClick={onClick} isLoading={signingUp} type="button">Sign Up</Button>
            </div>
            {error && <Error message={error}/>}
        </form>
    );
}
