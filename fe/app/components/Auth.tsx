import React from 'react';
import Signin from './Signin';
import Signup from './Signup';
import {SigninResp, SignupResp} from "@/app/fetch";

interface AuthProps {
    onSignin: (signin: SigninResp) => void
    onSignup: (signup: SignupResp) => void
}

export default function ({onSignin, onSignup}: AuthProps) {

    return (
        <div className="flex items-center justify-center border-solid bg-white rounded-lg border-2 px-8 py-6">
            <div className="flex">
                <Signin onSignin={onSignin}></Signin>
                <Signup onSignup={onSignup}></Signup>
            </div>
        </div>
    );
}
