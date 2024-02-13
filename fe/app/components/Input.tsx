import React from 'react';

interface InputProps {
    type: string;
    label: string;
    value: string;
    onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

export default function ({type, label, value, onChange}: InputProps) {
    return (
        <div className="mb-4">
            <label className="form-control w-full max-w-xs">
                <div className="label">
                    <span className="label-text">{label}</span>
                </div>
                <input className="input input-bordered w-full max-w-xs"
                       type={type}
                       onChange={onChange}
                       value={value}/>
            </label>
        </div>
    )
}