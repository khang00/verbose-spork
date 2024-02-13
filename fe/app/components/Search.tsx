import FileInput from "@/app/components/FileInput";
import {useState} from "react";
import {GetResultDetails, Search} from "@/app/fetch";
import KeywordView from "@/app/components/KeywordView";
import DetailView from "@/app/components/DetailView";

interface Result {
    id: number,
    keyword: string
}

interface ResultDetails {
    id: number,
    keyword: string,
    resultStats: number,
    numberOfLinks: number,
    numberOfAds: number,
    html: string,
}

export default function () {
    const [keywords, setKeywords] = useState<string[]>([]);
    const [results, setResults] = useState<Result[]>([]);
    const [details, setDetails] = useState<ResultDetails | undefined>();

    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | undefined>('');

    const onChange = async (file: File | undefined) => {
        if (file != undefined) {
            const text = await file.text();
            const keywordsArray = text.split(',');
            setKeywords(keywordsArray);
        }
    }

    const onSearch = async () => {
        if (keywords.length > 0) {
            setIsLoading(true)
            const results = await Search(keywords)
            setResults(results)
            setIsLoading(false)
        }
    }

    const onClickResult = async (result: Result) => {
        const details = await GetResultDetails(result)
        setDetails(details)
    }

    return (
        <div className="flex flex-col items-center bg-white rounded-lg px-8 py-6 mx-8">
            <div className="flex">
                <FileInput label={"Search file"}
                           onChange={onChange}
                           onClick={onSearch}
                           isLoading={isLoading}
                ></FileInput>
                {error && <div className="text-red-500 mt-2">{error}</div>}
            </div>

            <div className="flex w-full flex-nowrap py-6">
                {keywords.length > 0 && (
                    <div className="flex-1 min-w-90">
                        <KeywordView results={results} onClickResult={onClickResult}></KeywordView>
                    </div>
                )}

                {details && (
                    <div className="flex-1 px-6 py-6 min-w-90">
                        <DetailView details={details}></DetailView>
                    </div>)
                }
            </div>
        </div>
    );
}
