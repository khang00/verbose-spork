'use client';
import Auth from "@/app/components/Auth";
import React, {useState} from "react";
import {SigninResp} from "@/app/fetch";

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

    return (
        <main className="w-full h-full">
            {user.username == '' ?
                (<Auth onSignin={onSignin}></Auth>) :
                (<div></div>)}
        </main>
    );
}

export default Page