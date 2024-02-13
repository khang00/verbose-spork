import React from "react";

interface DetailsViewProps {
    details: ResultDetails
}

interface ResultDetails {
    id: number,
    keyword: string,
    resultStats: number,
    numberOfLinks: number,
    numberOfAds: number,
    html: string,
}

export default function ({details}: DetailsViewProps) {
    return (
        <table className="table">
            <tbody>
            <tr>
                <th>id</th>
                <td>{details.id}</td>
            </tr>
            <tr>
                <th>keyword</th>
                <td>{details.keyword}</td>
            </tr>
            <tr>
                <th>Number of Link</th>
                <td>{details.numberOfLinks}</td>
            </tr>
            <tr>
                <th>Number of adwords</th>
                <td>{details.numberOfAds}</td>
            </tr>
            <tr>
                <th>Number of search results</th>
                <td>{details.resultStats}</td>
            </tr>
            <tr>
                <th>HTML page</th>
                <td>
                    <textarea className="textarea w-full" placeholder="HTML page" disabled>
                        {details.html}
                    </textarea>
                </td>
            </tr>
            </tbody>
        </table>
    );
}