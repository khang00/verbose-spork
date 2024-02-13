import FileInput from "@/app/components/FileInput";
import {useState} from "react";
import {Search} from "@/app/fetch";
import KeywordView from "@/app/components/KeywordView";

interface Result {
    id: number,
    keyword: string
}

export default function () {
    const [keywords, setKeywords] = useState<string[]>([]);
    const [results, setResults] = useState<Result[]>([]);

    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>('');

    const onChange = async (file: File | undefined) => {
        if (file != null) {
            const text = await file.text();
            const keywordsArray = text.split(',');
            setKeywords(keywordsArray);
        }
    }

    const onSearch = async () => {
        setIsLoading(true)
        const results = await Search(keywords)
        setResults(results)
        setIsLoading(false)
    }

    return (
        <div>
            <div className="flex justify-center border-solid bg-white rounded-lg border-2 px-8 py-6">
                <div className="flex">
                    <FileInput label={"Search file"}
                               onChange={onChange}
                               onClick={onSearch}
                               isLoading={isLoading}
                    ></FileInput>
                    {error && <div className="text-red-500 mt-2">{error}</div>}
                </div>
            </div>
            <div className="flex justify-center">
                <KeywordView results={results}></KeywordView>
            </div>
        </div>
    );
}
