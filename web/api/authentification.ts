import axios from "axios";

export async function login(emailValue: string, passwordValue: string) : Promise<boolean> {
    axios.defaults.withCredentials = true;
    try {
        const response = await axios.post(`http://localhost:3000/login/`, {
            email: emailValue,
            password: passwordValue
        });

        console.log(response);
        console.log(response.data);

        localStorage.setItem('token', response.data);
        return true;
    } catch (error) {
        throw error;
    }
}
export function checkAuthentification()  {
    if (typeof window !== 'undefined') {
        const token = localStorage.getItem("token");
        console.log(token);
        return token != null;
    }
}
