interface UserLogInfosBody {
    token: string,
    user_id: number
}

interface UserCredentials {
    email: string,
    password: string
}

interface UserSignUpBody {
    email: string,
    password: string,
    first_name: string,
    last_name: string
}

interface OAuthLoginBody {
    service: string,
    code: string | null,
    redirect_uri: string
}
