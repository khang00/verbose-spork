import React, {useState} from 'react';
import Signin from './Signin';
import {SigninResp} from "@/app/fetch";

interface AuthProps {
    onSignin: (signin: SigninResp) => void
}

export default function ({onSignin}: AuthProps) {

    return (
        <div className="w-full h-full flex items-center justify-center">
            <div className="border-solid bg-white rounded-lg border-2">
                <Signin onSignin={onSignin}></Signin>
            </div>
        </div>
    );
}
