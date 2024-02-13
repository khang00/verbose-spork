import axios from "axios";

const host = 'http://localhost:8080'
const axiosInstance = axios.create({
    baseURL: host,
    timeout: 1000,
});

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

    axios.interceptors.request.use(function (config) {
        config.headers.Authorization = `Bearer ${signinResp.token}`
        return config;
    });

    return signinResp
}

const parseSigninResp = (data: any): SigninResp => {
    return {
        userID: data.userID,
        username: data.username,
        token: data.token
    }
}

export type {SigninResp, SigninReq}
export {Signin}