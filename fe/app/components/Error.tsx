import React from 'react';

interface ErrorProps {
    message: string;
}

export default function ({message}: ErrorProps) {
    return (
        <div className="mt-4 text-red-500">{message}</div>
    )
}