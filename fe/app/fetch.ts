import axios from "axios";

const host = 'http://localhost:8080'
const axiosInstance = axios.create({
    baseURL: host,
    timeout: 60000,
});

const GetToken = (): string | null => {
    return sessionStorage.getItem('token')
}

interface SigninReq {
    username: string,
    password: string,
}

interface SigninResp {
    userID: number,
    username: string,
    token: string,
}

const Signin = async (req: SigninReq): Promise<SigninResp> => {
    const resp = await axiosInstance.post("/api/user/signin", req)
    const signinResp = parseSigninResp(resp.data)
    sessionStorage.setItem('token', signinResp.token)

    return signinResp
}

const parseSigninResp = (data: any): SigninResp => {
    return {
        userID: data.userID,
        username: data.username,
        token: data.token
    }
}

interface SignupReq {
    username: string,
    password: string,
}

interface SignupResp {
    userID: number,
    username: string,
    token: string,
}

const Signup = async (req: SignupReq): Promise<SignupResp> => {
    const resp = await axiosInstance.post("/api/user/signup", req)
    const signupResp = parseSignupResp(resp.data)
    sessionStorage.setItem('token', signupResp.token)

    return signupResp
}

const parseSignupResp = (data: any): SignupResp => {
    return {
        userID: data.userID,
        username: data.username,
        token: data.token
    }
}

type Results = Result[]

interface Result {
    id: number,
    keyword: string,
}

const Search = async (keywords: string[]): Promise<Results> => {
    const token = GetToken()
    const resp = await axiosInstance.post("/api/keyword/upload", {
        keywords: keywords
    }, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })

    return parseKeywordsResponse(resp.data)
}

const parseKeywordsResponse = (data: any): Results => {
    return data.results.map((obj: any) => {
        return {
            id: obj.id,
            keyword: obj.keyword,
        }
    })
}

interface ResultDetails {
    id: number,
    keyword: string,
    resultStats: number,
    numberOfLinks: number,
    numberOfAds: number,
    html: string,
}

const GetResultDetails = async (result: Result): Promise<ResultDetails> => {
    const token = GetToken()
    const resp = await axiosInstance.get(`/api/keyword?id=${result.id}`, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })

    return parseResultDetailsResponse(resp.data)
}

const parseResultDetailsResponse = (data: any): ResultDetails => {
    return {
        id: data.id,
        keyword: data.keyword,
        resultStats: data.result_stats,
        numberOfLinks: data.number_of_links,
        numberOfAds: data.number_of_ads,
        html: data.html,
    }
}

export type {SigninResp, SigninReq, SignupReq, SignupResp, Results, Result}
export {Signin, Signup, Search, GetResultDetails}