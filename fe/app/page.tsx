'use client';
import Auth from "@/app/components/Auth";
import React, {useState} from "react";
import {SigninResp, SignupResp} from "@/app/fetch";
import Search from "@/app/components/Search";

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
        <main>
            {user.username == '' ?
                (<Auth onSignin={onSignin} onSignup={onSignup}></Auth>) :
                (<Search></Search>)}
        </main>
    );
}

export default Page