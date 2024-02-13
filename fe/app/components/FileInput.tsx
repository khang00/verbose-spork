'use client'
import React, {useState} from 'react';
import Button from "@/app/components/Button";

interface FileInputProps {
    label: string;
    onChange: (file: File | undefined) => void;
    onClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
    isLoading: boolean;
}

const FileInput = ({label, onChange, onClick, isLoading}: FileInputProps) => {
    const [selectedFile, setSelectedFile] = useState<File | undefined>(undefined);
    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        setSelectedFile(file);
        onChange(file);
    };

    return (
        <div>
            <label className="block text-gray-700 font-bold mb-2">{label}</label>
            <div className="flex space-x-4 items-center justify-center w-full">
                <input
                    type="file"
                    className="file-input file-input-bordered w-full max-w-xs"
                    accept="text/csv"
                    onChange={handleFileChange}
                />
                <div className="flex items-center justify-center">
                    <Button isLoading={isLoading} onClick={onClick}>Search</Button>
                </div>
            </div>

        </div>
    );
};

export default FileInput;
