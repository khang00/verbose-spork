import React from 'react';

interface Result {
    id: number,
    keyword: string
}

interface KeywordViewProps {
    results: Result[]
    onClickResult: (result: Result) => void
}

export default function ({results, onClickResult}: KeywordViewProps) {
    const resultsList = results.map(result => {
        return (
            <tr key={result.id} className="bg-base-200" onClick={() => onClickResult(result)}>
                <td>{result.id}</td>
                <td>{result.keyword}</td>
            </tr>
        )
    });

    return (
        <table className="table">
            <thead>
            <tr>
                <th>Id</th>
                <th>Keyword</th>
            </tr>
            </thead>
            <tbody>
            {resultsList}
            </tbody>
        </table>
    );
}