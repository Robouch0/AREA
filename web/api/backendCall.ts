import axios from "axios";
import Cookies from 'js-cookie';

export function login(emailValue: string, passwordValue: string) {
    axios.post(`http://localhost:3000/login/`, {email: emailValue, password: passwordValue})
        .then((res) => {
            console.log(res);
            console.log(res.data);
            Cookies.set('token', res.data, {expires: 7, secure: true});
            const token = Cookies.get('token');
            axios.defaults.withCredentials = true;
            axios.get(`http://localhost:3000/admin/`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
            .then((res) => {
                    console.log(res.data);
            });
        }).catch(
        function (error) {
            console.log('Show error notification!')
            return Promise.reject(error)
        });
}
