"use server";
import {getUserInfo, getUserTokens, userInfo} from "@/api/getUserInfos";
import ProfilePage from "@/components/pages/profile/ProfilePage";

export default async function Profile() {
    try {
        const data: userInfo = await getUserInfo();
        data.providers = await getUserTokens();
        return <ProfilePage {...data} />;
    } catch (err) {
        const data: userInfo = {id: -1, first_name: "", last_name: "", email: "", password: "", providers: []};
        console.log(err)
        return <ProfilePage {...data} />
    }
}
