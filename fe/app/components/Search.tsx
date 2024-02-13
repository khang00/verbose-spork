import FileInput from "@/app/components/FileInput";
import React, {useState} from "react";
import {GetResultDetails, Search} from "@/app/fetch";
import KeywordView from "@/app/components/KeywordView";
import DetailView from "@/app/components/DetailView";
import Input from "@/app/components/Input";

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

    const [searchKeyword, setSearchKeyword] = useState('');
    const [filteredResults, setFilteredResults] = useState<Result[]>([]);

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

            setFilteredResults(results)
            setIsLoading(false)
        }
    }

    const onClickResult = async (result: Result) => {
        const details = await GetResultDetails(result)
        setDetails(details)
    }

    const onSearchKeyword = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearchKeyword(e.target.value)
        setFilteredResults(results.filter(result => {
            return result.keyword.startsWith(e.target.value)
        }))
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
            <div className="flex items-center py-6 min-w-80">
                {keywords.length > 0 && (
                    <div className="w-full">
                        <Input
                            type="text"
                            label="search for keywords"
                            value={searchKeyword}
                            onChange={onSearchKeyword}
                        ></Input>
                    </div>
                )}
            </div>

            <div className="flex w-full flex-nowrap py-6">
                {keywords.length > 0 && (
                    <div className="flex-1 min-w-90">
                        <KeywordView results={filteredResults} onClickResult={onClickResult}></KeywordView>
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
