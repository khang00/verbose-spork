import React from 'react';

interface ButtonProps {
    type?: "button" | "submit" | "reset" | undefined;
    isLoading: boolean
    children: React.ReactNode;
    onClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

export default function ({type = 'button', isLoading, children, onClick}: ButtonProps) {
    return (
        <button
            type={type}
            className="btn btn-block"
            onClick={onClick}
        >
            {isLoading ? (<span className="loading loading-spinner"></span>) : (children)}
        </button>
    )
}