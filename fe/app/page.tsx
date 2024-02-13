'use client';
import Auth from "@/app/components/Auth";
import React, {useState} from "react";
import {SigninResp, SignupResp} from "@/app/fetch";

const Page = () => {
    const [user, setUser] = useState({
        userID: 0,
        username: '',
        token: '',
    });

    const onSignin = (signin: SigninResp) => {
        setUser({
            userID: signin.userID,
            username: signin.username,
            token: signin.token,

        })
    }

    const onSignup = (signup: SignupResp) => {
        setUser({
            userID: signup.userID,
            username: signup.username,
            token: signup.token,
        })
    }

    return (
        <main className="flex justify-center items-center w-full h-full">
            {user.username == '' ?
                (<Auth onSignin={onSignin} onSignup={onSignup}></Auth>) :
                (<div></div>)}
        </main>
    );
}

export default Page